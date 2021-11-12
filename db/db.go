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
  db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=gorm password=12345 sslmode=disable")

  if err != nil {
    fmt.Println("Failed to connect to database!")
    panic(err)
  }

  db.AutoMigrate(&models.User{})

  DB = db
}
