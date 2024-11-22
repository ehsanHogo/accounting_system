package detailedserv

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/validations"
	"fmt"

	"gorm.io/gorm"
)

func InsertDetailed(db *gorm.DB, d *models.Detailed) error {

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

func UpdateDetailed(db *gorm.DB, d *models.Detailed) error {
	err := validations.UpdateDetailedValidation(db, d)

	if err != nil {
		return fmt.Errorf("can not update detailed due to validation failure : %v", err)
	}

	newV := &models.Detailed{
		Code:  d.Code,
		Title: d.Title,
	}

	err = repositories.UpdateRecord[models.Detailed](db, newV, d.ID)
	if err != nil {
		return fmt.Errorf("can not update detailed due to database operation failure: %v", err)
	} else {

		return nil
	}

}

func DeleteDetailed(db *gorm.DB, d *models.Detailed) error {
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

func ReadDetailed(db *gorm.DB, id int64) (*models.Detailed, error) {

	res, err := repositories.ReadRecord[models.Detailed](db, id)
	if err != nil {
		return nil, fmt.Errorf("can not read detailed due to database operation failure : %v", err)
	} else {

		return res, nil
	}
}
