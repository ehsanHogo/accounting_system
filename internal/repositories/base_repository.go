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

func CreateRecord[T any](db *Repositories, v *T) {
	res := db.AccountingDB.Create(v)
	if res.Error != nil {
		fmt.Printf("Error creating record: %v\n", res.Error)
	} else {
		fmt.Println("Record created successfully")
	}

	

	println(res)
}
