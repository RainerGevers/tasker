package db

import (
	"github.com/RainerGevers/tasker/db/migrations"
	"github.com/RainerGevers/tasker/models"
	"gorm.io/gorm"
	"log"
)

func RunMigrations(db *gorm.DB) {
	_ = db.AutoMigrate(&models.Version{}, &models.User{})
	var versions []models.Version
	dbVersions := db.Select("version").Find(&versions)
	if dbVersions.Error != nil {
		log.Fatal(dbVersions.Error)
	}

	// TODO: Pluck https://v2.gorm.io/docs/advanced_query.html
	var versionNumbers []string
	for _, version := range versions {
		versionNumbers = append(versionNumbers, version.Version)
	}
	versionsToMigrate := difference(availableMigrations(), versionNumbers)
	if len(versionsToMigrate) == 0 {
		return
	}

	for _, versionToMigrate := range versionsToMigrate {
		switch versionToMigrate {
		case "13082020175800":
			migrations.CreateSessionsTable13082020175800(db)
		case "13082020205200":
			migrations.AddRefreshToSessionsTable13082020205200(db)

		}
	}

}

func availableMigrations() []string {
	return []string{
		"13082020175800",
		"13082020205200",
	}
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string

	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, s1)
		}
	}

	return diff
}
