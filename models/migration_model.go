package models

import "gorm.io/gorm"

type Migration struct {
	gorm.Model
	Name string `gorm:"index:,unique"`
}
