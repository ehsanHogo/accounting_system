package servieces

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/validations"

	"fmt"
)

func InsertDetailed(db *repositories.Repositories, d *models.Detailed) error {

	err := validations.InsertDetailedValidation(d)

	if err != nil {
		return fmt.Errorf("can not insert detailed due to validation failure : %v", err)
	}

	err = repositories.CreateRecord(db, d)
	if err != nil {
		return fmt.Errorf("can not insert detailed due to database operation failure: %v", err)
	} else {

		return nil
	}

}

func UpdateDetailed(db *repositories.Repositories, d *models.Detailed) error {
	err := validations.UpdateDetailedValidation(db, d)

	if err != nil {
		return fmt.Errorf("can not update detailed due to validation failure : %v", err)
	}

	newV := &models.Detailed{
		Code:    d.Code,
		Title:   d.Title,
		Version: d.Version + 1}

	err = repositories.UpdateRecord[models.Detailed](db, newV, d.Model.ID)
	if err != nil {
		return fmt.Errorf("can not update detailed due to database operation failure: %v", err)
	} else {

		return nil
	}

}

func DeleteDetailed(db *repositories.Repositories, d *models.Detailed) error {
	err := validations.DeleteDetailedValidation(db, d)

	if err != nil {
		return fmt.Errorf("can not delete detailed due to validation failure : %v", err)
	}

	err = repositories.DeleteRecord[models.Detailed](db, d)

	if err != nil {
		return fmt.Errorf("can not delete detailed due to database operation failure: %v", err)
	} else {

		return nil
	}

}

func ReadDetailed(db *repositories.Repositories, id uint) (*models.Detailed, error) {

	res, err := repositories.ReadRecord[models.Detailed](db, id)
	if err != nil {
		return nil, fmt.Errorf("can not read detailed due to database operation failure : %v", err)
	} else {

		return res, nil
	}
}

func InsertSubsidiary(db *repositories.Repositories, d *models.Subsidiary) error {

	err := validations.InsertSubsidiaryValidation(d)

	if err != nil {
		return fmt.Errorf("can not insert subsidiary due to validation failure : %v", err)
	}

	err = repositories.CreateRecord(db, d)
	if err != nil {
		return fmt.Errorf("can not insert subsidiary due to database operation failure: %v", err)
	} else {

		return nil
	}

}

func UpdateSubsidiary(db *repositories.Repositories, d *models.Subsidiary) error {
	err := validations.UpdateSubsidiaryValidation(db, d)

	if err != nil {
		return fmt.Errorf("can not update subsidiary due to validation failure : %v", err)
	}

	newV := &models.Subsidiary{
		Code:        d.Code,
		Title:       d.Title,
		HasDetailed: d.HasDetailed,
		Version:     d.Version + 1}

	err = repositories.UpdateRecord[models.Subsidiary](db, newV, d.Model.ID)
	if err != nil {
		return fmt.Errorf("can not update subsidiary due to database operation failure: %v", err)
	} else {

		return nil
	}

}

func DeleteSubsidiary(db *repositories.Repositories, d *models.Subsidiary) error {

	err := validations.DeleteSubsidiaryValidation(db, d)

	if err != nil {
		return fmt.Errorf("can not delete subsidiary due to validation failure : %v", err)
	}
	err = repositories.DeleteRecord[models.Subsidiary](db, d)
	if err != nil {
		return fmt.Errorf("can not delete subsidiary due to database operation failure : %v", err)
	} else {

		return nil
	}

}

func ReadSubsidiary(db *repositories.Repositories, id uint) (*models.Subsidiary, error) {

	res, err := repositories.ReadRecord[models.Subsidiary](db, id)
	if err != nil {
		return nil, fmt.Errorf("can not read subsidiary due to database operation failure : %v", err)
	} else {

		return res, nil
	}
}

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
