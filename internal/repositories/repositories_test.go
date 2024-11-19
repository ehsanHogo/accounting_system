package repositories

import (
	"accounting_system/config"
	"accounting_system/internal/models"
	randgenerator "accounting_system/internal/utils"
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
		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		detailed := &models.Detailed{Code: code, Title: title}
		err := CreateRecord(repo, detailed)
		fmt.Printf("detailed id : %v", detailed.Model.ID)
		assert.NoError(t, err, "expected detailed record to be created, but got error")
		var result models.Detailed
		err = repo.AccountingDB.First(&result, detailed.Model.ID).Error //Code is uniqe
		assert.NoError(t, err, " can not find the inserted detailed record :")

	})

	t.Run("the detailed record creation fail because duplication code", func(t *testing.T) {

		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		detailed := &models.Detailed{Code: code, Title: title}
		err := CreateRecord(repo, detailed)
		assert.NoError(t, err, "expected detailed record to be created, but got error")

		title = randgenerator.GenerateRandomTitle()
		detailed = &models.Detailed{Code: code, Title: title}
		err = CreateRecord(repo, detailed)

		assert.Error(t, err, "expected getting duplicate detailed code error")

	})

	t.Run("the detailed record creation fail because duplication title", func(t *testing.T) {

		code := randgenerator.GenerateRandomCode()
		title := randgenerator.GenerateRandomTitle()
		detailed := &models.Detailed{Code: code, Title: title}
		err := CreateRecord(repo, detailed)
		assert.NoError(t, err, "expected detailed record to be created, but got error")

		code = randgenerator.GenerateRandomCode()
		detailed = &models.Detailed{Code: code, Title: title}
		err = CreateRecord(repo, detailed)

		assert.Error(t, err, "expected getting duplicate detailed title error")

	})

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

		DeleteRecord(repo, detailed)

		result := repo.AccountingDB.First(&detailed)
		assert.Error(t, result.Error, "expected error indicate detailed record not found")

	})

	t.Run("deletion detailed record fail because record does not exist in database", func(t *testing.T) {
		detailed := createTempDetailed()
		DeleteRecord(repo, detailed)
		detailed.Model.ID = 1_000_000
		result := repo.AccountingDB.First(&detailed)
		assert.Error(t, result.Error, "expected error indicate detailed record not found")
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

		DeleteRecord(repo, subsidiary)

		result := repo.AccountingDB.First(&subsidiary)
		assert.Error(t, result.Error, "expected error indicate subsiduary record not found")

	})

	t.Run("deletion subsidiary record fail because record does not exist in database", func(t *testing.T) {
		subsidiary := createTempSubsidiary()
		DeleteRecord(repo, subsidiary)
		subsidiary.Model.ID = 1_000_000
		result := repo.AccountingDB.First(&subsidiary)
		assert.Error(t, result.Error, "expected error indicate subsiduary record not found")
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

		DeleteRecord(repo, voucher)

		result := repo.AccountingDB.First(&voucher)
		assert.Error(t, result.Error, "expected error indicate voucher record not found")

	})

	t.Run("deletion voucher record fail because record does not exist in database", func(t *testing.T) {
		voucher := createTempVoucher()
		DeleteRecord(repo, voucher)
		voucher.Model.ID = 1_000_000
		result := repo.AccountingDB.First(&voucher)
		assert.Error(t, result.Error, "expected error indicate voucher record not found")
	})
}
