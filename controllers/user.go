package controllers

import (
	"net/http"
	"strconv"

	"github.com/distributed-marketplace-system/db"
	"github.com/distributed-marketplace-system/models"
	"github.com/distributed-marketplace-system/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//UserController ...
type UserController struct{}

func (ctrl UserController) RegisterUser(c *gin.Context) {
	var input models.RegisterInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	var users models.User
	db.DB.Find(&users, "email=?", input.Email)

	if users.Email == input.Email {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Email is Already Registered", "success": false})
		return
	}

	bytePassword := []byte(input.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	// Create new user
	user := models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword)}

	db.DB.Create(&user)

	c.IndentedJSON(http.StatusOK, gin.H{"success": true})
}

func (ctrl UserController) LoginUser(c *gin.Context) {

	var input models.LoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	var user models.User
	if db.DB.Find(&user, "email=?", input.Email).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "You have not registered, please register first", "success": false})
		return
	}

	//Compare the password form and database if match
	bytePassword := []byte(input.Password)
	byteHashedPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Incorrect Password", "success": false})
		return
	}
	token, err := util.CreateToken(input.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.SetCookie("jwt_token", token, 10, "/", "localhost", false, true);
	c.JSON(http.StatusOK, gin.H{"success": true});
}

func (ctrl UserController) GetAll(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)

	c.IndentedJSON(http.StatusOK, gin.H{"data": users, "success": true})
}

func (ctrl UserController) GetOne(c *gin.Context) {
	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter", "success": false})
		return
	}

	var user models.User

	if db.DB.Find(&user, "id=?", getID).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "No User Found", "success": false})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": user, "success": true})
}

func (ctrl UserController) LogoutUser(c *gin.Context) {
	// Delete Authentication token
	c.SetCookie("jwt_token", "", -1, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"data": "Successfully Logged Out", "success": true})
}
