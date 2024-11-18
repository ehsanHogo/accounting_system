package models

import "gorm.io/gorm"

type Detailed struct {
	gorm.Model
	Code  string `gorm:"unique"`
	Title string `gorm:"unique"`
}

type Subsidiary struct {
	gorm.Model
	Code        string `gorm:"unique"`
	Title       string `gorm:"unique"`
	HasDetailed bool
}

type VoucherItem struct {
	gorm.Model
	VoucherID uint `gorm:"not null;constraint:OnDelete:CASCADE;"`

	DetailedId   uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	SubsidiaryId uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;default:null;"`
	Debit        int64
	Credit       int64
}

type Voucher struct {
	gorm.Model
	Number       string         `gorm:"unique"`
	VoucherItems []*VoucherItem `gorm:"foreignKey:VoucherID"`
}

//creare models
