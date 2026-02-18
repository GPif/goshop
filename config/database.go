package config

import (
	"goshop/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Item{}); err != nil {
		return nil, err
	}

	return db, nil
}
