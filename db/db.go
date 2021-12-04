package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRESQL_ADDRESS"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database!")
		panic(err)
	}

	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Product{})
	// db.AutoMigrate(&models.Order{})
	// db.AutoMigrate(&models.Store{})
	// db.AutoMigrate(&models.Deposit{})

	DB = db
}
