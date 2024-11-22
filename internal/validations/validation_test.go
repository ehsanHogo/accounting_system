package validations

import (
	"accounting_system/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEmpty(t *testing.T) {
	t.Run("return error when find empty field", func(t *testing.T) {
		s := ""
		err := CheckEmpty(s)
		assert.Error(t, err, "Expected error indicate empty field ")
	})

	t.Run("when field is non empty , successfully return nil", func(t *testing.T) {
		s := "test"
		err := CheckEmpty(s)
		assert.NoError(t, err, "Expected no error for non empty field")
	})
}

func TestCheckMaxLength(t *testing.T) {
	t.Run("return error when length is greater than 64", func(t *testing.T) {
		s := "a9b8c7d6e5f4g3h2i1j0k9l8m7n6o5p4q3r2s1t0u9v8w7x6y5z4a3b2c1d012345"
		err := CheckMaxLength(s, 64)
		assert.Error(t, err, "Expected error indicate length greater than max length")
	})

	t.Run("when field length is less than 64 , successfully return nil", func(t *testing.T) {
		s := "test"
		maxl := 64
		err := CheckMaxLength(s, maxl)
		assert.NoError(t, err, "Expected no error for length less than max length %d", maxl)
	})

	t.Run("when field length is equal 64 , successfully return nil", func(t *testing.T) {
		s := "a9b8c7d6e5f4g3h2i1j0k9l8m7n6o5p4q3r2s1t0u9v8w7x6y5z4a3b2c1d01234"
		maxl := 64
		err := CheckMaxLength(s, 64)
		assert.NoError(t, err, "Expected no error for length less than max length %d", maxl)
	})
}

func TestCheckDebitCredit(t *testing.T) {
	t.Run("return error if one of the debit or credit was negative", func(t *testing.T) {
		vi := &models.VoucherItem{Debit: 12, Credit: -1}
		err := CheckDebitCredit(vi.Credit, vi.Debit)
		assert.Error(t, err, "expected error indicate negative value is invalied")
		vi = &models.VoucherItem{Debit: -2, Credit: 11}
		err = CheckDebitCredit(vi.Credit, vi.Debit)
		assert.Error(t, err, "expected error indicate negative value is invalied")
	})

	t.Run("return error if both debit and credit were zero", func(t *testing.T) {
		vi := &models.VoucherItem{Debit: 0, Credit: 0}
		err := CheckDebitCredit(vi.Credit, vi.Debit)
		assert.Error(t, err, "expected error indicate both debit and credit cant be zero")
	})

	t.Run("return error if both debit and credit had positive value", func(t *testing.T) {
		vi := &models.VoucherItem{Debit: 3, Credit: 5}
		err := CheckDebitCredit(vi.Credit, vi.Debit)
		assert.Error(t, err, "expected error indicate both debit and credit cant have positive value")
	})

	t.Run("when one of them is zero and the other is positive value , successfully return nil ", func(t *testing.T) {
		vi := &models.VoucherItem{Debit: 2, Credit: 0}
		err := CheckDebitCredit(vi.Credit, vi.Debit)
		assert.NoError(t, err, "expected no error")

		vi = &models.VoucherItem{Debit: 0, Credit: 3}
		err = CheckDebitCredit(vi.Credit, vi.Debit)
		assert.NoError(t, err, "expected no error")
	})

}

func TestCheckBalance(t *testing.T) {
	t.Run("return error when sum of credits and dibits is different", func(t *testing.T) {
		temp := make([]*models.VoucherItem, 3)
		temp[0] = &models.VoucherItem{Credit: 0, Debit: 4}
		temp[1] = &models.VoucherItem{Credit: 2, Debit: 0}
		temp[2] = &models.VoucherItem{Credit: 1, Debit: 0}

		err := CheckBalance(temp)
		assert.Error(t, err, "expected error indicate sum of debits and credits is different")
	})

	t.Run("when sum of credits and dibits is same , successfully return nil", func(t *testing.T) {
		temp := make([]*models.VoucherItem, 3)
		temp[0] = &models.VoucherItem{Credit: 5, Debit: 0}
		temp[1] = &models.VoucherItem{Credit: 0, Debit: 2}
		temp[2] = &models.VoucherItem{Credit: 0, Debit: 3}

		err := CheckBalance(temp)
		assert.NoError(t, err, "expected no error")
	})

}

func TestCheckVoucherItemsLength(t *testing.T) {
	t.Run("the length of voucher item is valied", func(t *testing.T) {
		temp := make([]*models.VoucherItem, 3)
		temp[0] = &models.VoucherItem{Credit: 5, Debit: 0}
		temp[1] = &models.VoucherItem{Credit: 0, Debit: 2}
		temp[2] = &models.VoucherItem{Credit: 0, Debit: 3}

		err := CheckVoucherItemsLength(len(temp))
		assert.NoError(t, err, "ecpected no error in length of voucher item list")
	})

	t.Run("the length of voucher item is invalied due to length of less than 2", func(t *testing.T) {
		temp := make([]*models.VoucherItem, 1)
		temp[0] = &models.VoucherItem{Credit: 5, Debit: 0}

		err := CheckVoucherItemsLength(len(temp))
		assert.Error(t, err, "ecpected error indicate length of voucher items cant be less than 2")
	})

	t.Run("the length of voucher item is invalied due to length of greater than 500", func(t *testing.T) {
		temp := make([]*models.VoucherItem, 501)

		err := CheckVoucherItemsLength(len(temp))
		assert.Error(t, err, "ecpected error indicate length of voucher items cant be greater than 500")
	})
}
