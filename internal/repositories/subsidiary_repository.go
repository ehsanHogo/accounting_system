package repositories

import (
	"accounting_system/internal/models"
	"fmt"
)



func (db *Repositories) CreateSubsidiary(s *models.Subsidiary) {
	res := db.AccountingDB.Create(&models.Subsidiary{Code: s.Code, Title: s.Title, HasDetailed: s.HasDetailed})
	if res.Error != nil {
		fmt.Printf("Error creating record: %v\n", res.Error)
	} else {
		fmt.Println("Record created successfully")
	}

	println(res)
}