package detailedserv

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


