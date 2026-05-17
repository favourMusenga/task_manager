package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDb(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB(dbPath string) error {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		return err
	}

	// Migrate the schema
	db.AutoMigrate(&Category{}, &Todo{}, &FocusTimer{})
	return nil
}
