package repositories

import (
	"accounting_system/config"
	"accounting_system/internal/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateDetailed(t *testing.T) {

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

	repo := NewConnection(db)

	t.Run("the record successfully create", func(t *testing.T) {
		detailed := &models.Detailed{Code: "12", Title: "test"}
		CreateRecord(repo, detailed)
		var result models.Detailed
		err = repo.AccountingDB.First(&result, "code = ?", detailed.Code).Error //Code in uniqe
		assert.NoError(t, err, "expected detailed record to be created, but got error")

	})


}
