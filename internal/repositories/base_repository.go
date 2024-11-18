package repositories

import (
	"accounting_system/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type Repositories struct {
	AccountingDB *gorm.DB
}

func NewConnection(db *gorm.DB) *Repositories {
	return &Repositories{
		AccountingDB: db,
	}
}

func CreateRecord[T any](db *Repositories, v *T) error {
	res := db.AccountingDB.Create(v)
	if res.Error != nil {
		return fmt.Errorf("error creating record: %w", res.Error)

	} else {

		fmt.Println("Record created successfully")
		return nil
	}

}

func UpdateDetailed(db *Repositories, v *models.Detailed, id uint) error {
	var newV models.Detailed
	if err := db.AccountingDB.First(&newV, id).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	newV.Code = v.Code
	newV.Title = v.Title

	fmt.Printf("newval %v", newV)

	if err := db.AccountingDB.Save(&newV).Error; err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}
