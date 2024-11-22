package models

import (
	"gorm.io/gorm"
)

type Detailed struct {
	ID      int64  `gorm:"primaryKey"`
	Code    string `gorm:"unique"`
	Title   string `gorm:"unique"`
	Version int64  `gorm:"default:0"`
}

type Subsidiary struct {
	ID          int64  `gorm:"primaryKey"`
	Code        string `gorm:"unique"`
	Title       string `gorm:"unique"`
	HasDetailed bool
	Version     int64 `gorm:"default:0"`
}

type VoucherItem struct {
	ID        int64 `gorm:"primaryKey"`
	VoucherID int64 `gorm:"not null;constraint:OnDelete:CASCADE;"`

	DetailedId   int64 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;default:null;"`
	SubsidiaryId int64 `gorm:"not null;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Debit        int32
	Credit       int32
}

type Voucher struct {
	ID           int64          `gorm:"primaryKey"`
	Number       string         `gorm:"unique"`
	VoucherItems []*VoucherItem `gorm:"foreignKey:VoucherID"`
	Version      int64          `gorm:"default:0"`
}

func (u *Detailed) BeforeUpdate(tx *gorm.DB) (err error) {

	u.Version++
	return
}

func (u *Subsidiary) BeforeUpdate(tx *gorm.DB) (err error) {

	u.Version++
	return
}

func (u *Voucher) BeforeUpdate(tx *gorm.DB) (err error) {

	u.Version++
	return
}
