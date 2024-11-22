package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupConfig(t *testing.T) {

	t.Run("operation is seccessful when all environment variables are set", func(t *testing.T) {

		dbUrl, err := SetupConfig()
		assert.NoError(t, err, "expected no error")
		expectedUrl := "postgres://postgres:12551255@localhost:5432/accounting?sslmode=disable"
		assert.Equal(t, expectedUrl, dbUrl, "unexpected database URL")

	})

}
