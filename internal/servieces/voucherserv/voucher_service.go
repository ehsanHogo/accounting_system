package voucherserv

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/validations"
	"fmt"

	"gorm.io/gorm"
)

func InsertVoucher(db *gorm.DB, d *models.Voucher) error {

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

func UpdateVoucher(db *gorm.DB, d *models.Voucher, updatedItem []*models.VoucherItem, deletedItem []*models.VoucherItem, insertedItem []*models.VoucherItem) error {

	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("can not begin transaction: %v", tx.Error)
	}

	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		}
	}()

	err := validations.UpdateVoucherValidation(tx, d, updatedItem, deletedItem, insertedItem)

	if err != nil {
		return fmt.Errorf("can not update voucher due to validation failure: %v", err)
	}

	newV := &models.Voucher{Number: d.Number}

	for _, vi := range deletedItem {

		err := repositories.DeleteRecord(tx, vi)
		if err != nil {
			return fmt.Errorf("can not delete voucher item : %w", err)
		}
	}

	for _, vi := range updatedItem {
		err := UpdateVoucherItem(tx, vi, vi.ID)
		if err != nil {
			return fmt.Errorf("can not update voucher item : %w", err)
		}
	}

	for _, vi := range insertedItem {

		vi.VoucherID = d.ID
		err := repositories.CreateRecord(tx, vi)

		if err != nil {
			return fmt.Errorf("can not insert voucher item : %w", err)
		}
	}

	err = repositories.UpdateRecord[models.Voucher](tx, newV, d.ID)
	if err != nil {
		return fmt.Errorf("can not update voucher due to database operation failure : %v", err)
	} else {

		err = tx.Commit().Error
		if err != nil {
			return fmt.Errorf("can not commit transaction: %v", err)
		}

		return nil
	}

}

func DeleteVoucher(db *gorm.DB, d *models.Voucher) error {

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

func ReadVoucher(db *gorm.DB, id uint) (*models.Voucher, error) {

	res, err := repositories.ReadRecord[models.Voucher](db, id)
	if err != nil {
		return nil, fmt.Errorf("can not read voucher due to database operation failure : %v", err)
	} else {

		return res, nil
	}
}

func ReadVoucherItem(db *gorm.DB, id uint) (*models.VoucherItem, error) {

	res, err := repositories.ReadRecord[models.VoucherItem](db, id)
	if err != nil {
		return nil, fmt.Errorf("can not read voucher item due to database operation failure : %v", err)
	} else {

		return res, nil
	}
}

func UpdateVoucherItem(db *gorm.DB, v *models.VoucherItem, id uint) error {
	var newV models.VoucherItem
	if err := db.First(&newV, id).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	newV.Credit = v.Credit
	newV.Debit = v.Debit
	newV.DetailedId = v.DetailedId
	newV.SubsidiaryId = v.SubsidiaryId

	if err := db.Model(&newV).Where("id = ?", v.ID).Updates(newV).Error; err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}
