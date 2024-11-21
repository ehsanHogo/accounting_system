package repositories

import (
	"accounting_system/internal/models"
	"errors"
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

// func CreateConnectionForTest() (*Repositories, error) {
// 	dbUrl, err := config.SetupConfig()
// 	if err != nil {
// 		return nil, err
// 	}
// 	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	sqlDB, _ := db.DB()

// 	sqlDB.SetMaxOpenConns(100)                // Limit the maximum number of open connections
// 	sqlDB.SetMaxIdleConns(5)                  // Set idle connection limit
// 	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Limit connection lifetime

// 	return NewConnection(db), nil
// }

func CreateRecord[T any](db *Repositories, v *T) error {

	// Create record

	res := db.AccountingDB.Create(v)
	if res.Error != nil {

		return fmt.Errorf("can not create record due to : %v", res.Error)

	} else {

		fmt.Println("Record created successfully")
		return nil
	}

}

func DeleteVoucherItemRecord[T any](db *Repositories, v *T) error {
	res := db.AccountingDB.Unscoped().Delete(&v)

	if res.Error != nil {
		return fmt.Errorf("can not delete record due to: %v", res.Error)

	} else {

		fmt.Println("Record deleted successfully")
		return nil
	}

}

func DeleteDetailedRecord(db *Repositories, v *models.Detailed) error {

	// var prev *models.Detailed
	// var err error
	// prev, err := ReadRecord[models.Detailed](db, v.Model.ID)
	// if err != nil {
	// 	return fmt.Errorf("can not delete detailed record : %v", err)
	// } else {

	// if v.Version != prev.Version {
	// 	return errors.New("can not delete because of different version")
	// } else {

	res := db.AccountingDB.Unscoped().Delete(&v)

	if res.Error != nil {
		return fmt.Errorf("can not  delete detailed record due to : %v", res.Error)

	} else {

		fmt.Println("Record deleted successfully")
		return nil
	}
	// }
	// }

}

func DeleteSubsidiaryRecord(db *Repositories, v *models.Subsidiary) error {

	var prev *models.Subsidiary
	var err error
	prev, err = ReadRecord[models.Subsidiary](db, v.Model.ID)
	if err != nil {
		return fmt.Errorf("can not delete subsidiary record : %v", err)
	} else {

		if v.Version != prev.Version {
			return errors.New("can not delete because of different version")
		} else {

			res := db.AccountingDB.Unscoped().Delete(&v)

			if res.Error != nil {
				return fmt.Errorf("error in deleting record: %w", res.Error)

			} else {

				fmt.Println("Record deleted successfully")
				return nil
			}
		}
	}

}

func DeleteVoucherRecord(db *Repositories, v *models.Voucher) error {

	var prev *models.Voucher
	var err error
	prev, err = ReadRecord[models.Voucher](db, v.Model.ID)
	if err != nil {
		return fmt.Errorf("can not delete voucher record : %v", err)
	} else {

		if v.Version != prev.Version {
			return errors.New("can not delete because of different version")
		}

		res := db.AccountingDB.Unscoped().Delete(&v)

		if res.Error != nil {
			return fmt.Errorf("error in deleting record: %w", res.Error)

		} else {

			fmt.Println("Record deleted successfully")
			return nil
		}
		// }
	}

}

func ReadRecord[T any](db *Repositories, id uint) (*T, error) {
	var res T
	var genericType string
	switch any(res).(type) {
	case models.Detailed:
		genericType = "detailed"
	case models.Subsidiary:
		genericType = "subsidiary"

	case models.Voucher:
		genericType = "voucher"

	case models.VoucherItem:
		genericType = "voucherItem"
	default:
		genericType = ""
	}

	// fmt.Println("generic")
	// fmt.Println(genericType)
	// typeMap := map[any]string{
	// 	models.Detailed{}:    "detailed",
	// 	models.Subsidiary{}:  "subsidiary",
	// 	models.Voucher{}:     "voucher",
	// 	models.VoucherItem{}: "voucherItem",
	// }

	// genericType, ok := typeMap[any(res)]
	// if !ok {
	// 	return nil, fmt.Errorf("unsupported type for ReadRecord")
	// }

	if err := db.AccountingDB.First(&res, id).Error; err != nil {
		return nil, fmt.Errorf("%s record not found: %w", genericType, err)
	}
	return &res, nil
}

func UpdateDetailed(db *Repositories, v *models.Detailed, id uint) error {
	var newV models.Detailed
	if err := db.AccountingDB.First(&newV, id).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	if v.Version != newV.Version {
		return fmt.Errorf("can not update , the version of detailed record is different. expected version : %v", newV.Version)
	} else {

		var voucherHasThisDetailed models.VoucherItem
		if err := db.AccountingDB.First(&voucherHasThisDetailed, "detailed_id = ?", id).Error; err != nil {

			fmt.Printf("errrrrrr : %v\n", err)
			newV.Code = v.Code
			newV.Title = v.Title
			newV.Version += 1
			fmt.Printf("newval %v", newV)

			if err := db.AccountingDB.Save(&newV).Error; err != nil {
				return fmt.Errorf("failed to update record: %w", err)
			}

			return nil
		} else { //ther is a voucher has this detailed
			return fmt.Errorf("ccan not update detailed record because it is reffrenced by some voucherItems")
		}

	}
}

func UpdateSubsidiary(db *Repositories, v *models.Subsidiary, id uint) error {
	var newV models.Subsidiary
	if err := db.AccountingDB.First(&newV, id).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	if v.Version != newV.Version {
		return fmt.Errorf("can not update , the version of subsidiary record is different. expected version : %v", newV.Version)
	} else {

		newV.Code = v.Code
		newV.Title = v.Title
		newV.HasDetailed = v.HasDetailed
		newV.Version += 1
		if err := db.AccountingDB.Save(&newV).Error; err != nil {
			return fmt.Errorf("failed to update record: %w", err)
		}

		return nil
	}
}

func UpdateVoucher(db *Repositories, v *models.Voucher, updatedItem []*models.VoucherItem, deletedItem []*models.VoucherItem, insertedItem []*models.VoucherItem, id uint) error {

	var newV models.Voucher
	if err := db.AccountingDB.First(&newV, id).Error; err != nil {
		return fmt.Errorf("record not found: %w", err)
	}

	if v.Version != newV.Version {
		return fmt.Errorf("can not update , the version of detailed record is different. expected version : %v", newV.Version)
	} else {

		newV.Number = v.Number
		newV.VoucherItems = insertedItem
		newV.Version += 1

		for _, vi := range deletedItem {

			err := DeleteVoucherItemRecord(db, vi)
			if err != nil {
				return fmt.Errorf("can not update voucher item : %w", err)
			}
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

func FindRecord[T any, U any](db *Repositories, val U, columnName string) bool {
	var res T
	if err := db.AccountingDB.First(&res, fmt.Sprintf("%s = ?", columnName), val).Error; err != nil {
		return false
	}
	return true
}
