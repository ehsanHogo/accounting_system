package repositories

import (
	"accounting_system/config"
	"accounting_system/internal/models"
	randgenerator "accounting_system/internal/utils"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createConnectionForTest() (*Repositories, error) {
	dbUrl, err := config.SetupConfig()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return NewConnection(db), nil
}

func TestCreateDetailed(t *testing.T) {

	repo, err := createConnectionForTest()

	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("the detailed record successfully create", func(t *testing.T) {
		// Start a transaction
		tx := repo.AccountingDB.Begin()

		// Defer rollback or commit logic
		defer func() {
			if t.Failed() {
				// Rollback if the test fails
				tx.Rollback()
				fmt.Println("Transaction rolled back due to test failure")
			} else {
				// Commit if the test passes
				if err := tx.Commit().Error; err != nil {
					fmt.Printf("Transaction commit failed: %v\n", err)
				} else {
					fmt.Println("Transaction committed successfully")
				}
			}
		}()

		// Generate random code and title
		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()

		// Create a new detailed record within the transaction
		detailed := &models.Detailed{Code: code, Title: title}

		err := errors.New("")
		for err != nil {

			err = CreateRecord(repo, detailed)
			if err != nil {
				fmt.Printf("Error during record creation: %v\n", err)
				code = randgenerator.GenerateRandomCode()
				title = randgenerator.GenerateRandomTitle()
				detailed = &models.Detailed{Code: code, Title: title}
			}
		}

		fmt.Printf("Detailed ID: %v\n", detailed.Model.ID)
		assert.NoError(t, err, "expected detailed record to be created, but got error")

		// Query the record within the same transaction
		var result models.Detailed
		err = tx.First(&result, "id = ?", detailed.Model.ID).Error // Query using the transaction
		if err != nil {
			fmt.Printf("Error during record query: %v\n", err)
		}
		assert.NoError(t, err, "cannot find the inserted detailed record")
	})

	// t.Run("the detailed record successfully create", func(t *testing.T) {
	// 	code := randgenerator.GenerateRandomCode()
	// 	title := randgenerator.GenerateRandomTitle()
	// 	detailed := &models.Detailed{Code: code, Title: title}

	// 	err := errors.New("")
	// 	for err != nil {

	// 		err = CreateRecord(repo, detailed)
	// 		if err != nil {
	// 			fmt.Printf("Error during record creation: %v\n", err)
	// 			code = randgenerator.GenerateRandomCode()
	// 			title = randgenerator.GenerateRandomTitle()
	// 			detailed = &models.Detailed{Code: code, Title: title}
	// 		}
	// 	}

	// 	fmt.Printf("detailed id : %v", detailed.Model.ID)
	// 	assert.NoError(t, err, "expected detailed record to be created, but got error")
	// 	var result models.Detailed
	// 	err = repo.AccountingDB.First(&result, detailed.Model.ID).Error //Code is uniqe
	// 	assert.NoError(t, err, " can not find the inserted detailed record :")

	// })

	// t.Run("the detailed record creation fail because duplication code", func(t *testing.T) {

	// 	code := randgenerator.GenerateRandomCode()
	// 	title := randgenerator.GenerateRandomTitle()
	// 	detailed := &models.Detailed{Code: code, Title: title}
	// 	err := CreateRecord(repo, detailed)
	// 	assert.NoError(t, err, "expected detailed record to be created, but got error")

	// 	title = randgenerator.GenerateRandomTitle()
	// 	detailed = &models.Detailed{Code: code, Title: title}
	// 	err = CreateRecord(repo, detailed)

	// 	assert.Error(t, err, "expected getting duplicate detailed code error")

	// })

	// t.Run("the detailed record creation fail because duplication title", func(t *testing.T) {

	// 	code := randgenerator.GenerateRandomCode()
	// 	title := randgenerator.GenerateRandomTitle()
	// 	detailed := &models.Detailed{Code: code, Title: title}
	// 	err := CreateRecord(repo, detailed)
	// 	assert.NoError(t, err, "expected detailed record to be created, but got error")

	// 	code = randgenerator.GenerateRandomCode()
	// 	detailed = &models.Detailed{Code: code, Title: title}
	// 	err = CreateRecord(repo, detailed)

	// 	assert.Error(t, err, "expected getting duplicate detailed title error")

	// })

}

func TestCreateSubsidiary(t *testing.T) {

	repo, err := createConnectionForTest()

	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("the subsidiary record successfully create", func(t *testing.T) {
		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		subsidiary := &models.Subsidiary{Code: code, Title: title, HasDetailed: false}
		err := CreateRecord(repo, subsidiary)

		assert.NoError(t, err, "expected subsidiary record to be created, but got error")
		var result models.Subsidiary
		err = repo.AccountingDB.First(&result, subsidiary.Model.ID).Error //Code is uniqe
		assert.NoError(t, err, " can not find the inserted subsidiary record :")

	})

	t.Run("the subsidiary record creation fail because duplication code", func(t *testing.T) {

		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		subsidiary := &models.Subsidiary{Code: code, Title: title, HasDetailed: true}
		err := CreateRecord(repo, subsidiary)
		assert.NoError(t, err, "expected subsidiary record to be created, but got error")

		title = randgenerator.GenerateRandomTitle()
		subsidiary = &models.Subsidiary{Code: code, Title: title, HasDetailed: false}
		err = CreateRecord(repo, subsidiary)

		assert.Error(t, err, "expected getting duplicate subsidiary code error")

	})

	t.Run("the subsidiary record creation fail because duplication title", func(t *testing.T) {

		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		subsidiary := &models.Subsidiary{Code: code, Title: title, HasDetailed: false}
		err := CreateRecord(repo, subsidiary)
		assert.NoError(t, err, "expected subsidiary record to be created, but got error")

		code = randgenerator.GenerateRandomCode()
		subsidiary = &models.Subsidiary{Code: code, Title: title, HasDetailed: true}
		err = CreateRecord(repo, subsidiary)

		assert.Error(t, err, "expected getting duplicate subsidiary title error")

	})

}

func TestCreateVoucher(t *testing.T) {

	repo, err := createConnectionForTest()

	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("the voucher record successfully create", func(t *testing.T) {

		var tempDetailedId uint = 2
		var tempSubsidiaryId uint
		code := randgenerator.GenerateRandomCode()
		temp := make([]*models.VoucherItem, 2)
		temp[0] = &models.VoucherItem{Credit: 1, DetailedId: tempDetailedId, SubsidiaryId: tempSubsidiaryId}
		temp[1] = &models.VoucherItem{Credit: 2, DetailedId: tempDetailedId, SubsidiaryId: tempSubsidiaryId}
		voucher := &models.Voucher{Number: code, VoucherItems: temp}
		err := CreateRecord(repo, voucher)

		assert.NoError(t, err, "expected voucher record to be created, but got error")

		var result models.Voucher
		err = repo.AccountingDB.First(&result, voucher.Model.ID).Error //Number is uniqe
		assert.NoError(t, err, " can not find the inserted voucher record :")

	})

	t.Run("the voucher record creation fail because duplication number", func(t *testing.T) {
		var tempDetailedId uint = 2
		var tempSubsidiaryId uint = 4
		code := randgenerator.GenerateRandomCode()

		temp := make([]*models.VoucherItem, 2)
		temp[0] = &models.VoucherItem{Credit: 1, DetailedId: tempDetailedId, SubsidiaryId: tempSubsidiaryId}
		temp[1] = &models.VoucherItem{Credit: 2, DetailedId: tempDetailedId, SubsidiaryId: tempSubsidiaryId}
		voucher := &models.Voucher{Number: code, VoucherItems: temp}
		err := CreateRecord(repo, voucher)
		assert.NoError(t, err, "expected voucher record to be created, but got error")

		voucher = &models.Voucher{Number: code, VoucherItems: temp}
		err = CreateRecord(repo, voucher)

		assert.Error(t, err, "expected getting duplicate voucher number error")

	})

}

func TestUpdateDetailed(t *testing.T) {

	repo, err := createConnectionForTest()

	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}
	t.Run("can update detailed record successfully", func(t *testing.T) {
		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		detailed := &models.Detailed{Code: code, Title: title}
		CreateRecord(repo, detailed)

		prevDetailedId := detailed.Model.ID
		code = randgenerator.GenerateRandomCode()
		title = randgenerator.GenerateRandomTitle()
		detailed = &models.Detailed{Code: code, Title: title}
		err := UpdateDetailed(repo, detailed, prevDetailedId)
		assert.NoError(t, err, "expected no error")
	})

	t.Run("return error when update detailed record that is not in databse", func(t *testing.T) {
		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		detailed := &models.Detailed{Code: code, Title: title}

		err := UpdateDetailed(repo, detailed, 1_000_000)
		assert.Error(t, err, "expected error indicate there is such id in database")

	})

	t.Run("can not update detailed record if versions were  different", func(t *testing.T) {
		detailed := createTempDetailed()
		CreateRecord(repo, detailed)
		detailed.Code = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", detailed.Model.ID)
		fmt.Printf("code : %v\n", detailed.Code)
		fmt.Printf("prev version : %v\n", detailed.Version)
		UpdateDetailed(repo, detailed, detailed.Model.ID)
		detailed.Code = randgenerator.GenerateRandomCode()
		err := UpdateDetailed(repo, detailed, detailed.Model.ID)
		fmt.Printf("new version : %v\n", detailed.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can update detailed record if versions were same", func(t *testing.T) {
		detailed := createTempDetailed()
		CreateRecord(repo, detailed)
		detailed.Code = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", detailed.Model.ID)
		fmt.Printf("code : %v\n", detailed.Code)
		fmt.Printf("prev version : %v\n", detailed.Version)
		UpdateDetailed(repo, detailed, detailed.Model.ID)
		detailed, _ = ReadRecord[models.Detailed](repo, detailed.Model.ID, "detailed")
		detailed.Code = randgenerator.GenerateRandomCode()
		err := UpdateDetailed(repo, detailed, detailed.Model.ID)
		fmt.Printf("new version : %v\n", detailed.Version)
		assert.NoError(t, err, "expected no error")

	})

	t.Run("can not update detailed record if were reffrenced in some voucher items", func(t *testing.T) {
		detailed := createTempDetailed()
		CreateRecord(repo, detailed)
		voucher := createTempVoucher()
		voucher.VoucherItems = append(voucher.VoucherItems, &models.VoucherItem{DetailedId: detailed.Model.ID})
		fmt.Printf("detialed id : %v\n", detailed.Model.ID)
		CreateRecord(repo, voucher)
		fmt.Printf("voucher id : %v\n", voucher.Model.ID)
		err := UpdateDetailed(repo, detailed, detailed.Model.ID)

		assert.Error(t, err, "expected error indicate violation update forign key constraint")
	})

}

func TestUpdateSubsidiary(t *testing.T) {

	repo, err := createConnectionForTest()

	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}
	t.Run("can update subsidiary record successfully", func(t *testing.T) {
		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		subsidiary := &models.Subsidiary{Code: code, Title: title, HasDetailed: false}
		CreateRecord(repo, subsidiary)

		prevSubsidiaryId := subsidiary.Model.ID
		code = randgenerator.GenerateRandomCode()
		title = randgenerator.GenerateRandomTitle()
		subsidiary = &models.Subsidiary{Code: code, Title: title, HasDetailed: true}
		err := UpdateSubsidiary(repo, subsidiary, prevSubsidiaryId)
		assert.NoError(t, err, "expected no error")
	})

	t.Run("return error when update subsidiary record that is not in databse", func(t *testing.T) {
		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		subsidiary := &models.Subsidiary{Code: code, Title: title, HasDetailed: false}

		err := UpdateSubsidiary(repo, subsidiary, 1_000_000)
		assert.Error(t, err, "expected error indicate there is such id in database")

	})

	t.Run("can not update subsidiary record if versions were  different", func(t *testing.T) {
		subsidiary := createTempSubsidiary()
		CreateRecord(repo, subsidiary)
		subsidiary.Code = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
		fmt.Printf("code : %v\n", subsidiary.Code)
		fmt.Printf("prev version : %v\n", subsidiary.Version)
		UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
		subsidiary.Code = randgenerator.GenerateRandomCode()
		err := UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
		fmt.Printf("new version : %v\n", subsidiary.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can update subsidiary record if versions were same", func(t *testing.T) {
		subsidiary := createTempSubsidiary()
		CreateRecord(repo, subsidiary)
		subsidiary.Code = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
		fmt.Printf("code : %v\n", subsidiary.Code)
		fmt.Printf("prev version : %v\n", subsidiary.Version)
		UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
		subsidiary, _ = ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID, "subsidiary")
		subsidiary.Code = randgenerator.GenerateRandomCode()
		err := UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
		fmt.Printf("new version : %v\n", subsidiary.Version)
		assert.NoError(t, err, "expected no error")

	})

}

func TestUpdateVoucher(t *testing.T) {

	repo, err := createConnectionForTest()

	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}
	t.Run("can update voucher record successfully", func(t *testing.T) {
		// code := randgenerator.GenerateRandomCode()
		voucher := createTempVoucher()
		CreateRecord(repo, voucher)
		fmt.Printf("prev Code %v\n", voucher.Number)
		prevVoucherId := voucher.Model.ID
		code := randgenerator.GenerateRandomCode()
		temp := append(voucher.VoucherItems, createTempVoucherItem())
		temp[1].Credit = 12

		fmt.Printf("new Code %v\n", code)
		voucher = &models.Voucher{Number: code, VoucherItems: temp}
		err := UpdateVoucher(repo, voucher, []*models.VoucherItem{temp[1]}, []*models.VoucherItem{temp[0]}, []*models.VoucherItem{temp[2]}, prevVoucherId)
		assert.NoError(t, err, "expected no error")
	})

	t.Run("return error when update voucher record that is not in databse", func(t *testing.T) {

		voucher := createTempVoucher()

		err := UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, 1_000_000)
		assert.Error(t, err, "expected error indicate there is such id in database")

	})

	t.Run("can not update voucher record if versions were  different", func(t *testing.T) {
		voucher := createTempVoucher()
		CreateRecord(repo, voucher)
		voucher.Number = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", voucher.Model.ID)
		fmt.Printf("code : %v\n", voucher.Number)
		fmt.Printf("prev version : %v\n", voucher.Version)
		UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)
		voucher.Number = randgenerator.GenerateRandomCode()
		err := UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)

		fmt.Printf("new version : %v\n", voucher.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can update voucher record if versions were same", func(t *testing.T) {
		voucher := createTempVoucher()
		CreateRecord(repo, voucher)
		voucher.Number = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", voucher.Model.ID)
		fmt.Printf("code : %v\n", voucher.Number)
		fmt.Printf("prev version : %v\n", voucher.Version)
		UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)

		voucher, _ = ReadRecord[models.Voucher](repo, voucher.Model.ID, "voucher")
		voucher.Number = randgenerator.GenerateRandomCode()
		err := UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)

		fmt.Printf("new version : %v\n", voucher.Version)
		assert.NoError(t, err, "expected no error")

	})
}

func createTempVoucher() *models.Voucher {

	number := randgenerator.GenerateRandomCode()
	var tempDetailedId uint = 2
	var tempSubsidiaryId uint = 4

	temp := make([]*models.VoucherItem, 2)
	temp[0] = &models.VoucherItem{DetailedId: tempDetailedId, SubsidiaryId: tempSubsidiaryId}
	temp[1] = &models.VoucherItem{DetailedId: tempDetailedId, SubsidiaryId: tempSubsidiaryId}

	return &models.Voucher{Number: number, VoucherItems: temp}
}

func createTempVoucherItem() *models.VoucherItem {
	var tempDetailedId uint = 2
	var tempSubsidiaryId uint = 4

	return &models.VoucherItem{DetailedId: tempDetailedId, SubsidiaryId: tempSubsidiaryId}
}

func createTempSubsidiary() *models.Subsidiary {
	code := randgenerator.GenerateRandomCode()
	title := randgenerator.GenerateRandomTitle()
	return &models.Subsidiary{Code: code, Title: title, HasDetailed: false}
}

func createTempDetailed() *models.Detailed {
	code := randgenerator.GenerateRandomCode()
	title := randgenerator.GenerateRandomTitle()
	return &models.Detailed{Code: code, Title: title}
}

func TestDeleteDetailed(t *testing.T) {
	repo, err := createConnectionForTest()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("delete detailed record seccessfully", func(t *testing.T) {
		detailed := createTempDetailed()
		CreateRecord(repo, detailed)

		DeleteDetailedRecord(repo, detailed)

		result := repo.AccountingDB.First(&detailed)
		assert.Error(t, result.Error, "expected error indicate detailed record not found")

	})

	t.Run("deletion detailed record fail because record does not exist in database", func(t *testing.T) {
		detailed := createTempDetailed()
		DeleteDetailedRecord(repo, detailed)
		detailed.Model.ID = 1_000_000
		result := repo.AccountingDB.First(&detailed)
		assert.Error(t, result.Error, "expected error indicate detailed record not found")
	})

	t.Run("deletion detailed record fails because there is a reffrence in some voucher items  ", func(t *testing.T) {
		detailed := createTempDetailed()
		CreateRecord(repo, detailed)
		voucher := &models.Voucher{Number: randgenerator.GenerateRandomCode(), VoucherItems: []*models.VoucherItem{createTempVoucherItem()}}
		voucher.VoucherItems[0].DetailedId = detailed.Model.ID
		CreateRecord(repo, voucher)
		fmt.Printf("det : %v", detailed.Model.ID)
		fmt.Printf("vi : %v", voucher.VoucherItems[0].Model.ID)
		err := DeleteDetailedRecord(repo, detailed)

		assert.Error(t, err, "expected error indicate violation forignkey constraint")

	})

	t.Run("can not delete detailed record if versions were  different", func(t *testing.T) {
		detailed := createTempDetailed()
		CreateRecord(repo, detailed)
		detailed.Code = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", detailed.Model.ID)
		fmt.Printf("code : %v\n", detailed.Code)
		fmt.Printf("prev version : %v\n", detailed.Version)
		UpdateDetailed(repo, detailed, detailed.Model.ID)
		err := DeleteDetailedRecord(repo, detailed)
		fmt.Printf("new version : %v\n", detailed.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can delete detailed record if versions were same", func(t *testing.T) {
		detailed := createTempDetailed()
		CreateRecord(repo, detailed)
		detailed.Code = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", detailed.Model.ID)
		fmt.Printf("code : %v\n", detailed.Code)
		fmt.Printf("prev version : %v\n", detailed.Version)
		UpdateDetailed(repo, detailed, detailed.Model.ID)
		detailed, _ = ReadRecord[models.Detailed](repo, detailed.Model.ID, "detailed")

		err := DeleteDetailedRecord(repo, detailed)
		fmt.Printf("new version : %v\n", detailed.Version)
		assert.NoError(t, err, "expected no error")

	})
}

func TestDeleteSubsidiary(t *testing.T) {
	repo, err := createConnectionForTest()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("delete subsidiary record seccessfully", func(t *testing.T) {
		subsidiary := createTempSubsidiary()
		CreateRecord(repo, subsidiary)

		DeleteSubsidiaryRecord(repo, subsidiary)

		result := repo.AccountingDB.First(&subsidiary)
		assert.Error(t, result.Error, "expected error indicate subsiduary record not found")

	})

	t.Run("deletion subsidiary record fail because record does not exist in database", func(t *testing.T) {
		subsidiary := createTempSubsidiary()
		DeleteSubsidiaryRecord(repo, subsidiary)
		subsidiary.Model.ID = 1_000_000
		result := repo.AccountingDB.First(&subsidiary)
		assert.Error(t, result.Error, "expected error indicate subsiduary record not found")
	})

	t.Run("deletion subsidiary record fails because there is a reffrence in some voucher items  ", func(t *testing.T) {
		subsidiary := createTempSubsidiary()
		CreateRecord(repo, subsidiary)
		voucher := &models.Voucher{Number: randgenerator.GenerateRandomCode(), VoucherItems: []*models.VoucherItem{createTempVoucherItem()}}
		voucher.VoucherItems[0].SubsidiaryId = subsidiary.Model.ID
		CreateRecord(repo, voucher)
		fmt.Printf("det : %v", subsidiary.Model.ID)
		fmt.Printf("vi : %v", voucher.VoucherItems[0].Model.ID)
		err := DeleteSubsidiaryRecord(repo, subsidiary)

		assert.Error(t, err, "expected error indicate violation forignkey constraint")

	})

	t.Run("can not delete subsidiary record if versions were  different", func(t *testing.T) {
		subsidiary := createTempSubsidiary()
		CreateRecord(repo, subsidiary)
		subsidiary.Code = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
		fmt.Printf("code : %v\n", subsidiary.Code)
		fmt.Printf("prev version : %v\n", subsidiary.Version)
		UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
		err := DeleteSubsidiaryRecord(repo, subsidiary)
		fmt.Printf("new version : %v\n", subsidiary.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can delete subsidiary record if versions were same", func(t *testing.T) {
		subsidiary := createTempSubsidiary()
		CreateRecord(repo, subsidiary)
		subsidiary.Code = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
		fmt.Printf("code : %v\n", subsidiary.Code)
		fmt.Printf("prev version : %v\n", subsidiary.Version)
		UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
		subsidiary, _ = ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID, "subsidiary")

		err := DeleteSubsidiaryRecord(repo, subsidiary)
		fmt.Printf("new version : %v\n", subsidiary.Version)
		assert.NoError(t, err, "expected no error")

	})
}

func TestDeleteVoucher(t *testing.T) {
	repo, err := createConnectionForTest()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("delete voucher record seccessfully", func(t *testing.T) {

		voucher := createTempVoucher()
		CreateRecord(repo, voucher)
		fmt.Printf("voucher : %v", voucher.Model.ID)
		err := DeleteVoucherRecord(repo, voucher)
		assert.NoError(t, err, "expected no error %v", err)
		result := repo.AccountingDB.First(&voucher)
		assert.Error(t, result.Error, "expected error indicate voucher record not found")

	})

	t.Run("deletion voucher record fail because record does not exist in database", func(t *testing.T) {
		voucher := createTempVoucher()
		DeleteVoucherRecord(repo, voucher)
		voucher.Model.ID = 1_000_000
		result := repo.AccountingDB.First(&voucher)
		assert.Error(t, result.Error, "expected error indicate voucher record not found")
	})

	t.Run("can not delete voucher record if versions were  different", func(t *testing.T) {
		voucher := createTempVoucher()
		CreateRecord(repo, voucher)
		voucher.Number = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", voucher.Model.ID)
		fmt.Printf("code : %v\n", voucher.Number)
		fmt.Printf("prev version : %v\n", voucher.Version)
		UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)

		err := DeleteVoucherRecord(repo, voucher)
		fmt.Printf("new version : %v\n", voucher.Version)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can delete voucher record if versions were same", func(t *testing.T) {
		voucher := createTempVoucher()
		CreateRecord(repo, voucher)
		voucher.Number = randgenerator.GenerateRandomCode()
		fmt.Printf("prev id : %v\n", voucher.Model.ID)
		fmt.Printf("code : %v\n", voucher.Number)
		fmt.Printf("prev version : %v\n", voucher.Version)
		UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)

		voucher, _ = ReadRecord[models.Voucher](repo, voucher.Model.ID, "voucher")

		err := DeleteVoucherRecord(repo, voucher)
		fmt.Printf("new version : %v\n", voucher.Version)
		assert.NoError(t, err, "expected no error")

	})
}

func TestReadRecord(t *testing.T) {
	repo, err := createConnectionForTest()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("can read the detailed record successfully", func(t *testing.T) {
		detailed := createTempDetailed()
		CreateRecord(repo, detailed)

		res, err := ReadRecord[models.Detailed](repo, detailed.Model.ID, "detailed")
		assert.NoError(t, err, "expected no error")
		assert.Equal(t, res.Code, detailed.Code)
		assert.Equal(t, res.Title, detailed.Title)
	})

	t.Run("return error when the detailed record is not in database ", func(t *testing.T) {

		_, err := ReadRecord[models.Detailed](repo, 1_000_000, "detailed")
		assert.Error(t, err, "expected  error indicate can not found the detailed record")

	})

	t.Run("can read the subsidiary record successfully", func(t *testing.T) {
		subsidiary := createTempSubsidiary()
		CreateRecord(repo, subsidiary)

		res, err := ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID, "subsidiary")
		assert.NoError(t, err, "expected no error")
		assert.Equal(t, res.Code, subsidiary.Code)
		assert.Equal(t, res.Title, subsidiary.Title)
		assert.Equal(t, res.HasDetailed, subsidiary.HasDetailed)
	})

	t.Run("return error when the subsidiary record is not in database ", func(t *testing.T) {

		_, err := ReadRecord[models.Subsidiary](repo, 1_000_000, "subsidiary")
		assert.Error(t, err, "expected  error indicate can not found the subsidiary record")

	})

	t.Run("can read the voucher record successfully", func(t *testing.T) {
		voucher := createTempVoucher()
		CreateRecord(repo, voucher)

		res, err := ReadRecord[models.Voucher](repo, voucher.Model.ID, "voucher")
		assert.NoError(t, err, "expected no error")
		assert.Equal(t, res.Number, res.Number)

	})

	t.Run("return error when the voucher record is not in database ", func(t *testing.T) {

		_, err := ReadRecord[models.Voucher](repo, 1_000_000, "voucher")
		assert.Error(t, err, "expected  error indicate can not found the voucher record")

	})

}
