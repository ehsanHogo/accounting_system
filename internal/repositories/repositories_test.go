package repositories

import (
	"accounting_system/config"
	"accounting_system/internal/models"
	"errors"
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
		err := CreateRecord(repo, detailed)

		assert.NoError(t, err, "expected detailed record to be created, but got error")
		var result models.Detailed
		err = repo.AccountingDB.First(&result, "code = ?", detailed.Code).Error //Code in uniqe
		assert.NoError(t, err, " can not find the inserted record :")

	})

	t.Run("the record fail because duplication", func(t *testing.T) {
		// fmt.Println("jdgfsdgh")
		detailed := &models.Detailed{Code: "12", Title: "test"}
		err := CreateRecord(repo, detailed)
		assert.NoError(t, err, "expected detailed record to be created, but got error")

		detailed = &models.Detailed{Code: "12", Title: "dup"}
		// err = CreateRecord(repo, detailed)
		err = errors.New("sdfjksd")
		if err != nil {
			t.Fatalf("duplicated record : %v", err.Error())
		}

	})

}
