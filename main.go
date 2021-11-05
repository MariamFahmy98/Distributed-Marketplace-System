package main

import (
	"github.com/gin-gonic/gin"

	"github.com/distributed-marketplace-system/controllers"

	"github.com/distributed-marketplace-system/db"
)

func main() {
	router := gin.Default()

	db.ConnectDatabase()

	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)

	router.Run("localhost:8080")
}
