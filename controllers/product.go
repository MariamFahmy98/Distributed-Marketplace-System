package controllers

import (
	"context"
	_ "fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	uploadResult, err := util.CLD.Upload.Upload(
		ctx,
		img_path,
		uploader.UploadParams{Folder: os.ExpandEnv("CLOUDAINARY_STORAGE_FOLDER")},
	)
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

	var user models.User
	result := db.DB.Joins("Store").First(&user, "users.id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	img_url := ctrl.extractImage(c)

	// Store the new product in the database
	product := models.Product{
		UserID:   userId,
		Title:    input.Title,
		Content:  input.Content,
		Price:    input.Price,
		ImageURL: img_url,
		Status:   true,
	}
	db.DB.Create(&product)
	db.DB.Joins("User").Find(&product, "products.id = ?", product.ID)
	db.DB.Model(&user.Store).Association("Products").Append(&product)

	c.IndentedJSON(http.StatusOK, gin.H{"data": product.Serialize()})
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
		// c.AbortWithStatusJSON(422, gin.H{"error": err.Error()})
		return ""
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		// c.AbortWithStatusJSON(422, gin.H{"error": err.Error()})
		return ""
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// upload the image to cloudinary cloud
	img_url, err := UploadProductImage(tempFile.Name())
	if err != nil {
		// c.AbortWithStatusJSON(422, gin.H{"error": err.Error()})
		return ""
	}

	// remove the image from the server side after uploading it
	e := os.Remove(tempFile.Name())
	if e != nil {
		// c.AbortWithStatusJSON(422, gin.H{"error": e.Error()})
		return ""
	}

	return img_url
}

func (ctrl ProductController) GetAll(c *gin.Context) {
	var products []models.Product
	db.DB.Preload("User").Find(&products)

	data := make([]map[string]interface{}, len(products))
	for i, product := range products {
		data[i] = product.Serialize()
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}

func (ctrl ProductController) GetOne(c *gin.Context) {

	id := c.Param("id")

	productId, err := strconv.ParseInt(id, 10, 64)
	if productId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var product models.Product

	result := db.DB.Preload("User").First(&product, "id=?", productId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": product.Serialize()})
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
	result := db.DB.Preload("User").First(&product, productId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	id = c.Request.Header.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)
	if product.UserID != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrUnauthorized)
		return
	}

	img_url := ctrl.extractImage(c)

	productMap := make(map[string]interface{})
	if len(input.Title) != 0 {
		productMap["title"] = input.Title
	}

	if len(input.Content) != 0 {
		productMap["content"] = input.Content
	}

	if input.Price != 0 {
		productMap["price"] = input.Price
	}

	if len(img_url) != 0 {
		productMap["image_url"] = img_url
	}

	db.DB.Model(&product).Updates(productMap)
	c.IndentedJSON(http.StatusOK, gin.H{"data": product.Serialize()})
}

func (ctrl ProductController) MakeOrder(c *gin.Context) {
	id := c.Param("id")
	productId, err := strconv.ParseInt(id, 10, 64)
	if productId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var product models.Product
	result := db.DB.First(&product, "id = ?", productId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	if product.Status == false {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrNotForSales)
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

	if product.UserID == userId {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrCannotBuyYourProduct)
		return
	}

	if user.Balance < product.Price {
		c.AbortWithStatusJSON(422, errors.ErrBalanceNotEnough)
		return
	}

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&models.Order{
			BuyerID:   userId,
			SellerID:  product.UserID,
			ProductID: productId,
			Price:     product.Price,
		})

		db.DB.Create(&models.Transaction{
			UserID:        product.User.ID,
			Amount:        product.Price,
			BalanceBefore: product.User.Balance,
			Type: 	       "Item Sold",
		})

		db.DB.Create(&models.Transaction{
			UserID:        user.ID,
			Amount:        - product.Price,
			BalanceBefore: user.Balance,
			Type: 	       "Item Bought",
		})

		tx.Model(&user).Where("id = ?", userId).Update("balance", gorm.Expr("balance - ?", product.Price))
		tx.Model(&user).Where("id = ?", product.UserID).Update("balance", gorm.Expr("balance + ?", product.Price))
		tx.Model(&product).Updates(map[string]interface{}{"user_id": userId, "status": false})
		tx.Model(&product).Association("Stores").Delete(product.Stores)

		return nil
	})

	success := true
	if err != nil {
		success = false
	}

	c.JSON(http.StatusOK, gin.H{"success": success})
}

func (ctrl ProductController) SearchAll(c *gin.Context) {
	keyword := c.Query("q")

	if len(strings.TrimSpace(keyword)) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var products []models.Product
	db.DB.Where("title ILIKE ? OR content ILIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&products)

	data := make([]map[string]interface{}, len(products))
	for i, product := range products {
		data[i] = product.Serialize()
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}

func (ctrl ProductController) AddtoStore(c *gin.Context) {
	id := c.Request.Header.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)

	var user models.User
	result := db.DB.Joins("Store").First(&user, "id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	id = c.Param("id")
	productId, err := strconv.ParseInt(id, 10, 64)
	if productId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var product models.Product
	result = db.DB.First(&product, "id=?", productId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	if product.UserID == userId && product.Status == false {
		db.DB.Model(&product).Update("status", true)
	}

	if product.Status == false {
		c.AbortWithStatusJSON(422, errors.ErrNotForSales)
		return
	}

	db.DB.Model(&product).Association("Stores").Append(&user.Store)
	c.IndentedJSON(http.StatusOK, gin.H{"sccess": true})
}
