package repositories

import (
	"accounting_system/internal/models"
	"fmt"
)

func (db *Repositories) CreateVoucher(v *models.Voucher) {
	res := db.AccountingDB.Create(&models.Voucher{Number: v.Number, VoucherItems: v.VoucherItems})
	if res.Error != nil {
		fmt.Printf("Error creating voucher record: %v\n", res.Error)
	} else {
		fmt.Println("Record created successfully")
	}

	println(res)
}
