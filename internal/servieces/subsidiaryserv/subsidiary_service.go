package subsidiaryserv

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/validations"
	"fmt"

	"gorm.io/gorm"
)

func InsertSubsidiary(db *gorm.DB, d *models.Subsidiary) error {

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

func UpdateSubsidiary(db *gorm.DB, d *models.Subsidiary) error {
	err := validations.UpdateSubsidiaryValidation(db, d)

	if err != nil {
		return fmt.Errorf("can not update subsidiary due to validation failure : %v", err)
	}

	newV := &models.Subsidiary{
		Code:        d.Code,
		Title:       d.Title,
		HasDetailed: d.HasDetailed,
	}

	err = repositories.UpdateRecord[models.Subsidiary](db, newV, d.ID)
	if err != nil {
		return fmt.Errorf("can not update subsidiary due to database operation failure: %v", err)
	} else {

		return nil
	}

}

func DeleteSubsidiary(db *gorm.DB, d *models.Subsidiary) error {

	err := validations.DeleteSubsidiaryValidation(db, d)

	if err != nil {
		return fmt.Errorf("can not delete subsidiary due to validation failure : %v", err)
	}
	err = repositories.DeleteRecord[models.Subsidiary](db, d.ID)
	if err != nil {
		return fmt.Errorf("can not delete subsidiary due to database operation failure : %v", err)
	} else {

		return nil
	}

}

func ReadSubsidiary(db *gorm.DB, id int64) (*models.Subsidiary, error) {

	res, err := repositories.ReadRecord[models.Subsidiary](db, id)
	if err != nil {
		return nil, fmt.Errorf("can not read subsidiary due to database operation failure : %v", err)
	} else {

		return res, nil
	}
}
