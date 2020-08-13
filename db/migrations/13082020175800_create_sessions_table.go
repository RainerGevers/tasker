package migrations

import (
	"github.com/RainerGevers/tasker/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func CreateSessionsTable13082020175800(db *gorm.DB){

	if !db.Migrator().HasTable(models.Session{}) {
		db.Migrator().CreateTable(models.Session{})

		version := models.Version{Uuid: uuid.NewV4().String(), Version: "13082020175800"}
		db.Create(&version)
	}

}