package main

import (
	"log"
	"net/http"
	"os"

	"distributed-marketplace-system/controllers"
	"distributed-marketplace-system/db"
	"distributed-marketplace-system/errors"
	"distributed-marketplace-system/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func routes() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	router.Use(cors.New(config))
	// User APIs
	UserRoute := router.Group("/users")
	{
		user := new(controllers.UserController)

		UserRoute.GET("", user.GetAll)
		UserRoute.GET("/:id", user.GetOne)
		UserRoute.GET("/balance", util.AuthMiddleware(), user.GetBalance)
		UserRoute.GET("/:id/products", user.GetProducts)
		UserRoute.GET("/sold_products", util.AuthMiddleware(), user.GetSoldProducts)
		UserRoute.GET("/purchased_products", util.AuthMiddleware(), user.GetPurchasedProducts)
		UserRoute.GET("/report/orders", util.AuthMiddleware(), user.GetReportOnOrders)
		UserRoute.GET("/report/transactions", util.AuthMiddleware(), user.GetReportOnTransactions)
		UserRoute.POST("/signup", user.Signup)
		UserRoute.POST("/login", user.Login)
		UserRoute.POST("/balance", util.AuthMiddleware(), user.AddBalance)
		UserRoute.PUT("", util.AuthMiddleware(), user.EditOne)
	}

	// Product APIs
	ProductRoute := router.Group("/products")
	{
		product := new(controllers.ProductController)

		ProductRoute.GET("", product.GetAll)
		ProductRoute.GET("/:id", product.GetOne)
		ProductRoute.GET("/search", product.SearchAll)
		ProductRoute.POST("/:id/store", util.AuthMiddleware(), product.AddtoStore)
		ProductRoute.POST("/:id/order", util.AuthMiddleware(), product.MakeOrder)
		ProductRoute.POST("", util.AuthMiddleware(), product.AddProduct)
		ProductRoute.PUT("/:id", util.AuthMiddleware(), product.EditOne)
		ProductRoute.DELETE("/:id", util.AuthMiddleware(), product.DeleteOne)

	}

	// Store APIs
	StoreRoute := router.Group("/stores")
	{
		store := new(controllers.StoreController)
		StoreRoute.GET("", store.GetAll)
		StoreRoute.GET("/:id", store.GetOne)
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
	db.ConnectDatabase()
	util.ConnectCloudinary()

	routes()
}
