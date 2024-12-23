package subsidiaryserv

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/utils/randgenerator"
	"accounting_system/internal/utils/temporary"
	"accounting_system/internal/validations"

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
func TestInsertSubsidiary(t *testing.T) {

	repo := SetupTestRepo(t)
	t.Run("can insert subsidiary record successfully", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code"), Title: repositories.GenerateUniqeTitle[models.Subsidiary](repo.AccountingDB)}

		err := InsertSubsidiary(repo.AccountingDB, subsidiary)

		assert.NoError(t, err, "expected no error when inserting subsidiary")
	})

	t.Run("can not insert subsidiary record with empty code", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Title: repositories.GenerateUniqeTitle[models.Subsidiary](repo.AccountingDB)}

		err := InsertSubsidiary(repo.AccountingDB, subsidiary)

		assert.Error(t, err, "expected error indicate empty code not allowed")
	})

	t.Run("can not insert subsidiary record with empty title", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code")}

		err := InsertSubsidiary(repo.AccountingDB, subsidiary)

		assert.Error(t, err, "expected error indicate empty title not allowed")
	})

	t.Run("can not insert subsidiary record with code length greater than 64", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Title: repositories.GenerateUniqeTitle[models.Subsidiary](repo.AccountingDB)}
		s := "1"
		for i := 0; i < validations.MaxCodeLength+1; i++ {
			subsidiary.Code += s
		}
		err := InsertSubsidiary(repo.AccountingDB, subsidiary)

		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")
	})

	t.Run("can not insert subsidiary record with title length greater than 64", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code")}
		s := "a"
		for i := 0; i < validations.MaxTitleLength+1; i++ {
			subsidiary.Title += s
		}
		err := InsertSubsidiary(repo.AccountingDB, subsidiary)

		assert.Error(t, err, "expected error indicate title length should not be greater than 64 ")
	})

	t.Run("the subsidiary record creation fail because duplication code", func(t *testing.T) {

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "cexpected no error when inserting subsidiary")

		subsidiary.Title = randgenerator.GenerateRandomTitle()

		err = InsertSubsidiary(repo.AccountingDB, subsidiary)

		assert.Error(t, err, "expected getting duplicate subsidiary code error")

	})

	t.Run("the subsidiary record creation fail because duplication title", func(t *testing.T) {

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error when inserting subsidiary")

		subsidiary.Code = randgenerator.GenerateRandomCode()

		err = InsertSubsidiary(repo.AccountingDB, subsidiary)

		assert.Error(t, err, "expected getting duplicate subsidiary title error")

	})
}

func TestUpdateSubsidiary(t *testing.T) {

	repo := SetupTestRepo(t)

	t.Run("can update subsidiary successfully", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		fetchSubsidiary, err := ReadSubsidiary(repo.AccountingDB, subsidiary.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		fetchSubsidiary.Code = repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code")
		err = UpdateSubsidiary(repo.AccountingDB, fetchSubsidiary)
		assert.NoError(t, err, "expected no error when  updating subsidiary record")
		checkUpdated, err := ReadSubsidiary(repo.AccountingDB, fetchSubsidiary.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		assert.Equal(t, fetchSubsidiary.Code, checkUpdated.Code)
	})

	t.Run("can not update subsidiary due to empty title", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		fetchSubsidiary, err := ReadSubsidiary(repo.AccountingDB, subsidiary.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		fetchSubsidiary.Title = ""
		err = UpdateSubsidiary(repo.AccountingDB, fetchSubsidiary)
		assert.Error(t, err, "expected error indicate empty code is not allowed")

	})

	t.Run("can not update subsidiary due to empty code", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		fetchSubsidiary, err := ReadSubsidiary(repo.AccountingDB, subsidiary.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		fetchSubsidiary.Code = ""
		err = UpdateSubsidiary(repo.AccountingDB, fetchSubsidiary)
		assert.Error(t, err, "expected error indicate empty title is not allowed")

	})

	t.Run("can not update subsidiary when  code length is greater than 64", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		fetchSubsidiary, err := ReadSubsidiary(repo.AccountingDB, subsidiary.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		s := "1"
		fetchSubsidiary.Code = ""
		for i := 0; i < validations.MaxCodeLength+1; i++ {
			fetchSubsidiary.Code += s
		}

		err = UpdateSubsidiary(repo.AccountingDB, fetchSubsidiary)
		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")

	})

	t.Run("can not update subsidiary when  title length is greater than 64", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		fetchSubsidiary, err := ReadSubsidiary(repo.AccountingDB, subsidiary.ID)
		assert.NoError(t, err, "expected no error when reading subsidiary record ")
		s := "a"
		fetchSubsidiary.Title = ""
		for i := 0; i < validations.MaxTitleLength+1; i++ {
			fetchSubsidiary.Title += s
		}

		err = UpdateSubsidiary(repo.AccountingDB, fetchSubsidiary)
		assert.Error(t, err, "expected error indicate title length should not be greater than 64 ")

	})

	t.Run("return error when update subsidiary record that is not in databse", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")
		subsidiary.ID = InvalidRecordID
		err = UpdateSubsidiary(repo.AccountingDB, subsidiary)
		assert.Error(t, err, "expected error indicate there is such id in database")

	})

	t.Run("can not update subsidiary record if versions were  different", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		subsidiary.Code = repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code")

		err = UpdateSubsidiary(repo.AccountingDB, subsidiary)
		assert.NoError(t, err, "expected no error while updating")
		subsidiary.Code = repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code")
		err = UpdateSubsidiary(repo.AccountingDB, subsidiary)
		assert.Error(t, err, "expected no error while updating subsidiary")
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can update subsidiary record if versions were same", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "can not create subsidiary record due to")

		subsidiary.Code = repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code")

		UpdateSubsidiary(repo.AccountingDB, subsidiary)
		subsidiary, _ = ReadSubsidiary(repo.AccountingDB, subsidiary.ID)
		subsidiary.Code = repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code")
		err = UpdateSubsidiary(repo.AccountingDB, subsidiary)
		assert.NoError(t, err, "expected no error")

	})

	t.Run("can not update subsidiary record if were reffrenced in some voucher items", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		_, err = temporary.CreateTempVoucher(repo.AccountingDB, 0, subsidiary.ID)

		assert.NoError(t, err, "expected no error while inserting")
		subsidiary.Code = repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code")
		err = UpdateSubsidiary(repo.AccountingDB, subsidiary)
		assert.Error(t, err, "expected error indicate violation update forign key constraint")
	})

	t.Run("can not update subsidiary because duplication code", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)

		assert.NoError(t, err, "expected no error while updating subsidiary ")
		prevCode := subsidiary.Code

		subsidiary, err = temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while updating subsidiary ")

		subsidiary.Code = prevCode
		err = UpdateSubsidiary(repo.AccountingDB, subsidiary)

		assert.Error(t, err, "expected getting duplicate subsidiary code error")

	})

	t.Run("can not update subsidiary because duplication title", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)

		assert.NoError(t, err, "expected no error while updating subsidiary ")
		prevTitle := subsidiary.Title

		subsidiary, err = temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while updating subsidiary ")

		subsidiary.Title = prevTitle
		err = UpdateSubsidiary(repo.AccountingDB, subsidiary)

		assert.Error(t, err, "expected getting duplicate subsidiary number error")

	})
}

func TestDeleteSubsidiary(t *testing.T) {
	repo := SetupTestRepo(t)

	t.Run("delete subsidiary record seccessfully", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		err = DeleteSubsidiary(repo.AccountingDB, subsidiary)

		assert.NoError(t, err, "expected no error while deleting ")

	})

	t.Run("deletion subsidiary record fail because record does not exist in database", func(t *testing.T) {
		subsidiary := &models.Subsidiary{}
		subsidiary.ID = InvalidRecordID
		err := DeleteSubsidiary(repo.AccountingDB, subsidiary)

		assert.Error(t, err, "expected error indicate subsiduary record not found")
	})

	t.Run("deletion subsidiary record fails because there is a reffrence in some voucher items  ", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		_, err = temporary.CreateTempVoucher(repo.AccountingDB, 0, subsidiary.ID)
		assert.NoError(t, err, "expected no error while inserting")

		err = DeleteSubsidiary(repo.AccountingDB, subsidiary)

		assert.Error(t, err, "expected error indicate violation forignkey constraint")

	})

	t.Run("can not delete subsidiary record if versions were  different", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		subsidiary.Code = repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code")

		err = UpdateSubsidiary(repo.AccountingDB, subsidiary)
		assert.NoError(t, err, "can not update subsidiary record ")
		err = DeleteSubsidiary(repo.AccountingDB, subsidiary)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can delete subsidiary record if versions were same", func(t *testing.T) {
		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting")

		subsidiary.Code = repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code")

		err = UpdateSubsidiary(repo.AccountingDB, subsidiary)
		assert.NoError(t, err, "cann not update subsidiary record due to ")

		subsidiary, _ = ReadSubsidiary(repo.AccountingDB, subsidiary.ID)

		err = DeleteSubsidiary(repo.AccountingDB, subsidiary)
		assert.NoError(t, err, "expected no error while deleting")

	})

}

func TestReadSubsidiary(t *testing.T) {
	repo := SetupTestRepo(t)
	t.Run("can read the subsidairy record successfully", func(t *testing.T) {
		subsidairy, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while craeting subsidairy")

		_, err = ReadSubsidiary(repo.AccountingDB, subsidairy.ID)
		assert.NoError(t, err, "expected no error")

	})

	t.Run("return error when the subsidairy record is not in database ", func(t *testing.T) {

		_, err := ReadSubsidiary(repo.AccountingDB, InvalidRecordID)
		assert.Error(t, err, "expected  error indicate can not found the subsidairy record")

	})
}
