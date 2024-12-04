package config

import (
	"github.com/TranMinh2k3/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgress"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}