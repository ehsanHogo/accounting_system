package config

import (
	"errors"
	"fmt"

	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func SetupConfig() (string, error) {


	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	if dbHost == "" || dbPassword == "" || dbName == "" || dbUser == "" || dbPort == "" {
		// panic("Missing required environment variables for database connection")
		return "", errors.New("missing required environment variables for database connection")
	}
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	// m, err := migrate.New(
	// 	"file://internal/db/migrations",
	// 	dbUrl,
	// )
	// if err != nil {
	// 	fmt.Printf("failed to create migrate instance: %v\n", err)
	// 	return "", err
	// }

	// err = handleMigrateUp(m)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return "" , err
	// }
	// handleMigrateDown(m)
	return dbUrl, nil

}

// func handleMigrateUp(m *migrate.Migrate) error {
// 	if err := m.Up(); err != nil {
// 		if err.Error() == "no change" {
// 			fmt.Println("no change in database for migrate up")
// 			return nil
// 		}
// 		return fmt.Errorf("failed to perform migration up: %w", err)
// 	}
// 	return nil
// }

// func handleMigrateDown(m *migrate.Migrate) {
// 	if err := m.Down(); err != nil {
// 		if err.Error() == "no change" {
// 			log.Println("no change")
// 			return
// 		}
// 		log.Fatalf("failed to apply migration down: %v", err)
// 	}
// }
