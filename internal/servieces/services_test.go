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

	t.Run("can not update detailed record if were reffrenced in some voucher items", func(t *testing.T) {
		detailed, err := createTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting")
		fmt.Println("in me ")
		fmt.Printf("detialed id : %v\n", detailed.Model.ID)
		_, err = createTempVoucher(repo, detailed.Model.ID)

		fmt.Printf("detialed id : %v\n", detailed.Model.ID)
		assert.NoError(t, err, "expected no error while inserting")
		// fmt.Printf("voucher id : %v\n", voucher.Model.ID)
		detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
		err = UpdateDetailed(repo, detailed)
		assert.Error(t, err, "expected error indicate violation update forign key constraint")
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
	temp := make([]*models.VoucherItem, 3)
	if len(IDs) == 0 {

		detailed, err := createTempDetailed(repo)
		if err != nil {
			return nil, err
		}

		subsidiary, err := createTempSubsidiary(repo)
		if err != nil {
			return nil, err
		}

		temp[0] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Credit: 500}

		temp[1] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Debit: 250}

		temp[2] = &models.VoucherItem{DetailedId: detailed.Model.ID, Debit: 250}
	} else {
		temp = make([]*models.VoucherItem, 2)

		if len(IDs) == 1 {
			subsidiary, err := createTempSubsidiary(repo)
			if err != nil {
				return nil, err
			}
			temp[0] = &models.VoucherItem{DetailedId: IDs[0], Credit: 500}
			temp[1] = &models.VoucherItem{DetailedId: IDs[0], SubsidiaryId: subsidiary.Model.ID, Debit: 500}
		} else {
			detailed, err := createTempDetailed(repo)
			if err != nil {
				return nil, err
			}
			temp[0] = &models.VoucherItem{DetailedId: detailed.Model.ID, Credit: 500}
			temp[1] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: IDs[1], Debit: 500}

		}

	}

	number := generateUniqeCode[models.Voucher](repo, "number")
	voucher := &models.Voucher{Number: number, VoucherItems: temp}

	// err := errors.New("")
	// for err != nil {

	err := InsertVoucher(repo, voucher)
	if err != nil {
		return nil, fmt.Errorf("Error during record creation: %v\n", err)

	}

	// }

	return voucher, nil
}

func createTempVoucherItem(repo *repositories.Repositories) (*models.VoucherItem, error) {
	detailed, err := createTempDetailed(repo)
	if err != nil {
		return nil, err
	}

	subsidiary, err := createTempSubsidiary(repo)
	if err != nil {
		return nil, err
	}

	return &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Debit: 250}, nil
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