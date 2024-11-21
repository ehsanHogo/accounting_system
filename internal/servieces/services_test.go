package servieces

import (
	"accounting_system/config"
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	randgenerator "accounting_system/internal/utils"
	"fmt"
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
		_, err := createTempDetailed(repo)

		// err := InsertDetailed(repo, detailed)

		assert.NoError(t, err, "expected no error when inserting detailed")
	})

	t.Run("can not insert detailed record with empty code", func(t *testing.T) {
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

	t.Run("the detailed record creation fail because duplication code", func(t *testing.T) {

		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error when inserting")
		detailed.Title = randgenerator.GenerateRandomTitle()

		err = InsertDetailed(repo, detailed)

		assert.Error(t, err, "expected getting duplicate detailed code error")

	})

	t.Run("the detailed record creation fail because duplication title", func(t *testing.T) {

		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error when inserting")

		detailed.Code = randgenerator.GenerateRandomCode()

		err = InsertDetailed(repo, detailed)

		assert.Error(t, err, "expected getting duplicate detailed title error")

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

	t.Run("can update detailed record successfully", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed")
		insertedDetailed, err := ReadDetailed(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		insertedDetailed.Code = generateUniqeCode[models.Detailed](repo, "code")
		// insertedDetailed.Title = generateUniqeTitle[models.Detailed](repo)

		err = UpdateDetailed(repo, insertedDetailed)

		assert.NoError(t, err, "expected no error when updating detailed")

		checkUpdated, err := repositories.ReadRecord[models.Detailed](repo, insertedDetailed.Model.ID)
		assert.NoError(t, err, "expected no error when reading detailed record ")
		assert.Equal(t, insertedDetailed.Code, checkUpdated.Code)
	})

	t.Run("can not update detailed record with empty code", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		insertedDetailed.Code = ""
		err = UpdateDetailed(repo, insertedDetailed)

		assert.Error(t, err, "expected error indicate empty code not allowed")
	})

	t.Run("can not update detailed record with empty title", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		insertedDetailed.Title = ""
		err = UpdateDetailed(repo, insertedDetailed)

		assert.Error(t, err, "expected error indicate empty title not allowed")
	})

	t.Run("can not update detailed record with code length greater than 64", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		s := "1"
		for i := 0; i < 70; i++ {
			insertedDetailed.Code += s
		}
		err = UpdateDetailed(repo, insertedDetailed)

		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")
	})

	t.Run("can not update detailed record with title length greater than 64", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		s := "a"
		for i := 0; i < 70; i++ {
			insertedDetailed.Title += s
		}
		err = UpdateDetailed(repo, insertedDetailed)

		assert.Error(t, err, "expected error indicate title length should not be greater than 64 ")
	})

	t.Run("can not update detailed record that is not in databse", func(t *testing.T) {
		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		detailed := &models.Detailed{Code: code, Title: title}
		detailed.Model.ID = 1_000_000
		err := UpdateDetailed(repo, detailed)
		assert.Error(t, err, "expected error indicate there is such id in database")

	})

	t.Run("can not update detailed record if versions were  different", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "cexpected no error while inserting")

		detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
		// fmt.Printf("prev id : %v\n", detailed.Model.ID)
		// fmt.Printf("code : %v\n", detailed.Code)
		// fmt.Printf("prev version : %v\n", detailed.Version)
		err = UpdateDetailed(repo, detailed)
		assert.NoError(t, err, "expected no error while updating")

		detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
		err = UpdateDetailed(repo, detailed)
		// fmt.Printf("new version : %v\n", detailed.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can update detailed record if versions were same", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting")

		detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
		// fmt.Printf("prev id : %v\n", detailed.Model.ID)
		// fmt.Printf("code : %v\n", detailed.Code)
		// fmt.Printf("prev version : %v\n", detailed.Version)
		err = UpdateDetailed(repo, detailed)
		assert.NoError(t, err, "expected no error while updating")

		detailed, _ = repositories.ReadRecord[models.Detailed](repo, detailed.Model.ID)
		detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
		err = UpdateDetailed(repo, detailed)
		fmt.Printf("new version : %v\n", detailed.Version)
		assert.NoError(t, err, "expected no error")

	})

	t.Run("can not update detailed because duplication code", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)

		assert.NoError(t, err, "expected no error while updating detailed ")
		prevCode := detailed.Code

		detailed, err = createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while updating detailed ")

		detailed.Code = prevCode
		err = UpdateDetailed(repo, detailed)

		assert.Error(t, err, "expected getting duplicate detailed code error")

	})

	t.Run("can not update detailed because duplication title", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)

		assert.NoError(t, err, "expected no error while updating detailed ")
		prevTitle := detailed.Title

		detailed, err = createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while updating detailed ")

		detailed.Title = prevTitle
		err = UpdateDetailed(repo, detailed)

		assert.Error(t, err, "expected getting duplicate detailed number error")

	})

}

func TestDeleteDetailed(t *testing.T) {
	repo, err := createConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can delete detailed successfully", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting")

		err = DeleteDetailed(repo, detailed)
		assert.NoError(t, err, "expected no error while deleting")

		result := repo.AccountingDB.First(&detailed)
		assert.Error(t, result.Error, "expected error indicate detailed record not found")
	})

	t.Run("deletion detailed record fail because record does not exist in database", func(t *testing.T) {
		detailed := &models.Detailed{}
		detailed.Model.ID = 1_000_000
		err := DeleteDetailed(repo, detailed)
		// result := repo.AccountingDB.First(&detailed)
		assert.Error(t, err, "expected error indicate detailed record not found")
	})

	t.Run("deletion detailed record fails because there is a reffrence in some voucher items  ", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting")

		voucher, err := createTempVoucher(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while inserting")
		fmt.Printf("det : %v", detailed.Model.ID)
		fmt.Printf("vi : %v", voucher.VoucherItems[0].Model.ID)
		err = DeleteDetailed(repo, detailed)

		assert.Error(t, err, "expected error indicate violation forignkey constraint")

	})

	t.Run("can not delete detailed record if versions were  different", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting")

		detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
		// fmt.Printf("prev id : %v\n", detailed.Model.ID)
		// fmt.Printf("code : %v\n", detailed.Code)
		// fmt.Printf("prev version : %v\n", detailed.Version)
		err = UpdateDetailed(repo, detailed)
		assert.NoError(t, err, "expected no error while updating detailed record")
		err = DeleteDetailed(repo, detailed)
		// fmt.Printf("new version : %v\n", detailed.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can delete detailed record if versions were same", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting")

		detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
		// fmt.Printf("prev id : %v\n", detailed.Model.ID)
		// fmt.Printf("code : %v\n", detailed.Code)
		// fmt.Printf("prev version : %v\n", detailed.Version)
		err = UpdateDetailed(repo, detailed)
		assert.NoError(t, err, "can not update detailed record ")
		detailed, _ = repositories.ReadRecord[models.Detailed](repo, detailed.Model.ID)

		err = DeleteDetailed(repo, detailed)
		// fmt.Printf("new version : %v\n", detailed.Version)
		assert.NoError(t, err, "expected no error")

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

	t.Run("the subsidiary record creation fail because duplication code", func(t *testing.T) {

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "cexpected no error when inserting subsidiary")

		subsidiary.Title = randgenerator.GenerateRandomTitle()

		err = InsertSubsidiary(repo, subsidiary)

		assert.Error(t, err, "expected getting duplicate subsidiary code error")

	})

	t.Run("the subsidiary record creation fail because duplication title", func(t *testing.T) {

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error when inserting subsidiary")

		subsidiary.Code = randgenerator.GenerateRandomCode()

		err = InsertSubsidiary(repo, subsidiary)

		assert.Error(t, err, "expected getting duplicate subsidiary title error")

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
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")

		fetchSubsidiary, err := repositories.ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		fetchSubsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
		// fetchSubsidiary.Title = generateUniqeTitle[models.Subsidiary](repo)
		err = UpdateSubsidiary(repo, fetchSubsidiary)
		assert.NoError(t, err, "expected no error when  updating subsidiary record")
		checkUpdated, err := repositories.ReadRecord[models.Subsidiary](repo, fetchSubsidiary.Model.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		assert.Equal(t, fetchSubsidiary.Code, checkUpdated.Code)
	})

	t.Run("can not update subsidiary due to empty title", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")

		fetchSubsidiary, err := repositories.ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		fetchSubsidiary.Title = ""
		err = UpdateSubsidiary(repo, fetchSubsidiary)
		assert.Error(t, err, "expected error indicate empty code is not allowed")

	})

	t.Run("can not update subsidiary due to empty code", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")

		fetchSubsidiary, err := repositories.ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		fetchSubsidiary.Code = ""
		err = UpdateSubsidiary(repo, fetchSubsidiary)
		assert.Error(t, err, "expected error indicate empty title is not allowed")

	})

	t.Run("can not update subsidiary when  code length is greater than 64", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")

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
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")

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

	t.Run("return error when update subsidiary record that is not in databse", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")
		subsidiary.Model.ID = 1_000_000
		err = UpdateSubsidiary(repo, subsidiary)
		assert.Error(t, err, "expected error indicate there is such id in database")

	})

	t.Run("can not update subsidiary record if versions were  different", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")

		subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
		// fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
		// fmt.Printf("code : %v\n", subsidiary.Code)
		// fmt.Printf("prev version : %v\n", subsidiary.Version)
		err = UpdateSubsidiary(repo, subsidiary)
		assert.NoError(t, err, "expected no error while updating")
		subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
		err = UpdateSubsidiary(repo, subsidiary)
		assert.Error(t, err, "expected no error while updating subsidiary")
		// fmt.Printf("new version : %v\n", subsidiary.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can update subsidiary record if versions were same", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "can not create subsidiary record due to")

		subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
		// fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
		// fmt.Printf("code : %v\n", subsidiary.Code)
		// fmt.Printf("prev version : %v\n", subsidiary.Version)
		UpdateSubsidiary(repo, subsidiary)
		subsidiary, _ = repositories.ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
		err = UpdateSubsidiary(repo, subsidiary)
		// fmt.Printf("new version : %v\n", subsidiary.Version)
		assert.NoError(t, err, "expected no error")

	})

	t.Run("can not update subsidiary record if were reffrenced in some voucher items", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")
		fmt.Println("in me ")
		// fmt.Printf("detialed id : %v\n", subsidiary.Model.ID)
		_, err = createTempVoucher(repo, 0, subsidiary.Model.ID)

		// fmt.Printf("detialed id : %v\n", subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error while inserting")
		// fmt.Printf("voucher id : %v\n", voucher.Model.ID)
		subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
		err = UpdateSubsidiary(repo, subsidiary)
		assert.Error(t, err, "expected error indicate violation update forign key constraint")
	})

	t.Run("can not update subsidiary because duplication code", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)

		assert.NoError(t, err, "expected no error while updating subsidiary ")
		prevCode := subsidiary.Code

		subsidiary, err = createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while updating subsidiary ")

		subsidiary.Code = prevCode
		err = UpdateSubsidiary(repo, subsidiary)

		assert.Error(t, err, "expected getting duplicate subsidiary code error")

	})

	t.Run("can not update subsidiary because duplication title", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)

		assert.NoError(t, err, "expected no error while updating subsidiary ")
		prevTitle := subsidiary.Title

		subsidiary, err = createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while updating subsidiary ")

		subsidiary.Title = prevTitle
		err = UpdateSubsidiary(repo, subsidiary)

		assert.Error(t, err, "expected getting duplicate subsidiary number error")

	})
}

func TestDeleteSubsidiary(t *testing.T) {
	repo, err := createConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("delete subsidiary record seccessfully", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")

		err = DeleteSubsidiary(repo, subsidiary)

		assert.NoError(t, err, "expected no error while deleting ")

	})

	t.Run("deletion subsidiary record fail because record does not exist in database", func(t *testing.T) {
		subsidiary := &models.Subsidiary{}
		subsidiary.Model.ID = 1_000_000
		err := DeleteSubsidiary(repo, subsidiary)

		assert.Error(t, err, "expected error indicate subsiduary record not found")
	})

	t.Run("deletion subsidiary record fails because there is a reffrence in some voucher items  ", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")

		_, err = createTempVoucher(repo, 0, subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error while inserting")
		// fmt.Printf("det : %v", subsidiary.Model.ID)
		// fmt.Printf("vi : %v", voucher.VoucherItems[0].Model.ID)
		err = DeleteSubsidiary(repo, subsidiary)

		assert.Error(t, err, "expected error indicate violation forignkey constraint")

	})

	t.Run("can not delete subsidiary record if versions were  different", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")

		subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
		// fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
		// fmt.Printf("code : %v\n", subsidiary.Code)
		// fmt.Printf("prev version : %v\n", subsidiary.Version)
		err = UpdateSubsidiary(repo, subsidiary)
		assert.NoError(t, err, "can not update subsidiary record ")
		err = DeleteSubsidiary(repo, subsidiary)
		// fmt.Printf("new version : %v\n", subsidiary.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can delete subsidiary record if versions were same", func(t *testing.T) {
		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting")

		subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
		// fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
		// fmt.Printf("code : %v\n", subsidiary.Code)
		// fmt.Printf("prev version : %v\n", subsidiary.Version)
		err = UpdateSubsidiary(repo, subsidiary)
		assert.NoError(t, err, "cann not update subsidiary record due to ")

		subsidiary, _ = repositories.ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)

		err = DeleteSubsidiary(repo, subsidiary)
		// fmt.Printf("new version : %v\n", subsidiary.Version)
		assert.NoError(t, err, "expected no error while deleting")

	})

}

func TestInsertVoucher(t *testing.T) {
	repo, err := createConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can insert voucher successfully", func(t *testing.T) {
		voucher, err := createTempVoucher(repo)
		assert.NoError(t, err, "expected no error when inserting voucher")

		// assert.NoError(t, err, "expected voucher record to be created, but got error")

		var result models.Voucher
		err = repo.AccountingDB.First(&result, voucher.Model.ID).Error //Number is uniqe
		assert.NoError(t, err, " can not find the inserted voucher record :")
	})

	t.Run("can not insert voucher because duplication number", func(t *testing.T) {
		voucher, err := createTempVoucher(repo)

		assert.NoError(t, err, "expected no error while inserting voucher ")

		err = InsertVoucher(repo, voucher)

		assert.Error(t, err, "expected getting duplicate voucher number error")

	})

	t.Run("can not insert voucher record with empty number", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: 100}}}

		err = InsertVoucher(repo, voucher)

		assert.Error(t, err, "expected error indicate empty title not allowed")
	})

	t.Run("can not insert voucher record with number length greater than 64", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: 100}}}
		// detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}

		s := "1"
		for i := 0; i < 70; i++ {
			voucher.Number += s
		}
		err = InsertVoucher(repo, voucher)

		assert.Error(t, err, "expected error indicate number length should not be greater than 64 ")
	})

	t.Run("can not insert voucher record with imbalance voucher items", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 50}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: 100}}}
		// detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}

		err = InsertVoucher(repo, voucher)

		assert.Error(t, err, "expected error indicate voucher items should be balance ")
	})

	t.Run("can not insert voucher record with unvalied voucher items credits od debits", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: -100}}}
		err = InsertVoucher(repo, voucher)
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")

		voucher = &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID}}}
		err = InsertVoucher(repo, voucher)
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")

		voucher = &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100, Debit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100, Debit: 100}}}
		err = InsertVoucher(repo, voucher)
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")
	})

	t.Run("can not insert voucher record with invalied voucher items length ", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		tempList := make([]*models.VoucherItem, 0)

		for i := 0; i < 502; i += 2 {
			tempList = append(tempList, &models.VoucherItem{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100})

			tempList = append(tempList, &models.VoucherItem{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: 100})
		}
		voucher := &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: tempList}
		// detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}

		err = InsertVoucher(repo, voucher)

		assert.Error(t, err, "expected error indicate voucher items length should not be greater than 500 ")

		voucher = &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{}}

		err = InsertVoucher(repo, voucher)

		assert.Error(t, err, "expected error indicate voucher items length should not be less than 2 ")

	})

	t.Run("can not insert voucher record when voucher item subsidiary has detailed and not specifying detailed", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo), HasDetailed: true}
		err = InsertSubsidiary(repo, subsidiary)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, Credit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: 100}}}
		// detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}

		err = InsertVoucher(repo, voucher)

		assert.Error(t, err, "expected error indicate voucher items should have detailed")
	})

	t.Run("can not insert voucher record when voucher item subsidiary does not have detailed and  specifying detailed", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")
		fmt.Println("ftfugg")
		fmt.Println(detailed.Model.ID)
		voucher := &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, Credit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: 100}}}
		// detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}

		err = InsertVoucher(repo, voucher)

		assert.Error(t, err, "expected error indicate voucher items should not have detailed")
	})
}

func TestUpdateVoucher(t *testing.T) {
	repo, err := createConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can update voucher  seccessfully", func(t *testing.T) {
		// voucher, err := createTempVoucher(repo)
		// assert.NoError(t, err, "expected no error while inserting")

		// // fmt.Printf("prev Code %v\n", voucher.Number)
		// // prevVoucherId := voucher.Model.ID

		// // fmt.Printf("new Code %v\n", code)
		// voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
		// err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		// assert.NoError(t, err, "can not update voucher ")

		voucher, err := createTempVoucher(repo)
		assert.NoError(t, err, "can not create voucher record")

		// fmt.Printf("prev Code %v\n", voucher.Number)

		newVoucherItem, err := createTempVoucherItem(repo)
		assert.NoError(t, err, "can not create voucher item record")
		temp := append(voucher.VoucherItems, newVoucherItem)

		newVoucherItem, err = createTempVoucherItem(repo)
		assert.NoError(t, err, "can not create voucher item record")
		temp = append(temp, newVoucherItem)
		temp[len(temp)-1].Credit = 250
		temp[len(temp)-1].Debit = 0

		temp[1].Credit -= 1
		temp[2].Debit -= 1

		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{temp[1], temp[2]}, []*models.VoucherItem{temp[0], temp[3]}, []*models.VoucherItem{temp[4], temp[5]})

		// fmt.Printf("new Code %v\n", code)
		assert.NoError(t, err, "expected no error while updating voucher ")
		_, err = repositories.ReadRecord[models.VoucherItem](repo, voucher.VoucherItems[0].Model.ID)
		assert.Error(t, err, "expected error indicate voucher item not found")

		_, err = repositories.ReadRecord[models.VoucherItem](repo, voucher.VoucherItems[1].Model.ID)
		assert.NoError(t, err, "expexted no error when reading the voucherItem record")
		// assert.Equal(t, voucher.VoucherItems[1].DetailedId, newVoucherItem.DetailedId)
		// assert.Equal(t, voucher.VoucherItems[1].SubsidiaryId, newVoucherItem.SubsidiaryId)
		// assert.Equal(t, voucher.VoucherItems[1].Debit, newVoucherItem.Debit)
		// assert.Equal(t, voucher.VoucherItems[1].Credit, newVoucherItem.Credit)
		// assert.Equal(t, voucher.VoucherItems[1].VoucherID, newVoucherItem.VoucherID)

		_, err = repositories.ReadRecord[models.VoucherItem](repo, voucher.VoucherItems[2].Model.ID)
		assert.NoError(t, err, "expexted no error when reading the voucherItem record")

		_, err = repositories.ReadRecord[models.VoucherItem](repo, voucher.VoucherItems[3].Model.ID)
		assert.Error(t, err, "expexted no error when reading the voucherItem record")

		_, err = repositories.ReadRecord[models.VoucherItem](repo, temp[4].Model.ID)
		assert.NoError(t, err, "expexted no error when reading the voucherItem record")

		_, err = repositories.ReadRecord[models.VoucherItem](repo, temp[5].Model.ID)
		assert.NoError(t, err, "expexted no error when reading the voucherItem record")
	})

	t.Run("can not update voucher because duplication number", func(t *testing.T) {
		voucher, err := createTempVoucher(repo)

		assert.NoError(t, err, "expected no error while inserting voucher ")
		prevNumber := voucher.Number

		voucher, err = createTempVoucher(repo)
		assert.NoError(t, err, "expected no error while inserting voucher ")

		voucher.Number = prevNumber
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate  duplicate voucher number")

	})

	t.Run("can not update voucher because empty number", func(t *testing.T) {
		voucher, err := createTempVoucher(repo)

		assert.NoError(t, err, "expected no error while updating voucher ")

		voucher.Number = ""
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate  empty voucher number ")

	})

	t.Run("can not update voucher record with number length greater than 64", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		_, err = createTempVoucher(repo)

		assert.NoError(t, err, "expected no error while updating voucher ")

		voucher := &models.Voucher{VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: 100}}}
		// detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}

		s := "1"
		for i := 0; i < 70; i++ {
			voucher.Number += s
		}
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate number length should not be greater than 64 ")
	})

	t.Run("can not update voucher record with unvalied voucher items credits od debits", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")
		_, err = createTempVoucher(repo)
		assert.NoError(t, err, "expected no error while inserting voucher ")

		voucher := &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: -100}}}
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")

		voucher = &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID}}}
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")

		voucher = &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100, Debit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100, Debit: 100}}}
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")
	})

	t.Run("can not update voucher record with invalied voucher items length ", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher, err := createTempVoucher(repo)
		assert.NoError(t, err, "expected no error while inserting voucher ")
		tempList := make([]*models.VoucherItem, 0)

		for i := 0; i < 502; i += 2 {
			tempList = append(tempList, &models.VoucherItem{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100})

			tempList = append(tempList, &models.VoucherItem{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: 100})
		}
		// voucher := &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: tempList}
		// detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}

		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, tempList)

		assert.Error(t, err, "expected error indicate voucher items length should not be greater than 500 ")

		tempList = make([]*models.VoucherItem, 0)

		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, tempList, []*models.VoucherItem{})
		assert.NoError(t, err, "expected no error while deleting voucher")
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, tempList)

		assert.Error(t, err, "expected error indicate voucher items length should not be less than 2 ")

	})

	t.Run("can not update voucher record when voucher item subsidiary has detailed and not specifying detailed", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo), HasDetailed: true}
		err = InsertSubsidiary(repo, subsidiary)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")
		_, err = createTempVoucher(repo)
		assert.NoError(t, err, "expected no error while inserting voucher ")

		voucher := &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, Credit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: 100}}}
		// detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}

		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate voucher items should have detailed")
	})

	t.Run("can not update voucher record when voucher item subsidiary does not have detailed and  specifying detailed", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")
		fmt.Println("ftfugg")
		fmt.Println(detailed.Model.ID)

		_, err = createTempVoucher(repo)
		assert.NoError(t, err, "expected no error while inserting voucher ")

		voucher := &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, Credit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: 100}}}
		// detailed := &models.Detailed{Title: generateUniqeTitle[models.Detailed](repo)}

		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate voucher items should not have detailed")
	})

	t.Run("return error when update voucher record that is not in databse", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := createTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{Number: generateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Credit: 100}, {SubsidiaryId: subsidiary.Model.ID, DetailedId: detailed.Model.ID, Debit: -100}}}
		voucher.Model.ID = 1_000_000
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.Error(t, err, "expected error indicate there is such id in database")

	})

	t.Run("can not update voucher record if versions were  different", func(t *testing.T) {
		voucher, err := createTempVoucher(repo)
		assert.NoError(t, err, "can not create voucher record")

		voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
		// fmt.Printf("prev id : %v\n", voucher.Model.ID)
		// fmt.Printf("code : %v\n", voucher.Number)
		fmt.Printf("prev version : %v\n", voucher.Version)
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.NoError(t, err, "can not update voucher record")

		voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		fmt.Printf("new version : %v\n", voucher.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can update voucher record if versions were same", func(t *testing.T) {
		voucher, err := createTempVoucher(repo)
		assert.NoError(t, err, "can not create voucher record")

		voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
		// fmt.Printf("prev id : %v\n", voucher.Model.ID)
		// fmt.Printf("code : %v\n", voucher.Number)
		// fmt.Printf("prev version : %v\n", voucher.Version)
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.NoError(t, err, "can not update voucher record")

		voucher, _ = repositories.ReadRecord[models.Voucher](repo, voucher.Model.ID)
		voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		// fmt.Printf("new version : %v\n", voucher.Version)
		assert.NoError(t, err, "expected no error")

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

func createTempVoucher(repo *repositories.Repositories, IDs ...uint) (*models.Voucher, error) {
	temp := make([]*models.VoucherItem, 4)

	subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo), HasDetailed: true}
	err := InsertSubsidiary(repo, subsidiary)
	if err != nil {
		return nil, err
	}

	if len(IDs) == 0 {

		detailed, err := createTempDetailed(repo)

		if err != nil {
			return nil, err
		}

		temp[0] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Credit: 250}

		temp[1] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Credit: 250}

		temp[2] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Debit: 250}
		temp[3] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Debit: 250}
	} else {
		temp = make([]*models.VoucherItem, 2)

		if len(IDs) == 1 {

			temp[0] = &models.VoucherItem{DetailedId: IDs[0], SubsidiaryId: subsidiary.Model.ID, Credit: 500}
			temp[1] = &models.VoucherItem{DetailedId: IDs[0], SubsidiaryId: subsidiary.Model.ID, Debit: 500}
		} else {
			// detailed, err := createTempDetailed(repo)
			// if err != nil {
			// 	return nil, err
			// }
			temp[0] = &models.VoucherItem{SubsidiaryId: IDs[1], Credit: 500}
			temp[1] = &models.VoucherItem{SubsidiaryId: IDs[1], Debit: 500}

		}

	}

	number := generateUniqeCode[models.Voucher](repo, "number")
	voucher := &models.Voucher{Number: number, VoucherItems: temp}

	// err := errors.New("")
	// for err != nil {

	err = InsertVoucher(repo, voucher)
	if err != nil {
		return nil, fmt.Errorf("Error during record creation: %v\n", err)

	}

	// }

	return voucher, nil
}

func createTempVoucherItem(repo *repositories.Repositories) (*models.VoucherItem, error) {

	subsidiary, err := createTempSubsidiary(repo)
	if err != nil {
		return nil, err
	}

	return &models.VoucherItem{SubsidiaryId: subsidiary.Model.ID, Debit: 250}, nil
}

func createTempSubsidiary(repo *repositories.Repositories) (*models.Subsidiary, error) {
	subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo), HasDetailed: false}

	err := InsertSubsidiary(repo, subsidiary)
	if err != nil {
		return nil, fmt.Errorf("Error during record creation: %v\n", err)

	}

	return subsidiary, nil
}

func createTempDetailed(repo *repositories.Repositories) (*models.Detailed, error) {

	detailed := &models.Detailed{Code: generateUniqeCode[models.Detailed](repo, "code"), Title: generateUniqeTitle[models.Detailed](repo)}

	// err := errors.New("")
	// for err != nil {

	err := InsertDetailed(repo, detailed)
	if err != nil {
		return nil, fmt.Errorf("Error during record creation: %v\n", err)

	}

	return detailed, nil

	// }
}
