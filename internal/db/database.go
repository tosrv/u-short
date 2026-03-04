package db

import (
	"u-short/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(dbUrl string) *gorm.DB {
	var err error

	database, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	database.AutoMigrate(&model.Url{})
	return database
}
