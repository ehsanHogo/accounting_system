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

func (db *Repositories) CreateDetailed(d *models.Detailed) {
	res := db.AccountingDB.Create(&models.Detailed{Code: d.Code, Title: d.Title})
	if res.Error != nil {
		fmt.Printf("Error creating record: %v\n", res.Error)
	} else {
		fmt.Println("Record created successfully")
	}

	println(res)
}

func (db *Repositories) CreateSubsidiary(s *models.Subsidiary) {
	res := db.AccountingDB.Create(&models.Subsidiary{Code: s.Code, Title: s.Title, HasDetailed: s.HasDetailed})
	if res.Error != nil {
		fmt.Printf("Error creating record: %v\n", res.Error)
	} else {
		fmt.Println("Record created successfully")
	}

	println(res)
}

func (db *Repositories) CreateVoucher(v *models.Voucher, vi []models.VoucherItem) {
	res := db.AccountingDB.Create(&models.Voucher{Number: v.Number, VoucherItems: vi})
	if res.Error != nil {
		fmt.Printf("Error creating voucher record: %v\n", res.Error)
	} else {
		fmt.Println("Record created successfully")
	}

	println(res)
}
