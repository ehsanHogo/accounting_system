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
	var prevDetailed *models.Detailed
	prevDetailed, err = repositories.ReadRecord[models.Detailed](db, d.Model.ID)

	if err != nil {
		return err
	}

	if prevDetailed.Version == d.Version {

		repositories.UpdateDetailed(db, d, d.Model.ID)
		return nil
	} else {
		return fmt.Errorf("the version is different : %v", err)
	}

}

func DeleteDetailed(db *repositories.Repositories, d *models.Detailed) error {
	var prevDetailed *models.Detailed
	prevDetailed, err := repositories.ReadRecord[models.Detailed](db, d.Model.ID)

	if err != nil {
		return err
	}

	if prevDetailed.Version == d.Version {

		repositories.DeleteDetailedRecord(db, d)
		return nil
	} else {
		return fmt.Errorf("the version is different : %v", err)
	}

}
