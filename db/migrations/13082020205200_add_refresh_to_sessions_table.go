package migrations

import (
"github.com/RainerGevers/tasker/models"
uuid "github.com/satori/go.uuid"
"gorm.io/gorm"
)

func AddRefreshToSessionsTable13082020205200(db *gorm.DB) {

		db.Migrator().AddColumn(models.Session{}, "RefreshCode")

		version := models.Version{Uuid: uuid.NewV4().String(), Version: "13082020205200"}
		db.Create(&version)

}

