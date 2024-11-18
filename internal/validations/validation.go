package validations

import (
	"accounting_system/internal/models"
	"errors"
	"fmt"
)

func CheckEmpty(s string) error {
	if len(s) == 0 {
		return errors.New("empty field not allowed")
	} else {
		return nil
	}
}

func CheckMaxLength(s string, maxLen int) error {
	if len(s) > maxLen {
		return fmt.Errorf("field length is greater than max length witch is %d", maxLen)
	} else {
		return nil
	}
}

func CheckDebitCredit(vi *models.VoucherItem) error {
	if vi.Debit < 0 || vi.Credit < 0 {
		return errors.New("debit or credit cant be negative")
	} else if vi.Debit == 0 && vi.Credit == 0 {
		return errors.New("both debit and credit cant be zero")
	} else if vi.Debit > 0 && vi.Credit > 0 {
		return errors.New("both debit and credit cant have positive value")
	} else {
		return nil
	}
}

func CheckBalance(v []*models.VoucherItem) error {
	var credits int64 = 0
	var debits int64 = 0

	for _, v := range v {
		if v.Credit != 0 {
			credits += v.Credit
		} else {
			debits += v.Debit
		}
	}
	if debits == credits {
		return nil
	} else {
		return errors.New("sum of credits and sum of debits cant be different")
	}
}

