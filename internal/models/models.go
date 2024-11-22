package models

import "gorm.io/gorm"

type Detailed struct {
	gorm.Model
	Code    string `gorm:"unique"`
	Title   string `gorm:"unique"`
	Version uint
}

type Subsidiary struct {
	gorm.Model
	Code        string `gorm:"unique"`
	Title       string `gorm:"unique"`
	HasDetailed bool
	Version     uint
}

type VoucherItem struct {
	gorm.Model
	VoucherID uint `gorm:"not null;constraint:OnDelete:CASCADE;"`

	DetailedId   uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;default:null;"`
	SubsidiaryId uint `gorm:"not null;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Debit        int64
	Credit       int64
}

type Voucher struct {
	gorm.Model
	Number       string         `gorm:"unique"`
	VoucherItems []*VoucherItem `gorm:"foreignKey:VoucherID"`
	Version      uint
}

//creare models

type InsertDetailedRequest struct {
	ID    string
	Code  string
	Title string

	Version string
}

type UpdateDetailedRequest struct {
	ID    string
	Code  string
	Title string

	Version string
}

type DeleteDetailedRequest struct {
	ID      string
	Version string
}

type GetDetailedRequest struct {
	ID string
}

type InsertSubsidiaryRequest struct {
	ID      string
	Code    string
	Title   string
	HasDl   bool
	Version string
}

type UpdateSubsidiaryRequest struct {
	ID      string
	Code    string
	Title   string
	HasDl   bool
	Version string
}

type DeleteSubsidiaryRequest struct {
	ID      string
	Version string
}

type GetSubsidiaryRequest struct {
	ID string
}
