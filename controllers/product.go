package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/distributed-marketplace-system/db"
	"github.com/distributed-marketplace-system/models"
	"github.com/distributed-marketplace-system/util"
	"github.com/gin-gonic/gin"
)

func UploadProductImage(img_path string) (url string, err error) {

	if _, err := os.Stat(img_path); os.IsNotExist(err) {
		return "", err
	}

	var ctx = context.Background()
	uploadResult, err := util.CLD.Upload.Upload(ctx, img_path, uploader.UploadParams{Folder: "asu"})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
		return "", err
	}

	//log.Println(uploadResult.SecureURL)
	return uploadResult.SecureURL, nil
}

//ProductController ...
type ProductController struct{}

func (ctrl ProductController) AddProduct(c *gin.Context) {

	// Get the email of the user to compare with email in the header
	var input models.AddProductInput
	err := c.ShouldBind(&input)

	var user models.User
	if db.DB.Find(&user, "id=?", input.UserID).RecordNotFound() {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "there is no user with id " + (string)(input.UserID)})
		return
	}

	email := c.Request.Header.Get("email")
	if email != user.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to be here"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// upload the image to cloudinary cloud
	img, err := UploadProductImage(input.ImagePath)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add new product
	product := models.Product{UserID: input.UserID, Title: input.Title, Content: input.Content, Price: input.Price, ImageURL: img}

	db.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (ctrl ProductController) GetAll(c *gin.Context) {
	var products []models.Product
	db.DB.Find(&products)

	c.IndentedJSON(http.StatusOK, gin.H{"data": products})
}

func (ctrl ProductController) GetOne(c *gin.Context) {

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var product models.Product

	if db.DB.Find(&product, "id=?", getID).RecordNotFound() {
		c.IndentedJSON(http.StatusOK, gin.H{"msg": "there is no product with id " + id})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": product})
}

func (ctrl ProductController) DeleteOne(c *gin.Context) {

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var user models.User
	var product models.Product
	if db.DB.Find(&product, "id=?", getID).RecordNotFound() {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "there is no product with id = " + strconv.Itoa((int)(getID))})
		return
	}

	if db.DB.Find(&user, "id=?", product.UserID).RecordNotFound() {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "there is no user with id " + strconv.Itoa((int)(product.UserID))})
		return
	}
	email := c.Request.Header.Get("email")
	fmt.Println(email)
	fmt.Println(user.Email)
	if email != user.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to be here"})
		c.Abort()
		return
	}

	// you are authorized to delete
	db.DB.Where("id=?", getID).Delete(&product)

	// check if the deletion is performed or not
	if db.DB.Find(&product, "id=?", getID).RecordNotFound() {
		c.IndentedJSON(http.StatusOK, gin.H{"msg": "the product with id " + strconv.Itoa((int)(getID)) + " is deleted"})
		return
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "there is no product with id = " + strconv.Itoa((int)(getID))})
		return
	}

}
