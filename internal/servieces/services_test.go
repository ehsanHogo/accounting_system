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

func TestInsertSubsidiary(t *testing.T) {

	repo, err := createConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can insert subsidiary record successfully", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo)}

		err := InsertSubsidiary(repo, subsidiary)

		assert.NoError(t, err, "expected no error when inserting subsidiary")
	})

	t.Run("can not insert subsidiary record with empty code", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Title: generateUniqeTitle[models.Subsidiary](repo)}

		err := InsertSubsidiary(repo, subsidiary)

		assert.Error(t, err, "expected error indicate empty code not allowed")
	})

	t.Run("can not insert subsidiary record with empty title", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code")}

		err := InsertSubsidiary(repo, subsidiary)

		assert.Error(t, err, "expected error indicate empty title not allowed")
	})

	t.Run("can not insert subsidiary record with code length greater than 64", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Title: generateUniqeTitle[models.Subsidiary](repo)}
		s := "1"
		for i := 0; i < 70; i++ {
			subsidiary.Code += s
		}
		err := InsertSubsidiary(repo, subsidiary)

		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")
	})

	t.Run("can not insert subsidiary record with title length greater than 64", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code")}
		s := "a"
		for i := 0; i < 70; i++ {
			subsidiary.Title += s
		}
		err := InsertSubsidiary(repo, subsidiary)

		assert.Error(t, err, "expected error indicate title length should not be greater than 64 ")
	})
}

func TestUpdateSubsidiary(t *testing.T) {

	repo, err := createConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can update subsidiary successfully", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo)}
		err := InsertSubsidiary(repo, subsidiary)
		assert.NoError(t, err, "can not insert subsidiary record")

		fetchSubsidiary, err := repositories.ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		fetchSubsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
		err = UpdateSubsidiary(repo, fetchSubsidiary)
		assert.NoError(t, err, "expected no error when  updating subsidiary record")
		checkUpdated, err := repositories.ReadRecord[models.Subsidiary](repo, fetchSubsidiary.Model.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		assert.Equal(t, fetchSubsidiary.Code, checkUpdated.Code)
	})

	t.Run("can not update subsidiary due to empty title", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo)}
		err := InsertSubsidiary(repo, subsidiary)
		assert.NoError(t, err, "can not insert subsidiary record")

		fetchSubsidiary, err := repositories.ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		fetchSubsidiary.Title = ""
		err = UpdateSubsidiary(repo, fetchSubsidiary)
		assert.Error(t, err, "expected error indicate empty code is not allowed")

	})

	t.Run("can not update subsidiary due to empty code", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo)}
		err := InsertSubsidiary(repo, subsidiary)
		assert.NoError(t, err, "can not insert subsidiary record")

		fetchSubsidiary, err := repositories.ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		fetchSubsidiary.Code = ""
		err = UpdateSubsidiary(repo, fetchSubsidiary)
		assert.Error(t, err, "expected error indicate empty title is not allowed")

	})

	t.Run("can not update subsidiary when  code length is greater than 64", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo)}
		err := InsertSubsidiary(repo, subsidiary)
		assert.NoError(t, err, "can not insert subsidiary record")

		fetchSubsidiary, err := repositories.ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		s := "1"
		fetchSubsidiary.Code = ""
		for i := 0; i < 70; i++ {
			fetchSubsidiary.Code += s
		}

		err = UpdateSubsidiary(repo, fetchSubsidiary)
		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")

	})

	t.Run("can not update subsidiary when  title length is greater than 64", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo)}
		err := InsertSubsidiary(repo, subsidiary)
		assert.NoError(t, err, "can not insert subsidiary record")

		fetchSubsidiary, err := repositories.ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		s := "a"
		fetchSubsidiary.Code = ""
		for i := 0; i < 70; i++ {
			fetchSubsidiary.Code += s
		}

		err = UpdateSubsidiary(repo, fetchSubsidiary)
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
