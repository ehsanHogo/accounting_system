package detailedserv

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/utils/randgenerator"
	"accounting_system/internal/utils/temporary"
	"accounting_system/internal/validations"

	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const InvalidRecordID = 1_000_000

func SetupTestRepo(t *testing.T) *repositories.Repositories {
	repo, err := repositories.CreateConnectionForTest()
	if err != nil {
		t.Fatalf("cannot connect to database: %v", err)
	}

	t.Cleanup(func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	})
	return repo
}
func TestInsertDetailed(t *testing.T) {

	repo := SetupTestRepo(t)
	t.Run("can insert detailed record successfully", func(t *testing.T) {
		_, err := temporary.CreateTempDetailed(repo.AccountingDB)

		assert.NoError(t, err, "expected no error when inserting detailed")
	})

	t.Run("can not insert detailed record with empty code", func(t *testing.T) {
		detailed := &models.Detailed{Title: repositories.GenerateUniqeTitle[models.Detailed](repo.AccountingDB)}

		err := InsertDetailed(repo.AccountingDB, detailed)

		assert.Error(t, err, "expected error indicate empty code not allowed")
	})

	t.Run("can not insert detailed record with empty title", func(t *testing.T) {
		detailed := &models.Detailed{Code: repositories.GenerateUniqeCode[models.Detailed](repo.AccountingDB, "code")}

		err := InsertDetailed(repo.AccountingDB, detailed)

		assert.Error(t, err, "expected error indicate empty title not allowed")
	})

	t.Run("can not insert detailed record with code length greater than 64", func(t *testing.T) {
		detailed := &models.Detailed{Title: repositories.GenerateUniqeTitle[models.Detailed](repo.AccountingDB)}
		s := "1"
		for i := 0; i < validations.MaxCodeLength+1; i++ {
			detailed.Code += s
		}
		err := InsertDetailed(repo.AccountingDB, detailed)

		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")
	})

	t.Run("can not insert detailed record with title length greater than 64", func(t *testing.T) {
		detailed := &models.Detailed{Code: repositories.GenerateUniqeCode[models.Detailed](repo.AccountingDB, "code")}
		s := "a"
		for i := 0; i < validations.MaxTitleLength+1; i++ {
			detailed.Title += s
		}
		err := InsertDetailed(repo.AccountingDB, detailed)

		assert.Error(t, err, "expected error indicate title length should not be greater than 64 ")
	})

	t.Run("the detailed record creation fail because duplication code", func(t *testing.T) {

		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error when inserting")
		detailed.Title = randgenerator.GenerateRandomTitle()

		err = InsertDetailed(repo.AccountingDB, detailed)

		assert.Error(t, err, "expected getting duplicate detailed code error")

	})

	t.Run("the detailed record creation fail because duplication title", func(t *testing.T) {

		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error when inserting")

		detailed.Code = randgenerator.GenerateRandomCode()

		err = InsertDetailed(repo.AccountingDB, detailed)

		assert.Error(t, err, "expected getting duplicate detailed title error")

	})
}

func TestUpdateDetailed(t *testing.T) {

	repo := SetupTestRepo(t)

	t.Run("can update detailed record successfully", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed")
		insertedDetailed, err := ReadDetailed(repo.AccountingDB, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		insertedDetailed.Code = repositories.GenerateUniqeCode[models.Detailed](repo.AccountingDB, "code")

		err = UpdateDetailed(repo.AccountingDB, insertedDetailed)

		assert.NoError(t, err, "expected no error when updating detailed")

		checkUpdated, err := ReadDetailed(repo.AccountingDB, insertedDetailed.Model.ID)
		assert.NoError(t, err, "expected no error when reading detailed record ")
		assert.Equal(t, insertedDetailed.Code, checkUpdated.Code)
	})

	t.Run("can not update detailed record with empty code", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo.AccountingDB, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		insertedDetailed.Code = ""
		err = UpdateDetailed(repo.AccountingDB, insertedDetailed)

		assert.Error(t, err, "expected error indicate empty code not allowed")
	})

	t.Run("can not update detailed record with empty title", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo.AccountingDB, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		insertedDetailed.Title = ""
		err = UpdateDetailed(repo.AccountingDB, insertedDetailed)

		assert.Error(t, err, "expected error indicate empty title not allowed")
	})

	t.Run("can not update detailed record with code length greater than 64", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo.AccountingDB, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		s := "1"
		for i := 0; i < validations.MaxCodeLength+1; i++ {
			insertedDetailed.Code += s
		}
		err = UpdateDetailed(repo.AccountingDB, insertedDetailed)

		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")
	})

	t.Run("can not update detailed record with title length greater than 64", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo.AccountingDB, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		s := "a"
		for i := 0; i < validations.MaxTitleLength+1; i++ {
			insertedDetailed.Title += s
		}
		err = UpdateDetailed(repo.AccountingDB, insertedDetailed)

		assert.Error(t, err, "expected error indicate title length should not be greater than 64 ")
	})

	t.Run("can not update detailed record that is not in databse", func(t *testing.T) {
		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		detailed := &models.Detailed{Code: code, Title: title}
		detailed.Model.ID = InvalidRecordID
		err := UpdateDetailed(repo.AccountingDB, detailed)
		assert.Error(t, err, "expected error indicate there is such id in database")

	})

	t.Run("can not update detailed record if versions were  different", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "cexpected no error while inserting")

		detailed.Code = repositories.GenerateUniqeCode[models.Detailed](repo.AccountingDB, "code")

		err = UpdateDetailed(repo.AccountingDB, detailed)
		assert.NoError(t, err, "expected no error while updating")

		detailed.Code = repositories.GenerateUniqeCode[models.Detailed](repo.AccountingDB, "code")
		err = UpdateDetailed(repo.AccountingDB, detailed)
		assert.Error(t, err, "expected error indicate the versions are different")
		// err = UpdateDetailed(repo.AccountingDB, detailed)
		// assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can update detailed record if versions were same", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		detailed.Code = repositories.GenerateUniqeCode[models.Detailed](repo.AccountingDB, "code")

		err = UpdateDetailed(repo.AccountingDB, detailed)
		assert.NoError(t, err, "expected no error while updating")

		detailed, _ = ReadDetailed(repo.AccountingDB, detailed.Model.ID)
		detailed.Code = repositories.GenerateUniqeCode[models.Detailed](repo.AccountingDB, "code")
		err = UpdateDetailed(repo.AccountingDB, detailed)
		fmt.Printf("new version : %v\n", detailed.Version)
		assert.NoError(t, err, "expected no error")

	})

	t.Run("can not update detailed because duplication code", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)

		assert.NoError(t, err, "expected no error while updating detailed ")
		prevCode := detailed.Code

		detailed, err = temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while updating detailed ")

		detailed.Code = prevCode
		err = UpdateDetailed(repo.AccountingDB, detailed)

		assert.Error(t, err, "expected getting duplicate detailed code error")

	})

	t.Run("can not update detailed because duplication title", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)

		assert.NoError(t, err, "expected no error while updating detailed ")
		prevTitle := detailed.Title

		detailed, err = temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while updating detailed ")

		detailed.Title = prevTitle
		err = UpdateDetailed(repo.AccountingDB, detailed)

		assert.Error(t, err, "expected getting duplicate detailed number error")

	})

}

func TestDeleteDetailed(t *testing.T) {
	repo := SetupTestRepo(t)
	t.Run("can delete detailed successfully", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		err = DeleteDetailed(repo.AccountingDB, detailed)
		assert.NoError(t, err, "expected no error while deleting")

		result := repo.AccountingDB.First(&detailed)
		assert.Error(t, result.Error, "expected error indicate detailed record not found")
	})

	t.Run("deletion detailed record fail because record does not exist in database", func(t *testing.T) {
		detailed := &models.Detailed{}
		detailed.Model.ID = InvalidRecordID
		err := DeleteDetailed(repo.AccountingDB, detailed)
		assert.Error(t, err, "expected error indicate detailed record not found")
	})

	t.Run("deletion detailed record fails because there is a reffrence in some voucher items  ", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while inserting")
		fmt.Printf("det : %v", detailed.Model.ID)
		fmt.Printf("vi : %v", voucher.VoucherItems[0].Model.ID)
		err = DeleteDetailed(repo.AccountingDB, detailed)

		assert.Error(t, err, "expected error indicate violation forignkey constraint")

	})

	t.Run("can not delete detailed record if versions were  different", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		detailed.Code = repositories.GenerateUniqeCode[models.Detailed](repo.AccountingDB, "code")

		err = UpdateDetailed(repo.AccountingDB, detailed)
		assert.NoError(t, err, "expected no error while updating detailed record")
		err = DeleteDetailed(repo.AccountingDB, detailed)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can delete detailed record if versions were same", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		detailed.Code = repositories.GenerateUniqeCode[models.Detailed](repo.AccountingDB, "code")

		err = UpdateDetailed(repo.AccountingDB, detailed)
		assert.NoError(t, err, "can not update detailed record ")
		detailed, _ = ReadDetailed(repo.AccountingDB, detailed.Model.ID)

		err = DeleteDetailed(repo.AccountingDB, detailed)
		assert.NoError(t, err, "expected no error")

	})

}

func TestReadDetailed(t *testing.T) {
	repo := SetupTestRepo(t)
	t.Run("can read the detailed record successfully", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while craeting detailed")

		_, err = ReadDetailed(repo.AccountingDB, detailed.Model.ID)
		assert.NoError(t, err, "expected no error")

	})

	t.Run("return error when the detailed record is not in database ", func(t *testing.T) {

		_, err := ReadDetailed(repo.AccountingDB, InvalidRecordID)
		assert.Error(t, err, "expected  error indicate can not found the detailed record")

	})
}
