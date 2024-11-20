package servieces

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/validations"
	"fmt"
)

func InsertDetailed(db *repositories.Repositories, d *models.Detailed) error {

	err := validations.ChackCodeValidation(d.Code)

	if err != nil {
		return fmt.Errorf("code validation fail due to : %v", err)
	}

	err = validations.CheckTitleValidaion(d.Title)

	if err != nil {
		return fmt.Errorf("title validation fail due to : %v", err)
	}

	repositories.CreateRecord(db, d)

	return nil
}

func UpdateDetailed(db *repositories.Repositories, d *models.Detailed) error {
	err := validations.ChackCodeValidation(d.Code)

	if err != nil {
		return fmt.Errorf("code validation fail due to : %v", err)
	}

	err = validations.CheckTitleValidaion(d.Title)

	if err != nil {
		return fmt.Errorf("title validation fail due to : %v", err)
	}

	err = repositories.UpdateDetailed(db, d, d.Model.ID)
	if err != nil {
		return fmt.Errorf("can not update : %v", err)
	} else {

		return nil
	}

}

func DeleteDetailed(db *repositories.Repositories, d *models.Detailed) error {

	err := repositories.DeleteDetailedRecord(db, d)
	if err != nil {
		return fmt.Errorf("can not delete : %v", err)
	} else {

		return nil
	}

}

func ReadDetailed(db *repositories.Repositories, d *models.Detailed) (*models.Detailed, error) {

	res, err := repositories.ReadRecord[models.Detailed](db, d.Model.ID)
	if err != nil {
		return nil, fmt.Errorf("can not read : %v", err)
	} else {

		return res, nil
	}
}



