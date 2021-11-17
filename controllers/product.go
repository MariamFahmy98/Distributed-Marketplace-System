package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/distributed-marketplace-system/db"
	"github.com/distributed-marketplace-system/models"
	"github.com/distributed-marketplace-system/util"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

//ProductController ...
type ProductController struct{}

func UploadProductImage(img_path string) (url string, err error) {

	if _, err := os.Stat(img_path); os.IsNotExist(err) {
		return "", err
	}

	var ctx = context.Background()
	uploadResult, err := util.CLD.Upload.Upload(ctx, img_path, uploader.UploadParams{Folder: os.ExpandEnv("CLOUDAINARY_STORAGE_FOLDER")})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}

func (ctrl ProductController) AddProduct(c *gin.Context) {

	// Get the email of the user to compare with email in the header
	var input models.AddProductInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	var user models.User
	if db.DB.Find(&user, "id=?", input.UserID).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Something Went Wrong", "success": false})
		return
	}

	email := c.Request.Header.Get("email")
	if email != user.Email {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action", "success": false})
		c.Abort()
		return
	}

	// Extract image from the form
	r := c.Request
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	defer file.Close()

	// Store image in the server side
	tempFile, err := ioutil.TempFile("res", "upload-*.png")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "success": false})
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "success": false})
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// upload the image to cloudinary cloud
	img_url, err := UploadProductImage(tempFile.Name())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "success": false})
		return
	}

	// remove the image from the server side after uploading it
	e := os.Remove(tempFile.Name())
	if e != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Error(), "success": false})
	}

	// Store the new product in the database
	product := models.Product{UserID: input.UserID, Title: input.Title, Content: input.Content, Price: input.Price, ImageURL: img_url}
	db.DB.Create(&product)

	c.IndentedJSON(http.StatusOK, gin.H{"data": product, "success": true})
}

func (ctrl ProductController) GetAll(c *gin.Context) {
	var products []models.Product
	db.DB.Find(&products)

	c.IndentedJSON(http.StatusOK, gin.H{"data": products, "success": true})
}

func (ctrl ProductController) GetOne(c *gin.Context) {

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter", "success": false})
		return
	}

	var product models.Product

	if db.DB.Find(&product, "id=?", getID).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something Went Wrong", "success": false})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": product, "success": true})
}

func (ctrl ProductController) DeleteOne(c *gin.Context) {

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter", "success": false})
		return
	}

	var user models.User
	var product models.Product
	if db.DB.Find(&product, "id=?", getID).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something Went Wrong", "success": false})
		return
	}

	if db.DB.Find(&user, "id=?", product.UserID).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something Went Wrong", "success": false})
		return
	}
	email := c.Request.Header.Get("email")
	fmt.Println(email)
	fmt.Println(user.Email)
	if email != user.Email {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action", "success": false})
		c.Abort()
		return
	}

	// you are authorized to delete
	db.DB.Where("id=?", getID).Delete(&product)

	// check if the deletion is performed or not
	if db.DB.Find(&product, "id=?", getID).RecordNotFound() {
		c.IndentedJSON(http.StatusOK, gin.H{"data": "The action is performed", "success": true})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "No Product Found", "success": false})
		return
	}

}
