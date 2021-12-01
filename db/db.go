package db

import (
  "fmt"
  "os"

  "distributed-marketplace-system/models"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
  dbinfo := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
  db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

  if err != nil {
    fmt.Println("Failed to connect to database!")
    panic(err)
  }

  db.AutoMigrate(&models.User{})
  db.AutoMigrate(&models.Product{})
  db.AutoMigrate(&models.Order{})
  db.AutoMigrate(&models.Store{})

  DB = db
}
