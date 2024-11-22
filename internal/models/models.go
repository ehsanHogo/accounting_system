package models

import (
	"gorm.io/gorm"
)

type Detailed struct {
	ID        uint           `gorm:"primaryKey"`
	Code    string `gorm:"unique"`
	Title   string `gorm:"unique"`
	Version uint   `gorm:"default:0"`
}

type Subsidiary struct {
	ID        uint           `gorm:"primaryKey"`
	Code        string `gorm:"unique"`
	Title       string `gorm:"unique"`
	HasDetailed bool
	Version     uint `gorm:"default:0"`
}

type VoucherItem struct {
	ID        uint           `gorm:"primaryKey"`
	VoucherID uint `gorm:"not null;constraint:OnDelete:CASCADE;"`

	DetailedId   uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;default:null;"`
	SubsidiaryId uint `gorm:"not null;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Debit        int64
	Credit       int64
}

type Voucher struct {
	ID        uint           `gorm:"primaryKey"`
	Number       string         `gorm:"unique"`
	VoucherItems []*VoucherItem `gorm:"foreignKey:VoucherID"`
	Version      uint           `gorm:"default:0"`
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
