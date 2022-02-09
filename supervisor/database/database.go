package database

import (
	"gosupervisor/configuration"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase(configuration configuration.Configuration) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(configuration.SqlitePath), &gorm.Config{})
}
