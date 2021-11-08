package controllers

import (
	"net/http"
	"strconv"

	"github.com/distributed-marketplace-system/db"
	"github.com/distributed-marketplace-system/models"
	"github.com/gin-gonic/gin"
)

//ProductController ...
type ProductController struct{}

func (ctrl ProductController) AddProduct(c *gin.Context) {
	var input models.AddProductInput

	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add new product
	product := models.Product{UserID: input.UserID, Title: input.Title, Content: input.Content}

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
	var product models.Product
	db.DB.Where("id=?", getID).Delete(&product)

	if db.DB.Find(&product, "user_id=?", getID).RecordNotFound() {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "the product with id:" + id + " is not deleted"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"msg": "the product with " + id + " is deleted"})
}
