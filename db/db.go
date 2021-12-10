package db

import (
	"distributed-marketplace-system/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRESQL_ADDRESS"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Failed to connect to database!")
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Store{})
	db.AutoMigrate(&models.Transaction{})

	DB = db
}
