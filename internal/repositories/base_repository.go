package repositories

import (
	"accounting_system/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type Repositories struct {
	AccountingDB *gorm.DB
}

func NewConnection(db *gorm.DB) *Repositories {
	return &Repositories{
		AccountingDB: db,
	}
}

func CreateRecord[T any](db *Repositories, v *T) error {
	res := db.AccountingDB.Create(v)
	if res.Error != nil {
		return fmt.Errorf("error creating record: %w", res.Error)

	} else {

		fmt.Println("Record created successfully")
		return nil
	}

}

func UpdateDetailed(db *Repositories, v *models.Detailed, id uint) error {
	var newV models.Detailed
	if err := db.AccountingDB.First(&newV, id).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	newV.Code = v.Code
	newV.Title = v.Title

	fmt.Printf("newval %v", newV)

	if err := db.AccountingDB.Save(&newV).Error; err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}

func UpdateSubsidiary(db *Repositories, v *models.Subsidiary, id uint) error {
	var newV models.Subsidiary
	if err := db.AccountingDB.First(&newV, id).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	newV.Code = v.Code
	newV.Title = v.Title
	newV.HasDetailed = v.HasDetailed

	if err := db.AccountingDB.Save(&newV).Error; err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}

func UpdateVoucher(db *Repositories, v *models.Voucher, updatedItem []*models.VoucherItem, deletedItem []*models.VoucherItem, insertedItem []*models.VoucherItem, id uint) error {
	var newV models.Voucher
	if err := db.AccountingDB.First(&newV, id).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}
	newV.Number = v.Number
	newV.VoucherItems = insertedItem

	for _, vi := range deletedItem {
		db.AccountingDB.Delete(&vi)
	}

	for _, vi := range updatedItem {
		err := updateVoucherItem(db, vi, vi.Model.ID)
		if err != nil {
			return fmt.Errorf("can not update voucher item : %w", err)
		}
	}

	if err := db.AccountingDB.Save(&newV).Error; err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}
	return nil
}

func updateVoucherItem(db *Repositories, v *models.VoucherItem, id uint) error {
	var newV models.VoucherItem
	if err := db.AccountingDB.First(&newV, id).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	newV.Credit = v.Credit
	newV.Debit = v.Debit
	newV.DetailedId = v.DetailedId
	newV.SubsidiaryId = v.SubsidiaryId

	if err := db.AccountingDB.Save(&newV).Error; err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}

func DeleteRecord[T any](db *Repositories, v *T) error {
	res := db.AccountingDB.Delete(&v)

	if res.Error != nil {
		return fmt.Errorf("error in deleting record: %w", res.Error)

	} else {

		fmt.Println("Record deleted successfully")
		return nil
	}
}
