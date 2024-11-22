package temporary

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateVoucher(t *testing.T) {
	repo, err := repositories.CreateConnectionForTest()
	defer func() {
		sqlDB, _ := repo.AccountingDB.DB()
		sqlDB.Close()
	}()
	if err != nil {
		t.Fatalf("can not connect to database %v", err)
	}

	t.Run("create temp voucher record", func(t *testing.T) {

		voucher, err := CreateTempVoucher(repo)
		assert.NoError(t, err, "expected no error while creating temp voucher")

		_, err = repositories.ReadRecord[models.Voucher](repo, voucher.Model.ID)
		assert.NoError(t, err, "expected no error while reading temp voucher")

	})

	t.Run("create temp subsidairy record", func(t *testing.T) {

		subsidairy, err := CreateTempSubsidiary(repo)
		assert.NoError(t, err, "expected no error while creating temp subsidairy")

		_, err = repositories.ReadRecord[models.Subsidiary](repo, subsidairy.Model.ID)
		assert.NoError(t, err, "expected no error while reading temp subsidairy")

	})

	t.Run("create temp detailed record", func(t *testing.T) {

		detailed, err := CreateTempDetailed(repo)
		assert.NoError(t, err, "expected no error while creating temp detailed")

		_, err = repositories.ReadRecord[models.Detailed](repo, detailed.Model.ID)
		assert.NoError(t, err, "expected no error while reading temp detailed")

	})

}
