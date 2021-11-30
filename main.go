package main

import (
	"log"
	"net/http"
	"os"

	"distributed-marketplace-system/controllers"
	"distributed-marketplace-system/db"
	"distributed-marketplace-system/errors"
	"distributed-marketplace-system/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func routes() {
  router := gin.Default()

  // User APIs
  user_r := router.Group("/users")
  {
    user := new(controllers.UserController)

    user_r.GET("", user.GetAll)
    user_r.GET("/:id", user.GetOne)
    user_r.GET("/:id/balance", util.AuthMiddleware(), user.GetBalance)
    user_r.POST("/:id/balance", util.AuthMiddleware(), user.AddBalance)
    user_r.POST("/signup", user.Signup)
    user_r.POST("/login", user.Login)
  }

  // Product APIs
  product_r := router.Group("/products")
  {
    product := new(controllers.ProductController)

    product_r.GET("", product.GetAll)
    product_r.GET("/:id", product.GetOne)
		product_r.POST("/:id/order", util.AuthMiddleware(), product.MakeOrder)
    product_r.PUT("/:id", util.AuthMiddleware(), product.EditOne)
    product_r.POST("", util.AuthMiddleware(), product.AddProduct)
    product_r.DELETE("/:id", util.AuthMiddleware(), product.DeleteOne)

  }

  // Invalid routes handler
  router.NoRoute(func(c *gin.Context) {
    c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
  })

  // Run the server
  port := os.Getenv("PORT")
  log.Printf("\n\n PORT: %s \n ENV: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("API_VERSION"))
  router.Run(":" + port)
}

func main() {
  // Load the .env file
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("error: failed to load the env file")
  }

  db.ConnectDatabase()
  util.ConnectCloudinary()

  routes()
}
