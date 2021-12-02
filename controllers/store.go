package controllers

import (
	_ "fmt"
	"net/http"
	"strconv"

	"distributed-marketplace-system/db"
	"distributed-marketplace-system/errors"
	"distributed-marketplace-system/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StoreController struct{}

func (ctrl StoreController) CreateStore(c *gin.Context) {
	var input models.CreateStoreInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Request.Header.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)

	store := models.Store{Title: input.Title, UserID: userId}
	db.DB.Create(&store)
	c.IndentedJSON(http.StatusOK, gin.H{"data": store})

}

func (ctrl StoreController) GetAll(c *gin.Context) {
	var stores []models.Store
	db.DB.Find(&stores)

	c.IndentedJSON(http.StatusOK, gin.H{"data": stores})

}

func (ctrl StoreController) GetOne(c *gin.Context) {
	id := c.Param("id")

	storeId, err := strconv.ParseInt(id, 10, 64)
	if storeId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var store models.Store

	result := db.DB.First(&store, "id=?", storeId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	var products []models.Product

	result = db.DB.Find(&products, "store_id=?", storeId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": products})
}
