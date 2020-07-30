package db

import (
	"github.com/RainerGevers/tasker/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}