package voucherserv

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/servieces/subsidiaryserv"
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

func TestInsertVoucher(t *testing.T) {
	repo := SetupTestRepo(t)

	t.Run("can insert voucher successfully", func(t *testing.T) {
		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "expected no error when inserting voucher")

		var result models.Voucher
		err = repo.AccountingDB.First(&result, voucher.ID).Error
		assert.NoError(t, err, " can not find the inserted voucher record :")
	})

	t.Run("can not insert voucher because duplication number", func(t *testing.T) {
		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)

		assert.NoError(t, err, "expected no error while inserting voucher ")

		err = InsertVoucher(repo.AccountingDB, voucher)

		assert.Error(t, err, "expected getting duplicate voucher number error")

	})

	t.Run("can not insert voucher record with empty number", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 100}}}

		err = InsertVoucher(repo.AccountingDB, voucher)

		assert.Error(t, err, "expected error indicate empty title not allowed")
	})

	t.Run("can not insert voucher record with number length greater than 64", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 100}}}

		s := "1"
		for i := 0; i < validations.MaxCodeLength; i++ {
			voucher.Number += s
		}
		err = InsertVoucher(repo.AccountingDB, voucher)

		assert.Error(t, err, "expected error indicate number length should not be greater than 64 ")
	})

	t.Run("can not insert voucher record with imbalance voucher items", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 50}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 100}}}

		err = InsertVoucher(repo.AccountingDB, voucher)

		assert.Error(t, err, "expected error indicate voucher items should be balance ")
	})

	t.Run("can not insert voucher record with unvalied voucher items credits od debits", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: -100}}}
		err = InsertVoucher(repo.AccountingDB, voucher)
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")

		voucher = &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID}}}
		err = InsertVoucher(repo.AccountingDB, voucher)
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")

		voucher = &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100, Debit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100, Debit: 100}}}
		err = InsertVoucher(repo.AccountingDB, voucher)
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")
	})

	t.Run("can not insert voucher record with invalied voucher items length ", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		tempList := make([]*models.VoucherItem, 0)

		for i := 0; i < 502; i += 2 {
			tempList = append(tempList, &models.VoucherItem{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100})

			tempList = append(tempList, &models.VoucherItem{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 100})
		}
		voucher := &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: tempList}

		err = InsertVoucher(repo.AccountingDB, voucher)

		assert.Error(t, err, "expected error indicate voucher items length should not be greater than 500 ")

		voucher = &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{}}

		err = InsertVoucher(repo.AccountingDB, voucher)

		assert.Error(t, err, "expected error indicate voucher items length should not be less than 2 ")

	})

	t.Run("can not insert voucher record when voucher item subsidiary has detailed and not specifying detailed", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary := &models.Subsidiary{Code: repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code"), Title: repositories.GenerateUniqeTitle[models.Subsidiary](repo.AccountingDB), HasDetailed: true}
		err = subsidiaryserv.InsertSubsidiary(repo.AccountingDB, subsidiary)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, Credit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 100}}}

		err = InsertVoucher(repo.AccountingDB, voucher)

		assert.Error(t, err, "expected error indicate voucher items should have detailed")
	})

	t.Run("can not insert voucher record when voucher item subsidiary does not have detailed and  specifying detailed", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, Credit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 100}}}

		err = InsertVoucher(repo.AccountingDB, voucher)

		assert.Error(t, err, "expected error indicate voucher items should not have detailed")
	})

	t.Run("can not insert voucher record when voucher item does not  specify subsidiary", func(t *testing.T) {

		voucherItem1 := &models.VoucherItem{Credit: 100}
		voucherItem2 := &models.VoucherItem{Debit: 100}

		voucher := &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{voucherItem1, voucherItem2}}

		err := InsertVoucher(repo.AccountingDB, voucher)

		assert.Error(t, err, "expected error indicate voucher items should have subsidiary")
	})
}

func TestUpdateVoucher(t *testing.T) {
	repo := SetupTestRepo(t)

	t.Run("can update voucher  seccessfully", func(t *testing.T) {

		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "can not create voucher record")

		newVoucherItem, err := temporary.ReturnTempVoucherItem(repo.AccountingDB)
		assert.NoError(t, err, "can not create voucher item record")
		temp := append(voucher.VoucherItems, newVoucherItem)

		newVoucherItem, err = temporary.ReturnTempVoucherItem(repo.AccountingDB)
		assert.NoError(t, err, "can not create voucher item record")
		temp = append(temp, newVoucherItem)
		temp[len(temp)-1].Credit = 250
		temp[len(temp)-1].Debit = 0

		temp[1].Credit -= 1
		temp[2].Debit -= 1

		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{temp[1], temp[2]}, []*models.VoucherItem{temp[0], temp[3]}, []*models.VoucherItem{temp[4], temp[5]})

		assert.NoError(t, err, "expected no error while updating voucher ")
		_, err = ReadVoucherItem(repo.AccountingDB, voucher.VoucherItems[0].ID)
		assert.Error(t, err, "expected error indicate voucher item not found")

		_, err = ReadVoucherItem(repo.AccountingDB, voucher.VoucherItems[1].ID)
		assert.NoError(t, err, "expexted no error when reading the voucherItem record")

		_, err = ReadVoucherItem(repo.AccountingDB, voucher.VoucherItems[2].ID)
		assert.NoError(t, err, "expexted no error when reading the voucherItem record")

		_, err = ReadVoucherItem(repo.AccountingDB, voucher.VoucherItems[3].ID)
		assert.Error(t, err, "expexted no error when reading the voucherItem record")

		_, err = ReadVoucherItem(repo.AccountingDB, temp[4].ID)
		assert.NoError(t, err, "expexted no error when reading the voucherItem record")

		_, err = ReadVoucherItem(repo.AccountingDB, temp[5].ID)
		assert.NoError(t, err, "expexted no error when reading the voucherItem record")
	})

	t.Run("can not update voucher because duplication number", func(t *testing.T) {
		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)

		assert.NoError(t, err, "expected no error while inserting voucher ")
		prevNumber := voucher.Number

		voucher, err = temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting voucher ")

		voucher.Number = prevNumber
		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate  duplicate voucher number")

	})

	t.Run("can not update voucher because empty number", func(t *testing.T) {
		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)

		assert.NoError(t, err, "expected no error while updating voucher ")

		voucher.Number = ""
		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate  empty voucher number ")

	})

	t.Run("can not update voucher record with number length greater than 64", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		_, err = temporary.CreateTempVoucher(repo.AccountingDB)

		assert.NoError(t, err, "expected no error while updating voucher ")

		voucher := &models.Voucher{VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 100}}}

		s := "1"
		for i := 0; i < validations.MaxCodeLength+1; i++ {
			voucher.Number += s
		}
		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate number length should not be greater than 64 ")
	})

	t.Run("can not update voucher record with unvalied voucher items credits or debits", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")
		_, err = temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting voucher ")

		voucher := &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: -100}}}
		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")

		voucher = &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID}}}
		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")

		voucher = &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100, Debit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100, Debit: 100}}}
		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.Error(t, err, "expected error indicate credits or dibits is invalied ")
	})

	t.Run("can not update voucher record with invalied voucher items length ", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting voucher ")
		tempList := make([]*models.VoucherItem, 0)

		for i := 0; i < 502; i += 2 {
			tempList = append(tempList, &models.VoucherItem{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100})

			tempList = append(tempList, &models.VoucherItem{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 100})
		}

		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, tempList)

		assert.Error(t, err, "expected error indicate voucher items length should not be greater than 500 ")

		tempList = make([]*models.VoucherItem, 0)

		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, tempList, []*models.VoucherItem{})
		assert.NoError(t, err, "expected no error while deleting voucher")
		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, tempList)

		assert.Error(t, err, "expected error indicate voucher items length should not be less than 2 ")

	})

	t.Run("can not update voucher record when voucher item subsidiary has detailed and not specifying detailed", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary := &models.Subsidiary{Code: repositories.GenerateUniqeCode[models.Subsidiary](repo.AccountingDB, "code"), Title: repositories.GenerateUniqeTitle[models.Subsidiary](repo.AccountingDB), HasDetailed: true}
		err = subsidiaryserv.InsertSubsidiary(repo.AccountingDB, subsidiary)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")
		_, err = temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting voucher ")

		voucher := &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, Credit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 100}}}

		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate voucher items should have detailed")
	})

	t.Run("can not update voucher record when voucher item subsidiary does not have detailed and  specifying detailed", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		_, err = temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting voucher ")

		voucher := &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, Credit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 100}}}

		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate voucher items should not have detailed")
	})

	t.Run("return error when update voucher record that is not in databse", func(t *testing.T) {
		detailed, err := temporary.CreateTempDetailed(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting detailed ")

		subsidiary, err := temporary.CreateTempSubsidiary(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting subsidiary ")

		voucher := &models.Voucher{Number: repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number"), VoucherItems: []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 100}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: -100}}}
		voucher.ID = InvalidRecordID
		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.Error(t, err, "expected error indicate there is such id in database")

	})

	t.Run("can not update voucher record if versions were  different", func(t *testing.T) {
		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "can not create voucher record")

		voucher.Number = repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number")

		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.NoError(t, err, "can not update voucher record")

		voucher.Number = repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number")
		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can update voucher record if versions were same", func(t *testing.T) {
		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "can not create voucher record")

		voucher.Number = repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number")

		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.NoError(t, err, "can not update voucher record")

		voucher, _ = repositories.ReadRecord[models.Voucher](repo.AccountingDB, voucher.ID)
		voucher.Number = repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number")
		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})

		assert.NoError(t, err, "expected no error")

	})
}

func TestDeleteVoucher(t *testing.T) {
	repo := SetupTestRepo(t)
	t.Run("can delete voucher seccessfully", func(t *testing.T) {

		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inserting voucher")

		err = DeleteVoucher(repo.AccountingDB, voucher)
		assert.NoError(t, err, "expected no error while deleting voucher")
		result := repo.AccountingDB.First(&voucher)
		assert.Error(t, result.Error, "expected error indicate voucher record not found")
	})

	t.Run("deletion voucher record fail because record does not exist in database", func(t *testing.T) {
		voucher := &models.Voucher{}
		voucher.ID = InvalidRecordID
		err := DeleteVoucher(repo.AccountingDB, voucher)
		assert.Error(t, err, "expected error indicate there is not such record in data base")

	})

	t.Run("can not delete voucher record if versions were  different", func(t *testing.T) {
		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inseting voucher")

		voucher.Number = repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number")

		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.NoError(t, err, "expected no error while updating voucher")

		err = DeleteVoucher(repo.AccountingDB, voucher)
		assert.Error(t, err, "expected error indicate the versions are different")

	})

	t.Run("can delete voucher record if versions were same", func(t *testing.T) {
		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while inseting voucher")

		voucher.Number = repositories.GenerateUniqeCode[models.Voucher](repo.AccountingDB, "number")

		err = UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{})
		assert.NoError(t, err, "expected no error while updating voucher")
		voucher, _ = repositories.ReadRecord[models.Voucher](repo.AccountingDB, voucher.ID)

		err = DeleteVoucher(repo.AccountingDB, voucher)
		assert.NoError(t, err, "expected no error while deleting")

	})

}

func TestReadVoucher(t *testing.T) {
	repo := SetupTestRepo(t)

	t.Run("can read the voucher record successfully", func(t *testing.T) {
		voucher, err := temporary.CreateTempVoucher(repo.AccountingDB)
		assert.NoError(t, err, "expected no error while craeting voucher")

		_, err = ReadVoucher(repo.AccountingDB, voucher.ID)
		assert.NoError(t, err, "expected no error")

	})

	t.Run("return error when the voucher record is not in database ", func(t *testing.T) {

		_, err := ReadVoucher(repo.AccountingDB, InvalidRecordID)
		assert.Error(t, err, "expected  error indicate can not found the voucher record")

	})

}
