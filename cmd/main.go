package main

import (
	"accounting_system/config"
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbUrl, err := config.SetupConfig()
	if err != nil {
		fmt.Printf("Cant set database config : %v", err)
		return
	}
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect to database: %v", err)
		return
	}

	repo := repositories.NewConnection(db)

	temp := make([]models.VoucherItem, 2)
	fmt.Printf("eeeee : %v\n", temp)
	repo.CreateVoucher(&models.Voucher{Number: "13", VoucherItems: temp})

}
