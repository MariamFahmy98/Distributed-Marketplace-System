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
	UserRoute := router.Group("/users")
	{
		user := new(controllers.UserController)

		UserRoute.GET("", user.GetAll)
		UserRoute.GET("/:id", user.GetOne)
		UserRoute.GET("/:id/balance", util.AuthMiddleware(), user.GetBalance)
		UserRoute.POST("/:id/balance", util.AuthMiddleware(), user.AddBalance)
		UserRoute.GET("/:id/sold_products", util.AuthMiddleware(), user.GetSoldProducts)
		UserRoute.GET("/:id/purchased_products", util.AuthMiddleware(), user.GetPurchasedProducts)
		UserRoute.GET("/:id/report/orders", util.AuthMiddleware(), user.GetReportOnOrders)
		UserRoute.GET("/:id/report/deposits", util.AuthMiddleware(), user.GetReportOnDeposits)
		UserRoute.POST("/signup", user.Signup)
		UserRoute.POST("/login", user.Login)
	}

	// Product APIs
	ProductRoute := router.Group("/products")
	{
		product := new(controllers.ProductController)

		ProductRoute.GET("", product.GetAll)
		ProductRoute.GET("/search", product.SearchAll)
		ProductRoute.GET("/:id", product.GetOne)
		ProductRoute.POST("/:id/store", util.AuthMiddleware(), product.AddtoStore)
		ProductRoute.POST("/:id/order", util.AuthMiddleware(), product.MakeOrder)
		ProductRoute.PUT("/:id", util.AuthMiddleware(), product.EditOne)
		ProductRoute.POST("", util.AuthMiddleware(), product.AddProduct)
		ProductRoute.DELETE("/:id", util.AuthMiddleware(), product.DeleteOne)

	}

	// Store APIs
	StoreRoute := router.Group("/stores")
	{
		store := new(controllers.StoreController)
		StoreRoute.GET("", store.GetAll)
		StoreRoute.GET("/:id", store.GetOne)
		StoreRoute.POST("", util.AuthMiddleware(), store.CreateStore)
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
