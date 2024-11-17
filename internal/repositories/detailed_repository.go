package repositories

import (
	"accounting_system/internal/models"
	"fmt"
)



func (db *Repositories) CreateDetailed(d *models.Detailed) {
	res := db.AccountingDB.Create(&models.Detailed{Code: d.Code, Title: d.Title})
	if res.Error != nil {
		fmt.Printf("Error creating record: %v\n", res.Error)
	} else {
		fmt.Println("Record created successfully")
	}

	println(res)
}