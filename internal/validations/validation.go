package validations

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const (
	MaxCodeLength  = 64
	MaxTitleLength = 64
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

func CheckDebitCredit(credit, debit int32) error {
	if debit < 0 || credit < 0 {
		return errors.New("debit or credit cant be negative")
	} else if debit == 0 && credit == 0 {
		return errors.New("both debit and credit cant be zero")
	} else if debit > 0 && credit > 0 {
		return errors.New("both debit and credit cant have positive value")
	} else {
		return nil
	}
}

func CheckBalance(v []*models.VoucherItem) error {
	var credits int32   = 0
	var debits int32 = 0

	for _, v := range v {
		credits += v.Credit
		debits += v.Debit
	}
	if debits == credits {
		return nil
	} else {
		return errors.New("sum of credits and sum of debits cant be different")
	}
}

func CheckVoucherItemsLength(length int) error {
	if length < 2 {
		return fmt.Errorf("number of voucherItems can be less than 2 ")
	} else if length > 500 {
		return fmt.Errorf("number of voucherItems can be greater than 500 ")

	} else {
		return nil
	}
}

func ChackCodeValidation(c string) error {
	err := CheckEmpty(c)
	if err != nil {
		return err
	}

	err = CheckMaxLength(c, MaxCodeLength)
	if err != nil {
		return err
	}

	return nil
}

func CheckTitleValidaion(t string) error {
	err := CheckEmpty(t)
	if err != nil {
		return err
	}

	err = CheckMaxLength(t, MaxTitleLength)
	if err != nil {
		return err
	}

	return nil
}

func InsertDetailedValidation(d *models.Detailed) error {
	err := ChackCodeValidation(d.Code)

	if err != nil {
		return fmt.Errorf("code validation fail due to  : %v", err)
	}

	err = CheckTitleValidaion(d.Title)

	if err != nil {
		return fmt.Errorf("title validation fail due to : %v", err)
	}

	return nil
}

func UpdateDetailedValidation(repo *gorm.DB, d *models.Detailed) error {

	prevDetailed, err := repositories.ReadRecord[models.Detailed](repo, d.ID)
	if err != nil {
		return fmt.Errorf("update validation fail due to absence of detailed id in database  : %v", err)
	}

	if d.Version != prevDetailed.Version {
		return fmt.Errorf("delete validation fail due to different versions , expected version = %d , got : %d", prevDetailed.Version, d.Version)

	}

	err = ChackCodeValidation(d.Code)

	if err != nil {
		return fmt.Errorf("code validation fail due to  : %v", err)
	}

	err = CheckTitleValidaion(d.Title)

	if err != nil {
		return fmt.Errorf("title validation fail due to : %v", err)
	}

	return nil
}

func InsertSubsidiaryValidation(d *models.Subsidiary) error {
	err := ChackCodeValidation(d.Code)

	if err != nil {
		return fmt.Errorf("code validation fail due to  : %v", err)
	}

	err = CheckTitleValidaion(d.Title)

	if err != nil {
		return fmt.Errorf("title validation fail due to : %v", err)
	}

	return nil
}

func UpdateSubsidiaryValidation(repo *gorm.DB, d *models.Subsidiary) error {
	prevSubsidiary, err := repositories.ReadRecord[models.Subsidiary](repo, d.ID)
	if err != nil {
		return fmt.Errorf("update validation fail due to absence of subsidiary id in database  : %v", err)
	}

	if d.Version != prevSubsidiary.Version {
		return fmt.Errorf("delete validation fail due to different versions , expected version = %d , got : %d", prevSubsidiary.Version, d.Version)

	}

	var voucherHasThisSubsidiary models.VoucherItem
	if err := repo.First(&voucherHasThisSubsidiary, "subsidiary_id = ?", d.ID).Error; err == nil {

		return fmt.Errorf("can not update subsidiary record because it is reffrenced by some voucherItems")

	}

	err = ChackCodeValidation(d.Code)

	if err != nil {
		return fmt.Errorf("code validation fail due to  : %v", err)
	}

	err = CheckTitleValidaion(d.Title)

	if err != nil {
		return fmt.Errorf("title validation fail due to : %v", err)
	}

	return nil
}

func DeleteDetailedValidation(db *gorm.DB, d *models.Detailed) error {

	prevDetailed, err := repositories.ReadRecord[models.Detailed](db, d.ID)
	if err != nil {
		return fmt.Errorf("delete validation fail due to absence of detailed id in database  : %v", err)
	}

	if d.Version == prevDetailed.Version {
		return nil
	} else {
		return fmt.Errorf("delete validation fail due to different versions , expected version = %d , got : %d", prevDetailed.Version, d.Version)
	}
}

func DeleteSubsidiaryValidation(db *gorm.DB, d *models.Subsidiary) error {

	prevDetailed, err := repositories.ReadRecord[models.Subsidiary](db, d.ID)
	if err != nil {
		return fmt.Errorf("delete validation fail due to absence of subsidiary id in database  : %v", err)
	}

	if d.Version == prevDetailed.Version {
		return nil
	} else {
		return fmt.Errorf("delete validation fail due to different versions , expected version = %d , got : %d", prevDetailed.Version, d.Version)
	}
}

func InsertVoucherValidation(db *gorm.DB, d *models.Voucher) error {

	err := ChackCodeValidation(d.Number)

	if err != nil {
		return fmt.Errorf("number validation fail due to : %v", err)
	}

	err = CheckBalance(d.VoucherItems)

	if err != nil {
		return fmt.Errorf("balance voucherItem validation fail due to : %v", err)
	}

	for _, v := range d.VoucherItems {

		err = CheckDebitCredit(v.Credit, v.Debit)
		if err != nil {
			return fmt.Errorf("debit and credit validation fail due to : %v", err)
		}
	}

	err = CheckVoucherItemsLength(len(d.VoucherItems))
	if err != nil {
		return fmt.Errorf("length of voucher items is invalied due to : %v", err)
	}

	err = checkHasDetailed(db, d.VoucherItems)

	if err != nil {
		return fmt.Errorf("there are invalied voucher items due to : %v", err)
	}

	return nil
}

func UpdateVoucherValidation(db *gorm.DB, d *models.Voucher, updatedItem []*models.VoucherItem, deletedItem []*models.VoucherItem, insertedItem []*models.VoucherItem) error {

	prevVoucher, err := repositories.ReadRecord[models.Voucher](db, d.ID)
	if err != nil {
		return fmt.Errorf("delete validation fail due to absence of voucher id in database  : %v", err)
	}

	if prevVoucher.Version != d.Version {
		return fmt.Errorf("update validation fail due to different versions , expected version = %d , got : %d", prevVoucher.Version, d.Version)
	}
	err = ChackCodeValidation(d.Number)

	if err != nil {
		return fmt.Errorf("number validation fail due to : %v", err)
	}

	for _, v := range insertedItem {

		err = CheckDebitCredit(v.Credit, v.Debit)
		if err != nil {
			return fmt.Errorf("debit and credit validation in inserted voucher items fail due to : %v", err)
		}
	}

	err = CheckBalance(insertedItem)

	if err != nil {
		return fmt.Errorf("balance credit and debit validation in inserted voucher items fail  due to : %v", err)
	}

	for _, v := range updatedItem {

		err = CheckDebitCredit(v.Credit, v.Debit)
		if err != nil {
			return fmt.Errorf("debit and credit validation in updated voucher items fail due to : %v", err)
		}
	}

	err = CheckBalance(updatedItem)

	if err != nil {
		return fmt.Errorf("balance credit and debit validation in updated voucher items  fail due to : %v", err)
	}

	err = CheckBalance(deletedItem)

	if err != nil {
		return fmt.Errorf("balance redit and debit  validation in deleted voucher items fail due to : %v", err)
	}

	var prevVoucherItems []*models.VoucherItem

	result := db.Where("voucher_id = ?", fmt.Sprintf("%d", d.ID)).Find(&prevVoucherItems)

	if result.Error != nil {
		return fmt.Errorf("can not fetch prev voucherItems due to : %v ", result.Error)
	}

	exists := make(map[int64]bool)
	for _, val := range deletedItem {
		exists[val.ID] = true
	}

	newVoucherItems := []*models.VoucherItem{}
	for _, val := range prevVoucherItems {
		if !exists[val.ID] {
			newVoucherItems = append(newVoucherItems, val)
		}
	}

	newVoucherItems = append(newVoucherItems, insertedItem...)

	err = CheckVoucherItemsLength(len(newVoucherItems))

	if err != nil {
		return fmt.Errorf("length of voucher items is invalied due to : %v", err)
	}

	err = checkHasDetailed(db, insertedItem)

	if err != nil {
		return fmt.Errorf("there are invalied voucher items due to : %v", err)
	}

	err = checkHasDetailed(db, updatedItem)

	if err != nil {
		return fmt.Errorf("there are invalied voucher items due to : %v", err)
	}

	return nil
}

func checkHasDetailed(repo *gorm.DB, vi []*models.VoucherItem) error {
	var subsidiary *models.Subsidiary
	for _, v := range vi {
		err := repo.First(&subsidiary, v.SubsidiaryId).Error

		if err != nil {
			return fmt.Errorf("can not read subsidiary record %v : %v", v.SubsidiaryId, err)
		}

		if subsidiary.HasDetailed {
			if v.DetailedId == 0 {
				return fmt.Errorf("can not have empty detailed when subsidiary has detailed")
			}
		} else {
			if v.DetailedId != 0 {
				return fmt.Errorf("can not have detailed when subsidiary does not have detailed")

			}
		}

		subsidiary = nil
	}

	return nil
}

func DeleteVoucherValidation(db *gorm.DB, d *models.Voucher) error {
	prevVoucher, err := repositories.ReadRecord[models.Voucher](db, d.ID)
	if err != nil {
		return fmt.Errorf("can not find voucher record : %v", err)
	}

	if prevVoucher.Version != d.Version {
		return fmt.Errorf("update validation fail due to different versions , expected version = %d , got : %d", prevVoucher.Version, d.Version)
	}

	return nil

}
