package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)
import "gorm.io/driver/mysql"

func NewConnection() (*gorm.DB, error) {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	hostname := os.Getenv("MYSQL_HOSTNAME")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	dsn := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}