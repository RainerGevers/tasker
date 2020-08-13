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
	RefreshCode string	`gorm:"index;not null"`

	User User `gorm:"foreignKey:UserId"`
}
