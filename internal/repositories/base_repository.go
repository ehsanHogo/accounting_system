package repositories

import (
	"accounting_system/config"
	"accounting_system/internal/utils/randgenerator"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
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

func CreateConnectionForTest() (*Repositories, error) {
	dbUrl, err := config.SetupConfig()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)                // Limit the maximum number of open connections
	sqlDB.SetMaxIdleConns(5)                  // Set idle connection limit
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Limit connection lifetime

	return NewConnection(db), nil
}

func CreateRecord[T any](db *Repositories, v *T) error {

	res := db.AccountingDB.Create(v)
	if res.Error != nil {

		return fmt.Errorf("can not create record due to : %v", res.Error)

	} else {

		fmt.Println("Record created successfully")
		return nil
	}

}

func DeleteRecord[T any](db *Repositories, v *T) error {
	res := db.AccountingDB.Unscoped().Delete(&v)

	if res.Error != nil {
		return fmt.Errorf("can not delete record due to: %v", res.Error)

	} else {

		fmt.Println("Record deleted successfully")
		return nil
	}

}

func ReadRecord[T any](db *Repositories, id uint) (*T, error) {
	var res T

	if err := db.AccountingDB.First(&res, id).Error; err != nil {
		return nil, fmt.Errorf("record not found: %w", err)
	}
	return &res, nil
}

func UpdateRecord[T any](db *Repositories, v *T, id uint) error {

	if err := db.AccountingDB.Model(v).Where("id = ?", id).Updates(v).Error; err != nil {
		return fmt.Errorf("can not  update record due to : %v", err)
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

func GenerateUniqeCode[T any](repo *Repositories, columnName string) string {
	code := randgenerator.GenerateRandomCode()
	for {
		exist := FindRecord[T](repo, code, columnName)

		if exist {
			code = randgenerator.GenerateRandomCode()
		} else {
			break
		}
	}

	return code
}

func GenerateUniqeTitle[T any](repo *Repositories) string {
	title := randgenerator.GenerateRandomTitle()
	for {
		exist := FindRecord[T](repo, title, "title")

		if exist {
			title = randgenerator.GenerateRandomTitle()
		} else {
			break
		}
	}

	return title
}
