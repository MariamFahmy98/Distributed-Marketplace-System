package db

import (
	"fmt"
	"os"

	"github.com/distributed-marketplace-system/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	db, err := gorm.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println("Failed to connect to database!")
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})

  DB = db
}
