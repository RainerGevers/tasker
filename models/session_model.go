package models

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	Uuid      string `gorm:"index:,unique;not null"`
	UserId    uint
	ExpiresAt time.Time

	User User `gorm:"foreignKey:UserId"`
}
