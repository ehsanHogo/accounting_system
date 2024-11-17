package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupConfig(t *testing.T) {

	//for restore .env file after test
	originalEnv := map[string]string{
		"POSTGRES_USER":     os.Getenv("POSTGRES_USER"),
		"POSTGRES_PASSWORD": os.Getenv("POSTGRES_PASSWORD"),
		"POSTGRES_DB":       os.Getenv("POSTGRES_DB"),
		"POSTGRES_HOST":     os.Getenv("POSTGRES_HOST"),
		"POSTGRES_PORT":     os.Getenv("POSTGRES_PORT"),
	}
	// Restore environment variables after the test
	defer func() {
		for k, v := range originalEnv {
			os.Setenv(k, v)
		}
	}()
	t.Run("operation is seccessful when all environment variables are set", func(t *testing.T) {
		os.Setenv("POSTGRES_USER", "ehsan")
		os.Setenv("POSTGRES_PASSWORD", "12345")
		os.Setenv("POSTGRES_DB", "testdb")
		os.Setenv("POSTGRES_HOST", "localhost")
		os.Setenv("POSTGRES_PORT", "5432")

		fmt.Printf("DB_USER: %s, DB_PASSWORD: %s, DB_NAME: %s, DB_HOST: %s, DB_PORT: %s\n",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
		)
		dbUrl, err := SetupConfig()
		assert.NoError(t, err, "expected no error")
		expectedUrl := "postgres://ehsan:12345@localhost:5432/testdb?sslmode=disable"
		assert.Equal(t, expectedUrl, dbUrl, "unexpected database URL")

	})



	t.Run("Error when environment variables are missing", func(t *testing.T) {
		os.Setenv("POSTGRES_USER", "")
		os.Setenv("POSTGRES_PASSWORD", "123")
		os.Setenv("POSTGRES_DB", "d")
		os.Setenv("POSTGRES_HOST", "me")
		os.Setenv("POSTGRES_PORT", "d")

		_, err := SetupConfig()
		assert.Error(t, err, "expected an error when environment variables are missing")
		assert.Equal(t, "missing required environment variables for database connection", err.Error(), "unexpected error message")
	})

}
