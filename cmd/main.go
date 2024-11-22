package main

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"
	"accounting_system/internal/servieces/detailedserv"
	"accounting_system/internal/servieces/subsidiaryserv"
	"accounting_system/internal/servieces/voucherserv"
	"accounting_system/internal/utils/randgenerator"
	"fmt"
)

func main() {

	repo, err := repositories.CreateConnectionForTest()
	// err = errors.New("")

	if err != nil {
		fmt.Printf("can not connect to database due to : %v\n", err)
		return
	}

	detailed := &models.Detailed{Code: randgenerator.GenerateRandomCode(), Title: randgenerator.GenerateRandomTitle()}
	err = detailedserv.InsertDetailed(repo.AccountingDB, detailed)
	if err != nil {
		fmt.Printf("can not insert detailed due to : %v\n", err)
		return
	}
	subsidiary := &models.Subsidiary{Code: randgenerator.GenerateRandomCode(), Title: randgenerator.GenerateRandomTitle(), HasDetailed: true}
	err = subsidiaryserv.InsertSubsidiary(repo.AccountingDB, subsidiary)
	if err != nil {
		fmt.Printf("can not insert subsidiary due to : %v\n", err)
		return
	}

	voucher := &models.Voucher{Number: randgenerator.GenerateRandomCode()}
	voucher.VoucherItems = []*models.VoucherItem{{DetailedId: detailed.ID, SubsidiaryId: subsidiary.ID, Credit: 400}, {DetailedId: detailed.ID, SubsidiaryId: subsidiary.ID, Debit: 200}, {DetailedId: detailed.ID, SubsidiaryId: subsidiary.ID, Debit: 200}}
	// repositories.CreateRecord(repo.AccountingDB, detailed)
	err = voucherserv.InsertVoucher(repo.AccountingDB, voucher)
	if err != nil {
		fmt.Printf("can not insert voucher due to : %v\n", err)
		return
	}

	voucher.VoucherItems[0].Credit -= 100
	voucher.VoucherItems[1].Debit += 100

	err = voucherserv.UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{voucher.VoucherItems[0], voucher.VoucherItems[1]}, []int64{voucher.VoucherItems[2].ID}, []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Credit: 200}, {SubsidiaryId: subsidiary.ID, DetailedId: detailed.ID, Debit: 200}})

	if err != nil {
		fmt.Printf("can not update voucher due to : %v\n", err)
		return
	}

	// err = voucherserv.UpdateVoucher(repo.AccountingDB, voucher, []*models.VoucherItem{}, []*models.VoucherItem{}, []*models.VoucherItem{{SubsidiaryId: subsidiary.ID, DetailedId:  detailed.ID, Credit: 200}, {SubsidiaryId: subsidiary.ID, DetailedId:  detailed.ID,Debit: 200}})

}
