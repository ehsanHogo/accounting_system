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
		code := randgenerator.GenerateRandomCode()
		temp := make([]*models.VoucherItem, 2)
		temp[0] = &models.VoucherItem{Credit: 1,  DetailedId: 2}
		temp[1] = &models.VoucherItem{Credit: 2,  DetailedId: 2}
		voucher := &models.Voucher{Number: code, VoucherItems: temp}
		err := CreateRecord(repo, voucher)

		assert.NoError(t, err, "expected voucher record to be created, but got error")
		var result models.Voucher
		err = repo.AccountingDB.First(&result, voucher.Model.ID).Error //Number is uniqe
		assert.NoError(t, err, " can not find the inserted voucher record :")

	})

	t.Run("the voucher record creation fail because duplication number", func(t *testing.T) {

		code := randgenerator.GenerateRandomCode()

		temp := make([]*models.VoucherItem, 2)
		temp[0] = &models.VoucherItem{Credit: 1,  DetailedId: 2}
		temp[1] = &models.VoucherItem{Credit: 2,   DetailedId: 2}
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
		code := randgenerator.GenerateRandomCode()
		temp := make([]*models.VoucherItem, 2)
		temp[0] = &models.VoucherItem{Credit: 11,  DetailedId: 2}
		temp[1] = &models.VoucherItem{ DetailedId: 2}
		voucher := &models.Voucher{Number: code, VoucherItems: temp}
		CreateRecord(repo, voucher)
		fmt.Printf("prev Code %v\n", code)
		prevVoucherId := voucher.Model.ID
		code = randgenerator.GenerateRandomCode()
		temp = append(temp, &models.VoucherItem{Credit: 13,  DetailedId: 2})
		temp[1].Credit = 12

		fmt.Printf("new Code %v\n", code)
		voucher = &models.Voucher{Number: code, VoucherItems: temp}
		err := UpdateVoucher(repo, voucher, []*models.VoucherItem{temp[1]}, []*models.VoucherItem{temp[0]}, []*models.VoucherItem{temp[2]}, prevVoucherId)
		assert.NoError(t, err, "expected no error")
	})

	t.Run("return error when update voucher record that is not in databse", func(t *testing.T) {
		code := randgenerator.GenerateRandomCode()
		temp := make([]*models.VoucherItem, 2)
		temp[0] = &models.VoucherItem{ DetailedId: 2}
		temp[1] = &models.VoucherItem{ DetailedId: 2}
		voucher := &models.Voucher{Number: code, VoucherItems: temp}

		err := UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, 1_000_000)
		assert.Error(t, err, "expected error indicate there is such id in database")

	})
}
