package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB initializes the database connection
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("imgur_clone.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
