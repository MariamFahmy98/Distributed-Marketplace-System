package main

import (
	"log"
	"net/http"
	"os"

	"github.com/distributed-marketplace-system/controllers"
	"github.com/distributed-marketplace-system/db"

	"github.com/distributed-marketplace-system/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func routes() {
	router := gin.Default()

	// Load the static files
	router.LoadHTMLGlob("./public/html/*")
	router.Static("/public", "./public")

	user := new(controllers.UserController)
	router.POST("/register", user.RegisterUser)
	router.POST("/login", user.LoginUser)

	// User APIs
	user_r := router.Group("/users")
	{

		user_r.GET("/", user.GetAll)
		user_r.GET("/:id", user.GetOne)

	}

	// Product APIs
	product_r := router.Group("/products")
	{
		product := new(controllers.ProductController)

		product_r.POST("/", util.AuthMiddleware(), product.AddProduct)
		product_r.DELETE("/:id", util.AuthMiddleware(), product.DeleteOne)
		product_r.GET("/", product.GetAll)

		product_r.GET("/:id", product.GetOne)
	}

	// Invalid routes handler
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "404 Not Found", "success": false})
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
