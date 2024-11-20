package servieces

import (
	"accounting_system/config"
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	randgenerator "accounting_system/internal/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createConnectionForTest() (*repositories.Repositories, error) {
	dbUrl, err := config.SetupConfig()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)                // Limit the maximum number of open connections
	sqlDB.SetMaxIdleConns(5)                  // Set idle connection limit
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Limit connection lifetime

	return repositories.NewConnection(db), nil
}
func TestInsertDetailed(t *testing.T) {

	repo, err := createConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can insert detailed record successfully", func(t *testing.T) {
		detailed := &models.Detailed{Code: generateUniqeCode[models.Detailed](repo, "code"), Title: generateUniqeTitle[models.Detailed](repo)}

		err := InsertDetailed(repo, detailed)

		assert.NoError(t, err, "expected no error when inserting detailed")
	})

	t.Run("can not insert detailed record with emptu code", func(t *testing.T) {
		detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}

		err := InsertDetailed(repo, detailed)

		assert.Error(t, err, "expected error indicate empty code not allowed")
	})

	t.Run("can not insert detailed record with empty title", func(t *testing.T) {
		detailed := &models.Detailed{Code: generateUniqeCode[models.Detailed](repo, "code")}

		err := InsertDetailed(repo, detailed)

		assert.Error(t, err, "expected error indicate empty title not allowed")
	})

	t.Run("can not insert detailed record with code length greater than 64", func(t *testing.T) {
		detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}
		s := "1"
		for i := 0; i < 70; i++ {
			detailed.Code += s
		}
		err := InsertDetailed(repo, detailed)

		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")
	})

	t.Run("can not insert detailed record with title length greater than 64", func(t *testing.T) {
		detailed := &models.Detailed{Code: generateUniqeCode[models.Detailed](repo, "code")}
		s := "a"
		for i := 0; i < 70; i++ {
			detailed.Title += s
		}
		err := InsertDetailed(repo, detailed)

		assert.Error(t, err, "expected error indicate title length should not be greater than 64 ")
	})
}

func TestUpdateDetailed(t *testing.T) {

	repo, err := createConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can update detailed successfully", func(t *testing.T) {
		detailed := &models.Detailed{Code: generateUniqeCode[models.Detailed](repo, "code"), Title: generateUniqeTitle[models.Detailed](repo)}
		err := InsertDetailed(repo, detailed)
		assert.NoError(t, err, "can not insert detailed record")

		fetchDetailed, err := repositories.ReadRecord[models.Detailed](repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error when reading detailed record ")
		fetchDetailed.Code = generateUniqeCode[models.Detailed](repo, "code")
		err = UpdateDetailed(repo, fetchDetailed)
		assert.NoError(t, err, "expected no error when  updating detailed record")
		checkUpdated, err := repositories.ReadRecord[models.Detailed](repo, fetchDetailed.Model.ID)
		assert.NoError(t, err, "expected no error when reading detailed record ")
		assert.Equal(t, fetchDetailed.Code, checkUpdated.Code)
	})

	t.Run("can not update detailed due to empty title", func(t *testing.T) {
		detailed := &models.Detailed{Code: generateUniqeCode[models.Detailed](repo, "code"), Title: generateUniqeTitle[models.Detailed](repo)}
		err := InsertDetailed(repo, detailed)
		assert.NoError(t, err, "can not insert detailed record")

		fetchDetailed, err := repositories.ReadRecord[models.Detailed](repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error when reading detailed record ")
		fetchDetailed.Title = ""
		err = UpdateDetailed(repo, fetchDetailed)
		assert.Error(t, err, "expected error indicate empty code is not allowed")

	})

	t.Run("can not update detailed due to empty code", func(t *testing.T) {
		detailed := &models.Detailed{Code: generateUniqeCode[models.Detailed](repo, "code"), Title: generateUniqeTitle[models.Detailed](repo)}
		err := InsertDetailed(repo, detailed)
		assert.NoError(t, err, "can not insert detailed record")

		fetchDetailed, err := repositories.ReadRecord[models.Detailed](repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error when reading detailed record ")
		fetchDetailed.Code = ""
		err = UpdateDetailed(repo, fetchDetailed)
		assert.Error(t, err, "expected error indicate empty title is not allowed")

	})

	t.Run("can not update detailed when  code length is greater than 64", func(t *testing.T) {
		detailed := &models.Detailed{Code: generateUniqeCode[models.Detailed](repo, "code"), Title: generateUniqeTitle[models.Detailed](repo)}
		err := InsertDetailed(repo, detailed)
		assert.NoError(t, err, "can not insert detailed record")

		fetchDetailed, err := repositories.ReadRecord[models.Detailed](repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error when reading detailed record ")
		s := "1"
		fetchDetailed.Code = ""
		for i := 0; i < 70; i++ {
			fetchDetailed.Code += s
		}

		err = UpdateDetailed(repo, fetchDetailed)
		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")

	})

	t.Run("can not update detailed when  title length is greater than 64", func(t *testing.T) {
		detailed := &models.Detailed{Code: generateUniqeCode[models.Detailed](repo, "code"), Title: generateUniqeTitle[models.Detailed](repo)}
		err := InsertDetailed(repo, detailed)
		assert.NoError(t, err, "can not insert detailed record")

		fetchDetailed, err := repositories.ReadRecord[models.Detailed](repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error when reading detailed record ")
		s := "a"
		fetchDetailed.Code = ""
		for i := 0; i < 70; i++ {
			fetchDetailed.Code += s
		}

		err = UpdateDetailed(repo, fetchDetailed)
		assert.Error(t, err, "expected error indicate title length should not be greater than 64 ")

	})
}

func generateUniqeCode[T any](repo *repositories.Repositories, columnName string) string {
	code := randgenerator.GenerateRandomCode()
	for {
		exist := repositories.FindRecord[T](repo, code, columnName)

		if exist {
			code = randgenerator.GenerateRandomCode()
		} else {
			break
		}
	}

	return code
}

func generateUniqeTitle[T any](repo *repositories.Repositories) string {
	title := randgenerator.GenerateRandomTitle()
	for {
		exist := repositories.FindRecord[T](repo, title, "title")

		if exist {
			title = randgenerator.GenerateRandomTitle()
		} else {
			break
		}
	}

	return title
}
