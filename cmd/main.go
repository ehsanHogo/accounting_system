package main

import (
	"accounting_system/config"
	"accounting_system/internal/repositories"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbUrl, err := config.SetupConfig()
	if err != nil {
		fmt.Println("Cant set database config")
		return
	}
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := repositories.NewConnection(db)
	fmt.Print(repo)

	println(dbUrl)
}
