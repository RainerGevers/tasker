package models

import (
	"gorm.io/gorm"
)

type Version struct {
	gorm.Model
	Uuid    string `gorm:"index:,unique;not null"`
	Version string `gorm:"index:,unique"`
}
