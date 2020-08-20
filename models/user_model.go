package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index:,unique"`
	Password string
	Username string    `gorm:"index:,unique"`
	Uuid     string    `gorm:"index:,unique"`
	Sessions []Session `gorm:"foreignKey:UserId"`
}
