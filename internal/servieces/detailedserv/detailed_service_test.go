package detailedserv

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/utils/casting"
	"accounting_system/internal/utils/randgenerator"
	"accounting_system/internal/utils/temporary"

	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertDetailed(t *testing.T) {

	repo, err := repositories.CreateConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can insert detailed record successfully", func(t *testing.T) {
		_, err := temporary.CreateTempDetailed(repo)

		// err := InsertDetailed(repo, detailed)

		assert.NoError(t, err, "expected no error when inserting detailed")
	})

	t.Run("can not insert detailed record with empty code", func(t *testing.T) {
		detailed := &models.InsertDetailedRequest{Title: randgenerator.GenerateUniqeTitle[models.Detailed](repo)}

		err := InsertDetailed(repo, detailed)

		assert.Error(t, err, "expected error indicate empty code not allowed")
	})

	t.Run("can not insert detailed record with empty title", func(t *testing.T) {
		detailed := &models.InsertDetailedRequest{Code: randgenerator.GenerateUniqeCode[models.Detailed](repo, "code")}

		err := InsertDetailed(repo, detailed)

		assert.Error(t, err, "expected error indicate empty title not allowed")
	})

	t.Run("can not insert detailed record with code length greater than 64", func(t *testing.T) {
		detailed := &models.InsertDetailedRequest{Title: randgenerator.GenerateUniqeTitle[models.Detailed](repo)}
		s := "1"
		for i := 0; i < 70; i++ {
			detailed.Code += s
		}
		err := InsertDetailed(repo, detailed)

		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")
	})

	t.Run("can not insert detailed record with title length greater than 64", func(t *testing.T) {
		detailed := &models.InsertDetailedRequest{Code: randgenerator.GenerateUniqeCode[models.Detailed](repo, "code")}
		s := "a"
		for i := 0; i < 70; i++ {
			detailed.Title += s
		}
		err := InsertDetailed(repo, detailed)

		assert.Error(t, err, "expected error indicate title length should not be greater than 64 ")
	})

	t.Run("the detailed record creation fail because duplication code", func(t *testing.T) {

		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error when inserting")
		// detailed.Title = randgenerator.GenerateRandomTitle()
		newDetailed := &models.InsertDetailedRequest{Code: detailed.Code, Title: randgenerator.GenerateRandomTitle()}
		err = InsertDetailed(repo, newDetailed)

		assert.Error(t, err, "expected getting duplicate detailed code error")

	})

	t.Run("the detailed record creation fail because duplication title", func(t *testing.T) {

		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error when inserting")

		// detailed.Code = randgenerator.GenerateRandomCode()
		newDetailed := &models.InsertDetailedRequest{Code: randgenerator.GenerateRandomCode(), Title: detailed.Title}

		err = InsertDetailed(repo, newDetailed)

		assert.Error(t, err, "expected getting duplicate detailed title error")

	})
}

func TestUpdateDetailed(t *testing.T) {

	repo, err := repositories.CreateConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can update detailed record successfully", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed")
		insertedDetailed, err := ReadDetailed(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		// insertedDetailed.Code = randgenerator.GenerateUniqeCode[models.Detailed](repo, "code")
		// insertedDetailed.Title = randgenerator.GenerateUniqeTitle[models.Detailed](repo)
		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(insertedDetailed.Model.ID), Code: randgenerator.GenerateUniqeCode[models.Detailed](repo, "code"), Title: insertedDetailed.Title}
		err = UpdateDetailed(repo, updateDetailed)

		assert.NoError(t, err, "expected no error when updating detailed")

		checkUpdated, err := repositories.ReadRecord[models.Detailed](repo, insertedDetailed.Model.ID)
		assert.NoError(t, err, "expected no error when reading detailed record ")
		assert.Equal(t, updateDetailed.Code, checkUpdated.Code)
	})

	t.Run("can not update detailed record with empty code", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		// insertedDetailed.Code = ""

		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(insertedDetailed.Model.ID), Title: insertedDetailed.Title}

		err = UpdateDetailed(repo, updateDetailed)

		assert.Error(t, err, "expected error indicate empty code not allowed")
	})

	t.Run("can not update detailed record with empty title", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")
		// insertedDetailed.Title = ""
		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(insertedDetailed.Model.ID), Code: insertedDetailed.Code}

		err = UpdateDetailed(repo, updateDetailed)

		assert.Error(t, err, "expected error indicate empty title not allowed")
	})

	t.Run("can not update detailed record with code length greater than 64", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")

		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(insertedDetailed.Model.ID), Title: insertedDetailed.Title}

		s := "1"
		for i := 0; i < 70; i++ {
			updateDetailed.Code += s
		}

		err = UpdateDetailed(repo, updateDetailed)

		assert.Error(t, err, "expected error indicate code length should not be greater than 64 ")
	})

	t.Run("can not update detailed record with title length greater than 64", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting detailed")

		insertedDetailed, err := ReadDetailed(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed")

		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(insertedDetailed.Model.ID), Code: insertedDetailed.Code}

		s := "a"
		for i := 0; i < 70; i++ {
			updateDetailed.Title += s
		}
		err = UpdateDetailed(repo, updateDetailed)

		assert.Error(t, err, "expected error indicate title length should not be greater than 64 ")
	})

	t.Run("can not update detailed record that is not in databse", func(t *testing.T) {
		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		detailed := &models.UpdateDetailedRequest{Code: code, Title: title}
		detailed.ID = "1000000"
		err := UpdateDetailed(repo, detailed)
		assert.Error(t, err, "expected error indicate there is such id in database")

	})

	t.Run("can not update detailed record if versions were  different", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "cexpected no error while inserting")

		// detailed.Code = randgenerator.GenerateUniqeCode[models.Detailed](repo, "code")

		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(detailed.Model.ID), Code: randgenerator.GenerateUniqeCode[models.Detailed](repo, "code"), Title: detailed.Title, Version: casting.UintToString(detailed.Version)}

		// fmt.Printf("prev id : %v\n", detailed.Model.ID)
		// fmt.Printf("code : %v\n", detailed.Code)
		// fmt.Printf("prev version : %v\n", detailed.Version)
		err = UpdateDetailed(repo, updateDetailed)
		assert.NoError(t, err, "expected no error while updating")

		updateDetailed.Code = randgenerator.GenerateUniqeCode[models.Detailed](repo, "code")
		err = UpdateDetailed(repo, updateDetailed)
		// fmt.Printf("new version : %v\n", detailed.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can update detailed record if versions were same", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting")

		// detailed.Code = randgenerator.GenerateUniqeCode[models.Detailed](repo, "code")

		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(detailed.Model.ID), Code: randgenerator.GenerateUniqeCode[models.Detailed](repo, "code"), Title: detailed.Title}

		// fmt.Printf("prev id : %v\n", detailed.Model.ID)
		// fmt.Printf("code : %v\n", detailed.Code)
		// fmt.Printf("prev version : %v\n", detailed.Version)
		err = UpdateDetailed(repo, updateDetailed)
		assert.NoError(t, err, "expected no error while updating")

		detailed, _ = repositories.ReadRecord[models.Detailed](repo, detailed.Model.ID)
		// detailed.Code = randgenerator.GenerateUniqeCode[models.Detailed](repo, "code")
		updateDetailed = &models.UpdateDetailedRequest{ID: casting.UintToString(detailed.Model.ID), Code: randgenerator.GenerateUniqeCode[models.Detailed](repo, "code"), Title: detailed.Title, Version: casting.UintToString(detailed.Version)}

		err = UpdateDetailed(repo, updateDetailed)
		fmt.Printf("new version : %v\n", detailed.Version)
		assert.NoError(t, err, "expected no error")

	})

	t.Run("can not update detailed because duplication code", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)

		assert.NoError(t, err, "expected no error while updating detailed ")
		prevCode := detailed.Code

		detailed, err = temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while updating detailed ")
		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(detailed.Model.ID), Code: prevCode, Title: detailed.Title}

		// detailed.Code = prevCode
		err = UpdateDetailed(repo, updateDetailed)

		assert.Error(t, err, "expected getting duplicate detailed code error")

	})

	t.Run("can not update detailed because duplication title", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)

		assert.NoError(t, err, "expected no error while updating detailed ")
		prevTitle := detailed.Title

		detailed, err = temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while updating detailed ")

		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(detailed.Model.ID), Code: detailed.Code, Title: prevTitle}

		// detailed.Title = prevTitle
		err = UpdateDetailed(repo, updateDetailed)

		assert.Error(t, err, "expected getting duplicate detailed number error")

	})

}

func TestDeleteDetailed(t *testing.T) {
	repo, err := repositories.CreateConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can delete detailed successfully", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)
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
		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting")

		voucher, err := temporary.CreateTempVoucher(repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while inserting")
		fmt.Printf("det : %v", detailed.Model.ID)
		fmt.Printf("vi : %v", voucher.VoucherItems[0].Model.ID)
		err = DeleteDetailed(repo, detailed)

		assert.Error(t, err, "expected error indicate violation forignkey constraint")

	})

	t.Run("can not delete detailed record if versions were  different", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting")

		// detailed.Code = randgenerator.GenerateUniqeCode[models.Detailed](repo, "code")
		// fmt.Printf("prev id : %v\n", detailed.Model.ID)
		// fmt.Printf("code : %v\n", detailed.Code)
		// fmt.Printf("prev version : %v\n", detailed.Version)
		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(detailed.Model.ID), Code: randgenerator.GenerateUniqeCode[models.Detailed](repo, "code"), Title: detailed.Title}

		err = UpdateDetailed(repo, updateDetailed)
		assert.NoError(t, err, "expected no error while updating detailed record")
		err = DeleteDetailed(repo, detailed)
		// fmt.Printf("new version : %v\n", detailed.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can delete detailed record if versions were same", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while inserting")

		// detailed.Code = randgenerator.GenerateUniqeCode[models.Detailed](repo, "code")
		// fmt.Printf("prev id : %v\n", detailed.Model.ID)
		// fmt.Printf("code : %v\n", detailed.Code)
		// fmt.Printf("prev version : %v\n", detailed.Version)

		updateDetailed := &models.UpdateDetailedRequest{ID: casting.UintToString(detailed.Model.ID), Code: randgenerator.GenerateUniqeCode[models.Detailed](repo, "code"), Title: detailed.Title}

		err = UpdateDetailed(repo, updateDetailed)
		assert.NoError(t, err, "can not update detailed record ")
		detailed, _ = repositories.ReadRecord[models.Detailed](repo, detailed.Model.ID)

		err = DeleteDetailed(repo, detailed)
		// fmt.Printf("new version : %v\n", detailed.Version)
		assert.NoError(t, err, "expected no error")

	})

}

// func temporary.CreateTempDetailed(repo *repositories.Repositories) (*models.Detailed, error) {

// 	detailed := &models.Detailed{Code: randgenerator.GenerateUniqeCode[models.Detailed](repo, "code"), Title: randgenerator.GenerateUniqeTitle[models.Detailed](repo)}

// 	// err := errors.New("")
// 	// for err != nil {

// 	err := detailedserv.InsertDetailed(repo, detailed)
// 	if err != nil {
// 		return nil, fmt.Errorf("error during record creation: %v", err)

// 	}

// 	return detailed, nil

// 	// }
// }
