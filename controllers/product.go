package controllers

import (
	// "context"
	_ "fmt"
	// "io/ioutil"
	"sync"
	"net/http"
	// "os"
	"strconv"
	"strings"

	"distributed-marketplace-system/db"
	"distributed-marketplace-system/errors"
	"distributed-marketplace-system/models"
	"distributed-marketplace-system/util"

	// "github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var productLock = &sync.Mutex{}

//ProductController ...
type ProductController struct{}


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
	result := db.DB.Preload("Store").First(&user, "users.id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	img_url, _ := util.UploadImage(input.ImageURL)

	// Store the new product in the database
	productLock.Lock()
	db.DB.Model(&user.Store).Association("Products").Append(&models.Product{
		UserID:   userId,
		Title:    input.Title,
		Content:  input.Content,
		Price:    input.Price,
		ImageURL: img_url,
		Status:   true,
	})

	var productId int64
	row := db.DB.Table("products").Select("max(id)").Row()
	row.Scan(&productId)
	productLock.Unlock()

	var product models.Product
	db.DB.Preload("User").Find(&product, "products.id = ?", productId)

	c.IndentedJSON(http.StatusOK, gin.H{"data": product.Serialize()})
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

	img_url, _ := util.UploadImage(input.ImageURL)

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

	db.DB.Model(&models.Product{}).Where("id = ?", productId).Updates(productMap)
	db.DB.Preload("User").First(&product, productId)
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

		tx.Create(&models.Transaction{
			UserID:        product.User.ID,
			Amount:        product.Price,
			BalanceBefore: product.User.Balance,
			Type: 	       "Item Sold",
		})

		tx.Create(&models.Transaction{
			UserID:        user.ID,
			Amount:        - product.Price,
			BalanceBefore: user.Balance,
			Type: 	       "Item Bought",
		})

		tx.Model(&models.User{}).Where("id = ?", userId).Update("balance", gorm.Expr("balance - ?", product.Price))
		tx.Model(&models.User{}).Where("id = ?", product.UserID).Update("balance", gorm.Expr("balance + ?", product.Price))
		tx.Model(&models.Product{}).Where("id = ?", product.ID).Updates(map[string]interface{}{"user_id": userId, "status": false})
		tx.Model(&models.Product{}).Where("id = ?", product.ID).Association("Stores").Delete(product.Stores)

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
	result = db.DB.First(&product, "id = ?", productId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	if product.UserID == userId && product.Status == false {
		db.DB.Model(&models.Product{}).Where("id = ?", product.ID).Update("status", true)
	}

	if product.Status == false {
		c.AbortWithStatusJSON(422, errors.ErrNotForSales)
		return
	}

	db.DB.Model(&models.Product{}).Where("id = ?", product.ID).Association("Stores").Append(&user.Store)
	c.IndentedJSON(http.StatusOK, gin.H{"sccess": true})
}
