package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
import "gorm.io/driver/mysql"

func NewConnection() (*gorm.DB, interface{}) {
	dsn := "root:root@tcp(127.0.0.1:3306)/tasker_dev?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, "Could not connect to database!!!"
	}

	return db, nil
}