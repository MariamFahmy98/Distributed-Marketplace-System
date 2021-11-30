package controllers

import (
	"context"
	_ "fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"distributed-marketplace-system/db"
	"distributed-marketplace-system/errors"
	"distributed-marketplace-system/models"
	"distributed-marketplace-system/util"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	var input models.AddProductInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Request.Header.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)

	img_url := ctrl.extractImage(c)

	// Store the new product in the database
	product := models.Product{UserID: userId, Title: input.Title, Content: input.Content, Price: input.Price, ImageURL: img_url}
	db.DB.Create(&product)

	c.IndentedJSON(http.StatusOK, gin.H{"data": product})
}

func (ctrl ProductController) extractImage(c *gin.Context) string {
	// Extract image from the form
	r := c.Request
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("image")
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return ""
	}
	defer file.Close()

	// Store image in the server side
	tempFile, err := ioutil.TempFile("res", "upload-*.png")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// upload the image to cloudinary cloud
	img_url, err := UploadProductImage(tempFile.Name())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return ""
	}

	// remove the image from the server side after uploading it
	e := os.Remove(tempFile.Name())
	if e != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
	}

	return img_url
}

func (ctrl ProductController) GetAll(c *gin.Context) {
	var products []models.Product
	db.DB.Find(&products)

	c.IndentedJSON(http.StatusOK, gin.H{"data": products})
}

func (ctrl ProductController) GetOne(c *gin.Context) {

	id := c.Param("id")

	productId, err := strconv.ParseInt(id, 10, 64)
	if productId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var product models.Product

	result := db.DB.First(&product, "id=?", productId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": product})
}

func (ctrl ProductController) DeleteOne(c *gin.Context) {

	id := c.Param("id")

	productId, err := strconv.ParseInt(id, 10, 64)
	if productId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var product models.Product
	result := db.DB.First(&product, "id=?", productId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	id = c.Request.Header.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)
	if product.UserID != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrUnauthorized)
		c.Abort()
		return
	}

	// you are authorized to delete
	result = db.DB.Where("id=?", productId).Delete(&product)

	if result.Error != nil {
		c.AbortWithStatusJSON(422, gin.H{"success": false})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"success": true})
	}

}

func (ctrl ProductController) EditOne(c *gin.Context) {

	id := c.Param("id")
	productId, err := strconv.ParseInt(id, 10, 64)
	if productId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var input models.EditProductInput
	err = c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product

	db.DB.First(&product, productId)

	img_url := ctrl.extractImage(c)

	if len(input.Title) != 0 {
		product.Title = input.Title
	}

	if len(input.Content) != 0 {
		product.Content = input.Content
	}

	if input.Price != 0 {
		product.Price = input.Price
	}

	if len(img_url) != 0 {
		product.ImageURL = img_url
	}

	db.DB.Save(&product)

	c.IndentedJSON(http.StatusOK, gin.H{"data": product})
}

func (ctrl ProductController) MakeOrder(c *gin.Context) {
	id := c.Param("id")
	productId, err := strconv.ParseInt(id, 10, 64)
	if productId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var product models.Product
	result := db.DB.First(&product, "id=?", productId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	id = c.Request.Header.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)

	var user models.User
	result = db.DB.First(&user, "id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	if user.Balance < int64(product.Price) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Your balance isn't sufficient!"})
		return
	}

	db.DB.Model(&user).Where("id = ?", userId).Update("balance", user.Balance-int64(product.Price))

	c.JSON(http.StatusOK, gin.H{"success": true, "balance": user.Balance})

}
