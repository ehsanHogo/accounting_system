package models

import "gorm.io/gorm"

type Detailed struct {
	gorm.Model
	Code  string
	Title string
}

type Subsidiary struct {
	gorm.Model
	Code        string
	Title       string
	HasDetailed bool
}

type VoucherItem struct {
	gorm.Model
	VoucherID    uint
	DatailedId   int64
	SubsidiaryId int64
	debit        int64
	credit       int64
}

type Voucher struct {
	gorm.Model
	Number       string
	VoucherItems []VoucherItem
}

//creare models
