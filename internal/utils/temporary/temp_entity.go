package temporary

import (
	"accounting_system/internal/models"
	"accounting_system/internal/repositories"

	"fmt"
)

func CreateTempVoucher(repo *repositories.Repositories, IDs ...uint) (*models.Voucher, error) {
	temp := make([]*models.VoucherItem, 4)

	subsidiary := &models.Subsidiary{Code: repositories.GenerateUniqeCode[models.Subsidiary](repo, "code"), Title: repositories.GenerateUniqeTitle[models.Subsidiary](repo), HasDetailed: true}
	err := repositories.CreateRecord(repo, subsidiary)
	if err != nil {
		return nil, err
	}

	if len(IDs) == 0 {

		detailed, err := CreateTempDetailed(repo)

		if err != nil {
			return nil, err
		}

		temp[0] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Credit: 250}

		temp[1] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Credit: 250}

		temp[2] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Debit: 250}
		temp[3] = &models.VoucherItem{DetailedId: detailed.Model.ID, SubsidiaryId: subsidiary.Model.ID, Debit: 250}
	} else {
		temp = make([]*models.VoucherItem, 2)

		if len(IDs) == 1 {

			temp[0] = &models.VoucherItem{DetailedId: IDs[0], SubsidiaryId: subsidiary.Model.ID, Credit: 500}
			temp[1] = &models.VoucherItem{DetailedId: IDs[0], SubsidiaryId: subsidiary.Model.ID, Debit: 500}
		} else {

			temp[0] = &models.VoucherItem{SubsidiaryId: IDs[1], Credit: 500}
			temp[1] = &models.VoucherItem{SubsidiaryId: IDs[1], Debit: 500}

		}

	}

	number := repositories.GenerateUniqeCode[models.Voucher](repo, "number")
	voucher := &models.Voucher{Number: number, VoucherItems: temp}



	err = repositories.CreateRecord(repo, voucher)
	if err != nil {
		return nil, fmt.Errorf("error during record creation: %v", err)

	}


	return voucher, nil
}

func ReturnTempVoucherItem(repo *repositories.Repositories) (*models.VoucherItem, error) {

	subsidiary, err := CreateTempSubsidiary(repo)
	if err != nil {
		return nil, err
	}

	return &models.VoucherItem{SubsidiaryId: subsidiary.Model.ID, Debit: 250}, nil
}

func CreateTempSubsidiary(repo *repositories.Repositories) (*models.Subsidiary, error) {
	subsidiary := &models.Subsidiary{Code: repositories.GenerateUniqeCode[models.Subsidiary](repo, "code"), Title: repositories.GenerateUniqeTitle[models.Subsidiary](repo), HasDetailed: false}

	err := repositories.CreateRecord(repo, subsidiary)
	if err != nil {
		return nil, fmt.Errorf("error during record creation: %v", err)

	}

	return subsidiary, nil
}

func CreateTempDetailed(repo *repositories.Repositories) (*models.Detailed, error) {

	detailed := &models.Detailed{Code: repositories.GenerateUniqeCode[models.Detailed](repo, "code"), Title: repositories.GenerateUniqeTitle[models.Detailed](repo)}

	err := repositories.CreateRecord(repo, detailed)
	if err != nil {
		return nil, fmt.Errorf("error during record creation: %v", err)

	}

	return detailed, nil

}
