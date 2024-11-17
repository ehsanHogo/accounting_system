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
	DatailedId   uint 
	SubsidiaryId uint
	Debit        int64
	Credit       int64
}

type Voucher struct {
	gorm.Model
	Number       string
	VoucherItems []VoucherItem
}

//creare models
