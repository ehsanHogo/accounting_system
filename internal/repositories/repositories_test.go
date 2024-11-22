package repositories

import (
	"accounting_system/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositories(t *testing.T) {
	repo, err := CreateConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("scenario : create -> read -> update -> raed -> delete -> read for detailed record", func(t *testing.T) {
		detailed := &models.Detailed{Code: GenerateUniqeCode[models.Detailed](repo, "code"), Title: GenerateUniqeTitle[models.Detailed](repo)}

		err = CreateRecord(repo, detailed)
		assert.NoError(t, err, "expected no error while creating detailed record ")

		_, err := ReadRecord[models.Detailed](repo, detailed.Model.ID)

		assert.NoError(t, err, "expected no error while reading detailed record ")

		detailed.Code = GenerateUniqeCode[models.Detailed](repo, "code")
		err = UpdateRecord(repo, detailed, detailed.Model.ID)

		assert.NoError(t, err, "expected no error while updating detailed record")

		fetchDetailed, err := ReadRecord[models.Detailed](repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading detailed record ")
		assert.Equal(t, detailed.Code, fetchDetailed.Code)

		err = DeleteRecord(repo, fetchDetailed)
		assert.NoError(t, err, "expected no error while deleting detailed record ")

		_, err = ReadRecord[models.Detailed](repo, detailed.Model.ID)
		assert.Error(t, err, "expected  error indicate detailed record not found  ")

	})
	t.Run("scenario : create -> read -> update -> raed -> delete -> read for subsidiary record", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: GenerateUniqeCode[models.Subsidiary](repo, "code"), Title: GenerateUniqeTitle[models.Subsidiary](repo)}

		err = CreateRecord(repo, subsidiary)
		assert.NoError(t, err, "expected no error while creating subsidiary record ")

		_, err := ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)

		assert.NoError(t, err, "expected no error while reading subsidiary record ")

		subsidiary.Code = GenerateUniqeCode[models.Subsidiary](repo, "code")
		err = UpdateRecord(repo, subsidiary, subsidiary.Model.ID)

		assert.NoError(t, err, "expected no error while updating subsidiary record")

		fetchSubsidiary, err := ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		assert.NoError(t, err, "expected no error while reading subsidiary record ")
		assert.Equal(t, subsidiary.Code, fetchSubsidiary.Code)

		err = DeleteRecord(repo, fetchSubsidiary)
		assert.NoError(t, err, "expected no error while deleting subsidiary record ")

		_, err = ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
		assert.Error(t, err, "expected  error indicate subsidiary record not found  ")

	})

	t.Run("scenario : create -> read -> update -> raed -> delete -> read for voucher record", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: GenerateUniqeCode[models.Subsidiary](repo, "code"), Title: GenerateUniqeTitle[models.Subsidiary](repo)}

		err = CreateRecord(repo, subsidiary)
		assert.NoError(t, err, "expected no error while creating subsidiary record ")

		voucherItem1 := &models.VoucherItem{SubsidiaryId: subsidiary.Model.ID, Debit: 250}
		voucherItem2 := &models.VoucherItem{SubsidiaryId: subsidiary.Model.ID, Credit: 250}
		voucher := &models.Voucher{Number: GenerateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{voucherItem1, voucherItem2}}

		err = CreateRecord(repo, voucher)
		assert.NoError(t, err, "expected no error while creating voucher record ")

		_, err := ReadRecord[models.Voucher](repo, voucher.Model.ID)

		assert.NoError(t, err, "expected no error while reading voucher record ")

		fetchVoucherItem1, err := ReadRecord[models.VoucherItem](repo, voucherItem1.Model.ID)
		assert.NoError(t, err, "expected no error while reading voucherItem1 record ")
		assert.Equal(t, voucher.Model.ID, fetchVoucherItem1.VoucherID)

		fetchVoucherItem2, err := ReadRecord[models.VoucherItem](repo, voucherItem2.Model.ID)
		assert.NoError(t, err, "expected no error while reading voucherItem2 record ")
		assert.Equal(t, voucher.Model.ID, fetchVoucherItem2.VoucherID)

		voucher.Number = GenerateUniqeCode[models.Voucher](repo, "number")
		err = UpdateRecord(repo, voucher, voucher.Model.ID)

		assert.NoError(t, err, "expected no error while updating voucher record")

		fetchVoucher, err := ReadRecord[models.Voucher](repo, voucher.Model.ID)
		assert.NoError(t, err, "expected no error while reading voucher record ")
		assert.Equal(t, voucher.Number, fetchVoucher.Number)

		err = DeleteRecord(repo, fetchVoucher)
		assert.NoError(t, err, "expected no error while deleting voucher record ")

		_, err = ReadRecord[models.Voucher](repo, voucher.Model.ID)
		assert.Error(t, err, "expected  error indicate voucher record not found  ")

		_, err = ReadRecord[models.VoucherItem](repo, voucherItem1.Model.ID)
		assert.Error(t, err, "expected  error indicate voucherItem1 record not found  ")

		_, err = ReadRecord[models.VoucherItem](repo, voucherItem2.Model.ID)
		assert.Error(t, err, "expected  error indicate voucherItem2 record not found  ")

	})

	t.Run("scenario : create -> read -> update -> raed -> delete -> read for voucher item record", func(t *testing.T) {
		subsidiary := &models.Subsidiary{Code: GenerateUniqeCode[models.Subsidiary](repo, "code"), Title: GenerateUniqeTitle[models.Subsidiary](repo)}

		err = CreateRecord(repo, subsidiary)
		assert.NoError(t, err, "expected no error while creating subsidiary record ")

		voucherItem1 := &models.VoucherItem{SubsidiaryId: subsidiary.Model.ID, Debit: 250}
		voucherItem2 := &models.VoucherItem{SubsidiaryId: subsidiary.Model.ID, Credit: 250}
		voucher := &models.Voucher{Number: GenerateUniqeCode[models.Voucher](repo, "number"), VoucherItems: []*models.VoucherItem{voucherItem1, voucherItem2}}

		err = CreateRecord(repo, voucher)
		assert.NoError(t, err, "expected no error while creating voucher record ")

		_, err := ReadRecord[models.Voucher](repo, voucher.Model.ID)

		assert.NoError(t, err, "expected no error while reading voucher record ")

		fetchVoucherItem1, err := ReadRecord[models.VoucherItem](repo, voucherItem1.Model.ID)
		assert.NoError(t, err, "expected no error while reading voucherItem1 record ")
		assert.Equal(t, voucher.Model.ID, fetchVoucherItem1.VoucherID)

		fetchVoucherItem2, err := ReadRecord[models.VoucherItem](repo, voucherItem2.Model.ID)
		assert.NoError(t, err, "expected no error while reading voucherItem2 record ")
		assert.Equal(t, voucher.Model.ID, fetchVoucherItem2.VoucherID)

		fetchVoucherItem1.Debit -= 100
		fetchVoucherItem2.Credit -= 100
		err = UpdateRecord(repo, fetchVoucherItem1, fetchVoucherItem1.Model.ID)

		assert.NoError(t, err, "expected no error while updating voucherItem1 record")

		err = UpdateRecord(repo, fetchVoucherItem2, fetchVoucherItem2.Model.ID)

		assert.NoError(t, err, "expected no error while updating voucherItem2 record")

		updatedVoucherItem1, err := ReadRecord[models.VoucherItem](repo, voucherItem1.Model.ID)
		assert.NoError(t, err, "expected no error while reading voucherItem1 record ")
		assert.Equal(t, updatedVoucherItem1.Debit, fetchVoucherItem1.Debit)

		updatedVoucherItem2, err := ReadRecord[models.VoucherItem](repo, voucherItem2.Model.ID)
		assert.NoError(t, err, "expected no error while reading voucherItem2 record ")
		assert.Equal(t, updatedVoucherItem2.Credit, fetchVoucherItem2.Credit)

		err = DeleteRecord(repo, updatedVoucherItem1)
		assert.NoError(t, err, "expected no error while deleting voucherItem1 record ")

		_, err = ReadRecord[models.VoucherItem](repo, updatedVoucherItem1.Model.ID)
		assert.Error(t, err, "expected  error indicate voucherItem1 record not found  ")

		err = DeleteRecord(repo, updatedVoucherItem2)
		assert.NoError(t, err, "expected no error while deleting voucherItem2 record ")

		_, err = ReadRecord[models.VoucherItem](repo, updatedVoucherItem2.Model.ID)
		assert.Error(t, err, "expected  error indicate voucherItem2 record not found  ")

	})
}
