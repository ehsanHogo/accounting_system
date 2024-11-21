package servieces

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/validations"
	"fmt"
)

func InsertDetailed(db *repositories.Repositories, d *models.Detailed) error {

	err := validations.DetailedValidation(d)

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
	err := validations.DetailedValidation(d)

	if err != nil {
		return fmt.Errorf("can not update detailed due to validation failure : %v", err)
	}

	err = repositories.UpdateDetailed(db, d, d.Model.ID)
	if err != nil {
		return fmt.Errorf("can not update detailed due to database operation failure: %v", err)
	} else {

		return nil
	}

}

func DeleteDetailed(db *repositories.Repositories, d *models.Detailed) error {

	err := repositories.DeleteDetailedRecord(db, d)
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

	err := validations.SubsidiaryValidation(d)

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
	err := validations.SubsidiaryValidation(d)

	if err != nil {
		return fmt.Errorf("can not update subsidiary due to validation failure : %v", err)
	}

	err = repositories.UpdateSubsidiary(db, d, d.Model.ID)
	if err != nil {
		return fmt.Errorf("can not update subsidiary due to database operation failure: %v", err)
	} else {

		return nil
	}

}

func DeleteSubsidiary(db *repositories.Repositories, d *models.Subsidiary) error {

	err := repositories.DeleteSubsidiaryRecord(db, d)
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

	err := validations.InsertVoucherValidation(d)

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

	err = repositories.UpdateVoucher(db, d, updatedItem, deletedItem, insertedItem, d.Model.ID)
	if err != nil {
		return fmt.Errorf("can not update voucher due to database operation failure : %v", err)
	} else {

		return nil
	}

}

func DeleteVoucher(db *repositories.Repositories, d *models.Voucher) error {

	err := repositories.DeleteVoucherRecord(db, d)
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
