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

func (ctrl StoreController) GetAll(c *gin.Context) {
	var stores []models.Store
	db.DB.Preload("Products.User").Find(&stores)

	data := make([]map[string]interface{}, len(stores))
	for i, store := range stores {
		data[i] = store.Serialize()
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}

func (ctrl StoreController) GetOne(c *gin.Context) {
	id := c.Param("id")

	storeId, err := strconv.ParseInt(id, 10, 64)
	if storeId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var store models.Store
	result := db.DB.Preload("Products.User").First(&store, "id = ?", storeId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": store.Serialize()})
}
