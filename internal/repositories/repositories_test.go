package repositories

// import (
// 	"accounting_system/config"
// 	"accounting_system/internal/models"

// 	randgenerator "accounting_system/internal/utils"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func createConnectionForTest() (*Repositories, error) {
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
// func TestCreateDetailed(t *testing.T) {

// 	repo, err := createConnectionForTest()
// 	defer func() {
// 		sqlDB, _ := repo.AccountingDB.DB()
// 		sqlDB.Close()
// 	}()

// 	if err != nil {
// 		t.Fatalf("can not connect to database %v", err)
// 	}

// 	t.Run("the detailed record successfully create", func(t *testing.T) {

// 		detailed, err := createTempDetailed(repo)

// 		assert.NoError(t, err, "can not create detailed record due to")

// 		// fmt.Printf("Detailed ID: %v\n", detailed.Model.ID)

// 		var result models.Detailed
// 		err = repo.AccountingDB.First(&result, detailed.Model.ID).Error

// 		assert.NoError(t, err, "can not find the inserted detailed record")
// 	})

// 	// t.Run("the detailed record creation fail because duplication code", func(t *testing.T) {

// 	// 	detailed, err := createTempDetailed(repo)
// 	// 	assert.NoError(t, err, "can not create detailed record due to")
// 	// 	detailed.Title = randgenerator.GenerateRandomTitle()

// 	// 	err = CreateRecord(repo, detailed)

// 	// 	assert.Error(t, err, "expected getting duplicate detailed code error")

// 	// })

// 	// t.Run("the detailed record creation fail because duplication title", func(t *testing.T) {

// 	// 	detailed, err := createTempDetailed(repo)
// 	// 	assert.NoError(t, err, "can not create detailed record due to")

// 	// 	detailed.Code = randgenerator.GenerateRandomCode()

// 	// 	err = CreateRecord(repo, detailed)

// 	// 	assert.Error(t, err, "expected getting duplicate detailed title error")

// 	// })

// }

// func TestCreateSubsidiary(t *testing.T) {

// 	repo, err := createConnectionForTest()
// 	defer func() {
// 		sqlDB, _ := repo.AccountingDB.DB()
// 		sqlDB.Close()
// 	}()

// 	if err != nil {
// 		t.Fatalf("can not connect to database %v", err)
// 	}

// 	t.Run("the subsidiary record successfully create", func(t *testing.T) {

// 		subsidiary, err := createTempSubsidiary(repo)
// 		assert.NoError(t, err, "can not create subsidiary record due to")

// 		var result models.Subsidiary
// 		err = repo.AccountingDB.First(&result, subsidiary.Model.ID).Error //Code is uniqe
// 		assert.NoError(t, err, " can not find the inserted subsidiary record :")

// 	})

// 	// t.Run("the subsidiary record creation fail because duplication code", func(t *testing.T) {

// 	// 	subsidiary, err := createTempSubsidiary(repo)
// 	// 	assert.NoError(t, err, "can not create subsidiary record due to")

// 	// 	subsidiary.Title = randgenerator.GenerateRandomTitle()

// 	// 	err = CreateRecord(repo, subsidiary)

// 	// 	assert.Error(t, err, "expected getting duplicate subsidiary code error")

// 	// })

// 	// t.Run("the subsidiary record creation fail because duplication title", func(t *testing.T) {

// 	// 	subsidiary, err := createTempSubsidiary(repo)
// 	// 	assert.NoError(t, err, "can not create subsidiary record due to")

// 	// 	subsidiary.Code = randgenerator.GenerateRandomCode()

// 	// 	err = CreateRecord(repo, subsidiary)

// 	// 	assert.Error(t, err, "expected getting duplicate subsidiary title error")

// 	// })

// }

// func TestCreateVoucher(t *testing.T) {

// 	repo, err := createConnectionForTest()
// 	defer func() {
// 		sqlDB, _ := repo.AccountingDB.DB()
// 		sqlDB.Close()
// 	}()
// 	if err != nil {
// 		t.Fatalf("can not connect to database %v", err)
// 	}

// 	t.Run("the voucher record successfully create", func(t *testing.T) {

// 		voucher, err := createTempVoucher(repo)
// 		assert.NoError(t, err, "can not create voucher record")

// 		// assert.NoError(t, err, "expected voucher record to be created, but got error")

// 		var result models.Voucher
// 		err = repo.AccountingDB.First(&result, voucher.Model.ID).Error //Number is uniqe
// 		assert.NoError(t, err, " can not find the inserted voucher record :")

// 	})

// 	// t.Run("the voucher record creation fail because duplication number", func(t *testing.T) {
// 	// 	voucher, err := createTempVoucher(repo)

// 	// 	assert.NoError(t, err, "can not create voucehr record")

// 	// 	err = CreateRecord(repo, voucher)

// 	// 	assert.Error(t, err, "expected getting duplicate voucher number error")

// 	// })

// }

// func TestUpdateDetailed(t *testing.T) {

// 	repo, err := createConnectionForTest()
// 	defer func() {
// 		sqlDB, _ := repo.AccountingDB.DB()
// 		sqlDB.Close()
// 	}()
// 	if err != nil {
// 		t.Fatalf("can not connect to database %v", err)
// 	}
// 	t.Run("can update detailed record successfully", func(t *testing.T) {

// 		detailed, err := createTempDetailed(repo)
// 		assert.NoError(t, err, "can not create detailed record")

// 		prevDetailedId := detailed.Model.ID
// 		detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
// 		detailed.Title = generateUniqeTitle[models.Detailed](repo)
// 		// detailed = &models.Detailed{Code: code, Title: title}
// 		err = UpdateDetailed(repo, detailed, prevDetailedId)
// 		assert.NoError(t, err, "can not update detailed record")
// 	})

// 	// t.Run("return error when update detailed record that is not in databse", func(t *testing.T) {
// 	// 	code := randgenerator.GenerateRandomCode()
// 	// 	title := randgenerator.GenerateRandomTitle()
// 	// 	detailed := &models.Detailed{Code: code, Title: title}

// 	// 	err := UpdateDetailed(repo, detailed, 1_000_000)
// 	// 	assert.Error(t, err, "expected error indicate there is such id in database")

// 	// })

// 	// t.Run("can not update detailed record if versions were  different", func(t *testing.T) {
// 	// 	detailed, err := createTempDetailed(repo)
// 	// 	assert.NoError(t, err, "can not create detailed record due to")

// 	// 	detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
// 	// 	// fmt.Printf("prev id : %v\n", detailed.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", detailed.Code)
// 	// 	// fmt.Printf("prev version : %v\n", detailed.Version)
// 	// 	err = UpdateDetailed(repo, detailed, detailed.Model.ID)
// 	// 	assert.NoError(t, err, "can not update detailed record")

// 	// 	detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
// 	// 	err = UpdateDetailed(repo, detailed, detailed.Model.ID)
// 	// 	// fmt.Printf("new version : %v\n", detailed.Version)
// 	// 	assert.Error(t, err, "expected error indicate the versions are different")

// 	// })

// 	// t.Run("can update detailed record if versions were same", func(t *testing.T) {
// 	// 	detailed, err := createTempDetailed(repo)
// 	// 	assert.NoError(t, err, "can not create detailed record due to")

// 	// 	detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
// 	// 	// fmt.Printf("prev id : %v\n", detailed.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", detailed.Code)
// 	// 	// fmt.Printf("prev version : %v\n", detailed.Version)
// 	// 	err = UpdateDetailed(repo, detailed, detailed.Model.ID)
// 	// 	assert.NoError(t, err, "can not update detailed record")

// 	// 	detailed, _ = ReadRecord[models.Detailed](repo, detailed.Model.ID)
// 	// 	detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
// 	// 	err = UpdateDetailed(repo, detailed, detailed.Model.ID)
// 	// 	fmt.Printf("new version : %v\n", detailed.Version)
// 	// 	assert.NoError(t, err, "expected no error")

// 	// })

// 	// t.Run("can not update detailed record if were reffrenced in some voucher items", func(t *testing.T) {
// 	// 	detailed, err := createTempDetailed(repo)
// 	// 	assert.NoError(t, err, "can not create detailed record due to")
// 	// 	fmt.Println("in me ")
// 	// 	fmt.Printf("detialed id : %v\n", detailed.Model.ID)
// 	// 	_, err = createTempVoucher(repo, detailed.Model.ID)

// 	// 	fmt.Printf("detialed id : %v\n", detailed.Model.ID)
// 	// 	assert.NoError(t, err, "can not create voucher record")
// 	// 	// fmt.Printf("voucher id : %v\n", voucher.Model.ID)
// 	// 	detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
// 	// 	err = UpdateDetailed(repo, detailed, detailed.Model.ID)
// 	// 	assert.Error(t, err, "expected error indicate violation update forign key constraint")
// 	// })

// }

// func TestUpdateSubsidiary(t *testing.T) {

// 	repo, err := createConnectionForTest()
// 	defer func() {
// 		sqlDB, _ := repo.AccountingDB.DB()
// 		sqlDB.Close()
// 	}()
// 	if err != nil {
// 		t.Fatalf("can not connect to database %v", err)
// 	}
// 	t.Run("can update subsidiary record successfully", func(t *testing.T) {

// 		subsidiary, err := createTempSubsidiary(repo)
// 		assert.NoError(t, err, "can not create subsidiary record")

// 		prevSubsidiaryId := subsidiary.Model.ID
// 		code := generateUniqeCode[models.Subsidiary](repo, "code")
// 		title := generateUniqeTitle[models.Subsidiary](repo)

// 		subsidiary = &models.Subsidiary{Code: code, Title: title, HasDetailed: true}
// 		err = UpdateSubsidiary(repo, subsidiary, prevSubsidiaryId)
// 		assert.NoError(t, err, "expected no error")
// 	})

// 	// t.Run("return error when update subsidiary record that is not in databse", func(t *testing.T) {
// 	// 	subsidiary, err := createTempSubsidiary(repo)
// 	// 	assert.NoError(t, err, "can not create subsidiary record")

// 	// 	err = UpdateSubsidiary(repo, subsidiary, 1_000_000)
// 	// 	assert.Error(t, err, "expected error indicate there is such id in database")

// 	// })

// 	// t.Run("can not update subsidiary record if versions were  different", func(t *testing.T) {
// 	// 	subsidiary, err := createTempSubsidiary(repo)
// 	// 	assert.NoError(t, err, "can not create subsidiary record due to")

// 	// 	subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
// 	// 	// fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", subsidiary.Code)
// 	// 	// fmt.Printf("prev version : %v\n", subsidiary.Version)
// 	// 	UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
// 	// 	subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
// 	// 	err = UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
// 	// 	// fmt.Printf("new version : %v\n", subsidiary.Version)
// 	// 	assert.Error(t, err, "expected error indicate the versions are different")

// 	// })

// 	// t.Run("can update subsidiary record if versions were same", func(t *testing.T) {
// 	// 	subsidiary, err := createTempSubsidiary(repo)
// 	// 	assert.NoError(t, err, "can not create subsidiary record due to")

// 	// 	subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
// 	// 	// fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", subsidiary.Code)
// 	// 	// fmt.Printf("prev version : %v\n", subsidiary.Version)
// 	// 	UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
// 	// 	subsidiary, _ = ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
// 	// 	subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
// 	// 	err = UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
// 	// 	// fmt.Printf("new version : %v\n", subsidiary.Version)
// 	// 	assert.NoError(t, err, "expected no error")

// 	// })

// }

// func TestUpdateVoucher(t *testing.T) {

// 	repo, err := createConnectionForTest()
// 	defer func() {
// 		sqlDB, _ := repo.AccountingDB.DB()
// 		sqlDB.Close()
// 	}()
// 	if err != nil {
// 		t.Fatalf("can not connect to database %v", err)
// 	}
// 	t.Run("can update voucher number record successfully", func(t *testing.T) {
// 		// code := randgenerator.GenerateRandomCode()
// 		voucher, err := createTempVoucher(repo)
// 		assert.NoError(t, err, "can not create voucher record")

// 		// fmt.Printf("prev Code %v\n", voucher.Number)
// 		prevVoucherId := voucher.Model.ID
// 		code := generateUniqeCode[models.Voucher](repo, "number")

// 		// fmt.Printf("new Code %v\n", code)
// 		voucher.Number = code
// 		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, prevVoucherId)
// 		assert.NoError(t, err, "can not update voucher ")
// 	})

// 	t.Run("can update voucher record successfully by updating , deleting and creating voucherItems ", func(t *testing.T) {
// 		// code := randgenerator.GenerateRandomCode()
// 		voucher, err := createTempVoucher(repo)
// 		assert.NoError(t, err, "can not create voucher record")

// 		// fmt.Printf("prev Code %v\n", voucher.Number)
// 		prevVoucherId := voucher.Model.ID

// 		newVoucherItem, err := createTempVoucherItem(repo)
// 		assert.NoError(t, err, "can not create voucher item record")
// 		temp := append(voucher.VoucherItems, newVoucherItem)
// 		temp[1].Credit = 12

// 		err = UpdateVoucher(repo, voucher, []*models.VoucherItem{temp[1]}, []*models.VoucherItem{temp[0]}, []*models.VoucherItem{temp[2]}, prevVoucherId)

// 		// fmt.Printf("new Code %v\n", code)
// 		assert.NoError(t, err, "can not update voucher ")
// 		_, err = ReadRecord[models.VoucherItem](repo, voucher.VoucherItems[0].Model.ID)
// 		assert.Error(t, err, "expected error indicate voucher item not found")

// 		_, err = ReadRecord[models.VoucherItem](repo, voucher.VoucherItems[1].Model.ID)
// 		assert.NoError(t, err, "expexted no error when reading the voucherItem record")
// 		// assert.Equal(t, voucher.VoucherItems[1].DetailedId, newVoucherItem.DetailedId)
// 		// assert.Equal(t, voucher.VoucherItems[1].SubsidiaryId, newVoucherItem.SubsidiaryId)
// 		// assert.Equal(t, voucher.VoucherItems[1].Debit, newVoucherItem.Debit)
// 		// assert.Equal(t, voucher.VoucherItems[1].Credit, newVoucherItem.Credit)
// 		// assert.Equal(t, voucher.VoucherItems[1].VoucherID, newVoucherItem.VoucherID)

// 		_, err = ReadRecord[models.VoucherItem](repo, voucher.VoucherItems[2].Model.ID)
// 		assert.NoError(t, err, "expexted no error when reading the voucherItem record")

// 	})

// 	// t.Run("return error when update voucher record that is not in databse", func(t *testing.T) {

// 	// 	voucher := &models.Voucher{}
// 	// 	voucher.Model.ID = 1_000_000
// 	// 	err := UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, 1_000_000)
// 	// 	assert.Error(t, err, "expected error indicate there is such id in database")

// 	// })

// 	// t.Run("can not update voucher record if versions were  different", func(t *testing.T) {
// 	// 	voucher, err := createTempVoucher(repo)
// 	// 	assert.NoError(t, err, "can not create voucher record")

// 	// 	voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
// 	// 	// fmt.Printf("prev id : %v\n", voucher.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", voucher.Number)
// 	// 	// fmt.Printf("prev version : %v\n", voucher.Version)
// 	// 	err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)
// 	// 	assert.NoError(t, err, "can not update voucher record")

// 	// 	voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
// 	// 	err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)

// 	// 	// fmt.Printf("new version : %v\n", voucher.Version)
// 	// 	assert.Error(t, err, "expected error indicate the versions are different")

// 	// })

// 	// t.Run("can update voucher record if versions were same", func(t *testing.T) {
// 	// 	voucher, err := createTempVoucher(repo)
// 	// 	assert.NoError(t, err, "can not create voucher record")

// 	// 	voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
// 	// 	// fmt.Printf("prev id : %v\n", voucher.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", voucher.Number)
// 	// 	// fmt.Printf("prev version : %v\n", voucher.Version)
// 	// 	err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)
// 	// 	assert.NoError(t, err, "can not update voucher record")

// 	// 	voucher, _ = ReadRecord[models.Voucher](repo, voucher.Model.ID)
// 	// 	voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
// 	// 	err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)

// 	// 	// fmt.Printf("new version : %v\n", voucher.Version)
// 	// 	assert.NoError(t, err, "expected no error")

// 	// })
// }

// func TestDeleteDetailed(t *testing.T) {
// 	repo, err := createConnectionForTest()

// 	defer func() {
// 		sqlDB, _ := repo.AccountingDB.DB()
// 		sqlDB.Close()
// 	}()
// 	if err != nil {
// 		t.Fatalf("can not connect to database %v", err)
// 	}

// 	// t.Run("delete detailed record seccessfully", func(t *testing.T) {
// 	// 	detailed, err := createTempDetailed(repo)
// 	// 	assert.NoError(t, err, "can not create detailed record due to")

// 	// 	err = DeleteDetailedRecord(repo, detailed)
// 	// 	assert.NoError(t, err, "can not delete detailed record")

// 	// 	result := repo.AccountingDB.First(&detailed)
// 	// 	assert.Error(t, result.Error, "expected error indicate detailed record not found")

// 	// })

// 	// t.Run("deletion detailed record fail because record does not exist in database", func(t *testing.T) {
// 	// 	detailed := &models.Detailed{}
// 	// 	detailed.Model.ID = 1_000_000
// 	// 	err := DeleteDetailedRecord(repo, detailed)
// 	// 	// result := repo.AccountingDB.First(&detailed)
// 	// 	assert.Error(t, err, "expected error indicate detailed record not found")
// 	// })

// 	// t.Run("deletion detailed record fails because there is a reffrence in some voucher items  ", func(t *testing.T) {
// 	// 	detailed, err := createTempDetailed(repo)
// 	// 	assert.NoError(t, err, "can not create detailed record due to")

// 	// 	voucher, err := createTempVoucher(repo, detailed.Model.ID)
// 	// 	assert.NoError(t, err, "can not create detailed record ")
// 	// 	fmt.Printf("det : %v", detailed.Model.ID)
// 	// 	fmt.Printf("vi : %v", voucher.VoucherItems[0].Model.ID)
// 	// 	err = DeleteDetailedRecord(repo, detailed)

// 	// 	assert.Error(t, err, "expected error indicate violation forignkey constraint")

// 	// })

// 	// t.Run("can not delete detailed record if versions were  different", func(t *testing.T) {
// 	// 	detailed, err := createTempDetailed(repo)
// 	// 	assert.NoError(t, err, "can not create detailed record due to")

// 	// 	detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
// 	// 	// fmt.Printf("prev id : %v\n", detailed.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", detailed.Code)
// 	// 	// fmt.Printf("prev version : %v\n", detailed.Version)
// 	// 	UpdateDetailed(repo, detailed, detailed.Model.ID)
// 	// 	err = DeleteDetailedRecord(repo, detailed)
// 	// 	// fmt.Printf("new version : %v\n", detailed.Version)
// 	// 	assert.Error(t, err, "expected error indicate the versions are different")

// 	// })

// 	// t.Run("can delete detailed record if versions were same", func(t *testing.T) {
// 	// 	detailed, err := createTempDetailed(repo)
// 	// 	assert.NoError(t, err, "can not create detailed record due to")

// 	// 	detailed.Code = generateUniqeCode[models.Detailed](repo, "code")
// 	// 	// fmt.Printf("prev id : %v\n", detailed.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", detailed.Code)
// 	// 	// fmt.Printf("prev version : %v\n", detailed.Version)
// 	// 	err = UpdateDetailed(repo, detailed, detailed.Model.ID)
// 	// 	assert.NoError(t, err, "can not update detailed record ")
// 	// 	detailed, _ = ReadRecord[models.Detailed](repo, detailed.Model.ID)

// 	// 	err = DeleteDetailedRecord(repo, detailed)
// 	// 	// fmt.Printf("new version : %v\n", detailed.Version)
// 	// 	assert.NoError(t, err, "expected no error")

// 	// })
// }

// func TestDeleteSubsidiary(t *testing.T) {
// 	repo, err := createConnectionForTest()
// 	defer func() {
// 		sqlDB, _ := repo.AccountingDB.DB()
// 		sqlDB.Close()
// 	}()
// 	if err != nil {
// 		t.Fatalf("can not connect to database %v", err)
// 	}

// 	// t.Run("delete subsidiary record seccessfully", func(t *testing.T) {
// 	// 	subsidiary, err := createTempSubsidiary(repo)
// 	// 	assert.NoError(t, err, "can not create subsidiary record due to")

// 	// 	err = DeleteSubsidiaryRecord(repo, subsidiary)

// 	// 	assert.NoError(t, err, "can not delete subsidiary record due to ")

// 	// })

// 	// t.Run("deletion subsidiary record fail because record does not exist in database", func(t *testing.T) {
// 	// 	subsidiary := &models.Subsidiary{}
// 	// 	subsidiary.Model.ID = 1_000_000
// 	// 	err := DeleteSubsidiaryRecord(repo, subsidiary)

// 	// 	assert.Error(t, err, "expected error indicate subsiduary record not found")
// 	// })

// 	// t.Run("deletion subsidiary record fails because there is a reffrence in some voucher items  ", func(t *testing.T) {
// 	// 	subsidiary, err := createTempSubsidiary(repo)
// 	// 	assert.NoError(t, err, "can not create subsidiary record due to")

// 	// 	_, err = createTempVoucher(repo, 0, subsidiary.Model.ID)
// 	// 	assert.NoError(t, err, "can not create voucher item record")
// 	// 	// fmt.Printf("det : %v", subsidiary.Model.ID)
// 	// 	// fmt.Printf("vi : %v", voucher.VoucherItems[0].Model.ID)
// 	// 	err = DeleteSubsidiaryRecord(repo, subsidiary)

// 	// 	assert.Error(t, err, "expected error indicate violation forignkey constraint")

// 	// })

// 	// t.Run("can not delete subsidiary record if versions were  different", func(t *testing.T) {
// 	// 	subsidiary, err := createTempSubsidiary(repo)
// 	// 	assert.NoError(t, err, "can not create subsidiary record due to")

// 	// 	subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
// 	// 	// fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", subsidiary.Code)
// 	// 	// fmt.Printf("prev version : %v\n", subsidiary.Version)
// 	// 	err = UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
// 	// 	assert.NoError(t, err, "can not update subsidiary record ")
// 	// 	err = DeleteSubsidiaryRecord(repo, subsidiary)
// 	// 	// fmt.Printf("new version : %v\n", subsidiary.Version)
// 	// 	assert.Error(t, err, "expected error indicate the versions are different")

// 	// })

// 	// t.Run("can delete subsidiary record if versions were same", func(t *testing.T) {
// 	// 	subsidiary, err := createTempSubsidiary(repo)
// 	// 	assert.NoError(t, err, "can not create subsidiary record due to")

// 	// 	subsidiary.Code = generateUniqeCode[models.Subsidiary](repo, "code")
// 	// 	// fmt.Printf("prev id : %v\n", subsidiary.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", subsidiary.Code)
// 	// 	// fmt.Printf("prev version : %v\n", subsidiary.Version)
// 	// 	err = UpdateSubsidiary(repo, subsidiary, subsidiary.Model.ID)
// 	// 	assert.NoError(t, err, "cann not update subsidiary record due to ")

// 	// 	subsidiary, _ = ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)

// 	// 	err = DeleteSubsidiaryRecord(repo, subsidiary)
// 	// 	// fmt.Printf("new version : %v\n", subsidiary.Version)
// 	// 	assert.NoError(t, err, "expected no error")

// 	// })
// }

// func TestDeleteVoucher(t *testing.T) {
// 	repo, err := createConnectionForTest()
// 	defer func() {
// 		sqlDB, _ := repo.AccountingDB.DB()
// 		sqlDB.Close()
// 	}()
// 	if err != nil {
// 		t.Fatalf("can not connect to database %v", err)
// 	}

// 	// t.Run("delete voucher record seccessfully", func(t *testing.T) {

// 	// 	voucher, err := createTempVoucher(repo)
// 	// 	assert.NoError(t, err, "can not create voucher record")

// 	// 	// fmt.Printf("voucher : %v", voucher.Model.ID)
// 	// 	err = DeleteVoucherRecord(repo, voucher)
// 	// 	assert.NoError(t, err, "expected no error %v", err)
// 	// 	result := repo.AccountingDB.First(&voucher)
// 	// 	assert.Error(t, result.Error, "expected error indicate voucher record not found")

// 	// })

// 	// t.Run("deletion voucher record fail because record does not exist in database", func(t *testing.T) {
// 	// 	voucher := &models.Voucher{}
// 	// 	voucher.Model.ID = 1_000_000
// 	// 	err = DeleteVoucherRecord(repo, voucher)
// 	// 	assert.Error(t, err, "expected error indicate there is not such record in data base")

// 	// })

// 	// t.Run("can not delete voucher record if versions were  different", func(t *testing.T) {
// 	// 	voucher, err := createTempVoucher(repo)
// 	// 	assert.NoError(t, err, "can not create voucher record")

// 	// 	voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
// 	// 	// fmt.Printf("prev id : %v\n", voucher.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", voucher.Number)
// 	// 	// fmt.Printf("prev version : %v\n", voucher.Version)
// 	// 	err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)
// 	// 	assert.NoError(t, err, "can not update voucher record")

// 	// 	err = DeleteVoucherRecord(repo, voucher)
// 	// 	// fmt.Printf("new version : %v\n", voucher.Version)
// 	// 	assert.Error(t, err, "expected error indicate the versions are different")

// 	// })

// 	// t.Run("can delete voucher record if versions were same", func(t *testing.T) {
// 	// 	voucher, err := createTempVoucher(repo)
// 	// 	assert.NoError(t, err, "can not create voucher record")

// 	// 	voucher.Number = generateUniqeCode[models.Voucher](repo, "number")
// 	// 	// fmt.Printf("prev id : %v\n", voucher.Model.ID)
// 	// 	// fmt.Printf("code : %v\n", voucher.Number)
// 	// 	// fmt.Printf("prev version : %v\n", voucher.Version)
// 	// 	err = UpdateVoucher(repo, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{}, voucher.Model.ID)
// 	// 	assert.NoError(t, err, "can not update voucher record")
// 	// 	voucher, _ = ReadRecord[models.Voucher](repo, voucher.Model.ID)

// 	// 	err = DeleteVoucherRecord(repo, voucher)
// 	// 	// fmt.Printf("new version : %v\n", voucher.Version)
// 	// 	assert.NoError(t, err, "expected no error")

// 	// })
// }

// func TestReadRecord(t *testing.T) {
// 	repo, err := createConnectionForTest()
// 	defer func() {
// 		sqlDB, _ := repo.AccountingDB.DB()
// 		sqlDB.Close()
// 	}()
// 	if err != nil {
// 		t.Fatalf("can not connect to database %v", err)
// 	}

// 	t.Run("can read the detailed record successfully", func(t *testing.T) {
// 		detailed, err := createTempDetailed(repo)
// 		assert.NoError(t, err, "can not create detailed record due to")

// 		res, err := ReadRecord[models.Detailed](repo, detailed.Model.ID)
// 		assert.NoError(t, err, "expected no error")
// 		assert.Equal(t, res.Code, detailed.Code)
// 		assert.Equal(t, res.Title, detailed.Title)
// 	})

// 	t.Run("return error when the detailed record is not in database ", func(t *testing.T) {

// 		_, err := ReadRecord[models.Detailed](repo, 1_000_000)
// 		assert.Error(t, err, "expected  error indicate can not found the detailed record")

// 	})

// 	t.Run("can read the subsidiary record successfully", func(t *testing.T) {
// 		subsidiary, err := createTempSubsidiary(repo)
// 		assert.NoError(t, err, "can not create subsidiary record due to")

// 		res, err := ReadRecord[models.Subsidiary](repo, subsidiary.Model.ID)
// 		assert.NoError(t, err, "expected no error")
// 		assert.Equal(t, res.Code, subsidiary.Code)
// 		assert.Equal(t, res.Title, subsidiary.Title)
// 		assert.Equal(t, res.HasDetailed, subsidiary.HasDetailed)
// 	})

// 	t.Run("return error when the subsidiary record is not in database ", func(t *testing.T) {

// 		_, err := ReadRecord[models.Subsidiary](repo, 1_000_000)
// 		assert.Error(t, err, "expected  error indicate can not found the subsidiary record")

// 	})

// 	t.Run("can read the voucher record successfully", func(t *testing.T) {
// 		voucher, err := createTempVoucher(repo)
// 		assert.NoError(t, err, "can not create voucher record")

// 		res, err := ReadRecord[models.Voucher](repo, voucher.Model.ID)
// 		assert.NoError(t, err, "expected no error")
// 		assert.Equal(t, res.Number, res.Number)

// 	})

// 	t.Run("return error when the voucher record is not in database ", func(t *testing.T) {

// 		_, err := ReadRecord[models.Voucher](repo, 1_000_000)
// 		assert.Error(t, err, "expected  error indicate can not found the voucher record")

// 	})

// }

// func createTempVoucher(repo *Repositories, IDs ...uint) (*models.Voucher, error) {
// 	temp := make([]*models.VoucherItem, 3)
// 	if len(IDs) == 0 {

// 		detailed, err := createTempDetailed(repo)
// 		if err != nil {
// 			return nil, err
// 		}

// 		subsidiary, err := createTempSubsidiary(repo)
// 		if err != nil {
// 			return nil, err
// 		}

// 		temp[0] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Credit: 500}

// 		temp[1] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Debit: 250}

// 		temp[2] = &models.VoucherItem{DetailedId: detailed.Model.ID, Debit: 250}
// 	} else {
// 		temp = make([]*models.VoucherItem, 2)

// 		if len(IDs) == 1 {
// 			subsidiary, err := createTempSubsidiary(repo)
// 			if err != nil {
// 				return nil, err
// 			}
// 			temp[0] = &models.VoucherItem{DetailedId: IDs[0], Credit: 500}
// 			temp[1] = &models.VoucherItem{DetailedId: IDs[0], SubsidiaryId: subsidiary.Model.ID, Debit: 500}
// 		} else {
// 			detailed, err := createTempDetailed(repo)
// 			if err != nil {
// 				return nil, err
// 			}
// 			temp[0] = &models.VoucherItem{DetailedId: detailed.Model.ID, Credit: 500}
// 			temp[1] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: IDs[1], Debit: 500}

// 		}

// 	}

// 	number := generateUniqeCode[models.Voucher](repo, "number")
// 	voucher := &models.Voucher{Number: number, VoucherItems: temp}

// 	// err := errors.New("")
// 	// for err != nil {

// 	err := CreateRecord(repo, voucher)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error during record creation: %v\n", err)

// 	}

// 	// }

// 	return voucher, nil
// }

// func createTempVoucherItem(repo *Repositories) (*models.VoucherItem, error) {
// 	detailed, err := createTempDetailed(repo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	subsidiary, err := createTempSubsidiary(repo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Debit: 250}, nil
// }

// func createTempSubsidiary(repo *Repositories) (*models.Subsidiary, error) {
// 	subsidiary := &models.Subsidiary{Code: generateUniqeCode[models.Subsidiary](repo, "code"), Title: generateUniqeTitle[models.Subsidiary](repo), HasDetailed: false}

// 	err := CreateRecord(repo, subsidiary)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error during record creation: %v\n", err)

// 	}

// 	return subsidiary, nil
// }

// func createTempDetailed(repo *Repositories) (*models.Detailed, error) {

// 	detailed := &models.Detailed{Code: generateUniqeCode[models.Detailed](repo, "code"), Title: generateUniqeTitle[models.Detailed](repo)}

// 	// err := errors.New("")
// 	// for err != nil {

// 	err := CreateRecord(repo, detailed)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error during record creation: %v\n", err)

// 	}

// 	return detailed, nil

// 	// }
// }

// func generateUniqeCode[T any](repo *Repositories, columnName string) string {
// 	code := randgenerator.GenerateRandomCode()
// 	for {
// 		exist := FindRecord[T](repo, code, columnName)

// 		if exist {
// 			code = randgenerator.GenerateRandomCode()
// 		} else {
// 			break
// 		}
// 	}

// 	return code
// }

// func generateUniqeTitle[T any](repo *Repositories) string {
// 	title := randgenerator.GenerateRandomTitle()
// 	for {
// 		exist := FindRecord[T](repo, title, "title")

// 		if exist {
// 			title = randgenerator.GenerateRandomTitle()
// 		} else {
// 			break
// 		}
// 	}

// 	return title
// }
