package controllers

import (
	"net/http"
	"strconv"

	"github.com/distributed-marketplace-system/db"
	"github.com/distributed-marketplace-system/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//UserController ...
type UserController struct{}

func (ctrl UserController) RegisterUser(c *gin.Context) {
	var input models.RegisterInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var users models.User
	db.DB.Find(&users, "email=?", input.Email)

	if users.Email == input.Email {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "This email is already registered"})
		return
	}

	bytePassword := []byte(input.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new user
	user := models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword)}

	db.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl UserController) LoginUser(c *gin.Context) {

	var input models.LoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if db.DB.Find(&user, "email=?", input.Email).RecordNotFound() {
		c.IndentedJSON(http.StatusOK, gin.H{"msg": "Please regiter first"})
		return
	}

	//Compare the password form and database if match
	bytePassword := []byte(input.Password)
	byteHashedPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "This password is incorrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl UserController) GetAll(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)

	c.IndentedJSON(http.StatusOK, gin.H{"data": users})
}

func (ctrl UserController) GetOne(c *gin.Context) {
	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var user models.User

	if db.DB.Find(&user, "id=?", getID).RecordNotFound() {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "there is no user with id " + id})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl UserController) LogoutUser(c *gin.Context) {
	// Delete Authentication token

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
