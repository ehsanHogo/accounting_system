package voucherserv

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/validations"
	"fmt"
)

func InsertVoucher(db *repositories.Repositories, d *models.Voucher) error {

	err := validations.InsertVoucherValidation(db, d)

	if err != nil {
		return fmt.Errorf("can not insert voucher due to validation failure: %v", err)
	}

	err = repositories.CreateRecord(db, d)
	if err != nil {
		return fmt.Errorf("can not insert voucher due to database operation failure : %v", err)
	}
	return nil
}

func UpdateVoucher(db *repositories.Repositories, d *models.Voucher, updatedItem []*models.VoucherItem, deletedItem []*models.VoucherItem, insertedItem []*models.VoucherItem) error {
	err := validations.UpdateVoucherValidation(db, d, updatedItem, deletedItem, insertedItem)

	if err != nil {
		return fmt.Errorf("can not update voucher due to validation failure: %v", err)
	}

	newV := &models.Voucher{Number: d.Number, Version: d.Version + 1}
	fmt.Println("here newV")
	fmt.Println(newV.Number)
	for _, vi := range deletedItem {

		err := repositories.DeleteRecord(db, vi)
		if err != nil {
			return fmt.Errorf("can not delete voucher item : %w", err)
		}
	}

	for _, vi := range updatedItem {
		err := UpdateVoucherItem(db, vi, vi.Model.ID)
		if err != nil {
			return fmt.Errorf("can not update voucher item : %w", err)
		}
	}

	for _, vi := range insertedItem {

		vi.VoucherID = d.Model.ID
		err := repositories.CreateRecord(db, vi)

		if err != nil {
			return fmt.Errorf("can not insert voucher item : %w", err)
		}
	}

	err = repositories.UpdateRecord[models.Voucher](db, newV, d.Model.ID)
	if err != nil {
		return fmt.Errorf("can not update voucher due to database operation failure : %v", err)
	} else {

		return nil
	}

}

func DeleteVoucher(db *repositories.Repositories, d *models.Voucher) error {

	err := validations.DeleteVoucherValidation(db, d)

	if err != nil {
		return fmt.Errorf("can not delete voucher due to validation failure: %v", err)
	}

	err = repositories.DeleteRecord[models.Voucher](db, d)
	if err != nil {
		return fmt.Errorf("can not delete voucher due to database operation failure : %v", err)
	} else {

		return nil
	}

}

func ReadVoucher(db *repositories.Repositories, id uint) (*models.Voucher, error) {

	res, err := repositories.ReadRecord[models.Voucher](db, id)
	if err != nil {
		return nil, fmt.Errorf("can not read voucher due to database operation failure : %v", err)
	} else {

		return res, nil
	}
}

func UpdateVoucherItem(db *repositories.Repositories, v *models.VoucherItem, id uint) error {
	var newV models.VoucherItem
	if err := db.AccountingDB.First(&newV, id).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	newV.Credit = v.Credit
	newV.Debit = v.Debit
	newV.DetailedId = v.DetailedId
	newV.SubsidiaryId = v.SubsidiaryId

	if err := db.AccountingDB.Model(&newV).Where("id = ?", v.Model.ID).Updates(newV).Error; err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}
