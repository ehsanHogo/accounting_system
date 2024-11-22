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

	sqlDB.SetMaxOpenConns(100)                
	sqlDB.SetMaxIdleConns(5)                 
	sqlDB.SetConnMaxLifetime(5 * time.Minute) 

	return NewConnection(db), nil
}

func CreateRecord[T any](db *gorm.DB, v *T) error {

	res := db.Create(v)
	if res.Error != nil {

		return fmt.Errorf("can not create record due to : %v", res.Error)

	} else {

		fmt.Println("Record created successfully")
		return nil
	}

}

func DeleteRecord[T any](db *gorm.DB, v *T) error {
	res := db.Unscoped().Delete(&v)

	if res.Error != nil {
		return fmt.Errorf("can not delete record due to: %v", res.Error)

	} else {

		fmt.Println("Record deleted successfully")
		return nil
	}

}

func ReadRecord[T any](db *gorm.DB, id uint) (*T, error) {
	var res T

	if err := db.First(&res, id).Error; err != nil {
		return nil, fmt.Errorf("record not found: %w", err)
	}
	return &res, nil
}

func UpdateRecord[T any](db *gorm.DB, v *T, id uint) error {

	if err := db.Model(v).Where("id = ?", id).Updates(v).Error; err != nil {
		return fmt.Errorf("can not  update record due to : %v", err)
	}

	return nil

}

func FindRecord[T any, U any](db *gorm.DB, val U, columnName string) bool {
	var res T
	if err := db.First(&res, fmt.Sprintf("%s = ?", columnName), val).Error; err != nil {
		return false
	}
	return true
}

func GenerateUniqeCode[T any](repo *gorm.DB, columnName string) string {
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

func GenerateUniqeTitle[T any](repo *gorm.DB) string {
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
