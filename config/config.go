package config

import (
	"errors"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	DB       = "accounting"
	USER     = "postgres"
	PASSWORD = "12551255"
	HOST     = "localhost"
	PORT     = "5432"
)

func SetupConfig() (string, error) {

	if HOST == "" || PASSWORD == "" || DB == "" || USER == "" || PORT == "" {
		return "", errors.New("missing required environment variables for database connection")
	}
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", USER, PASSWORD, HOST, PORT, DB)

	return dbUrl, nil

}
