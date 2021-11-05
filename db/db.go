package db

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/distributed-marketplace-system/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=gorm password=12345 sslmode=disable")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.User{})

	DB = db
}
