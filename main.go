package main

import (
  "github.com/gin-gonic/gin"
  "github.com/distributed-marketplace-system/controllers"
  "github.com/distributed-marketplace-system/db"
  "github.com/joho/godotenv"
  "github.com/distributed-marketplace-system/util"
  _ "fmt"
)

func main() {
  err := godotenv.Load(".env")

  if err != nil {
    panic("Fatal Error: Couldn't loading .env file")
  }

  router := gin.Default()

  db.ConnectDatabase()

  router.GET("/users", auth.AuthMiddleware(), controllers.GetUsers)
  router.POST("/users", controllers.CreateUser)

  router.Run("localhost:8080")
}
