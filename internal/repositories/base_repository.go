package repositories

import (
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
