package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/distributed-marketplace-system/controllers"
	"github.com/distributed-marketplace-system/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	router := gin.Default()

	// Load the static files
	router.LoadHTMLGlob("./public/html/*")
	router.Static("/public", "./public")

	db.ConnectDatabase()

	// User APIs
	user := new(controllers.UserController)

	router.GET("/user/get_all", user.GetAll)
	router.GET("/user/get_one/:id", user.GetOne)
	router.POST("/user/register", user.RegisterUser)
	router.POST("/user/login", user.LoginUser)

	// Product APIs
	product := new(controllers.ProductController)

	router.POST("/product/add", product.AddProduct)
	router.GET("/product/get_all", product.GetAll)
	router.GET("/product/get_one/:id", product.GetOne)
	router.DELETE("/product/delete_one/:id", product.DeleteOne)

	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	// Invalid routes handler
	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	// Run the server
	port := os.Getenv("PORT")
	log.Printf("\n\n PORT: %s \n ENV: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("API_VERSION"))
	router.Run(":" + port)
}
