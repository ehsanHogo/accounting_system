package repositories

import "gorm.io/gorm"

type Repositories struct {
	DB *gorm.DB
}

func NewConnection(db *gorm.DB) *Repositories {
	return &Repositories{
		DB: db,
	}
}
