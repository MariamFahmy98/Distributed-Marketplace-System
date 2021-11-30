package controllers

import (
	"net/http"
	"strconv"

	"distributed-marketplace-system/db"
	"distributed-marketplace-system/errors"
	"distributed-marketplace-system/models"
	"distributed-marketplace-system/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct{}

func (ctrl UserController) Signup(c *gin.Context) {
  var input models.SignupInput
  err := c.ShouldBind(&input)
  if err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  var user models.User
  db.DB.Find(&user, "email=?", input.Email)

  if user.Email == input.Email {
    c.AbortWithStatusJSON(422, errors.ErrEmailAlreadyRegistered)
    return
  }

  bytePassword := []byte(input.Password)
  hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
  if err != nil {
    c.AbortWithStatusJSON(422, errors.ErrUnprocessable)
    return
  }

  // Create new user
  user = models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword)}
  db.DB.Create(&user)

  userId := strconv.FormatInt(user.ID, 10)
  token, _ := util.CreateToken(userId)

  c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}

func (ctrl UserController) Login(c *gin.Context) {
  var input models.LoginInput
  err := c.ShouldBind(&input)
  if err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  var user models.User
  result := db.DB.First(&user, "email=?", input.Email)
  if result.Error == gorm.ErrRecordNotFound {
    c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrNotRegistered)
    return
  }

  bytePassword := []byte(input.Password)
  byteHashedPassword := []byte(user.Password)

  err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
  if err != nil {
    c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrIncorrectPassword)
    return
  }

  userId := strconv.FormatInt(user.ID, 10)
  token, err := util.CreateToken(userId)
  if err != nil {
    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ctrl UserController) GetAll(c *gin.Context) {
  var users []models.User
  db.DB.Find(&users)

  c.IndentedJSON(http.StatusOK, gin.H{"data": users})
}

func (ctrl UserController) GetOne(c *gin.Context) {
  id := c.Param("id")

  userId, err := strconv.ParseInt(id, 10, 64)
  if userId == 0 || err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
    return
  }

  var user models.User
  result := db.DB.First(&user, "id=?", userId) 
  if result.Error == gorm.ErrRecordNotFound {
    c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
    return
  }

  c.IndentedJSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl UserController) GetProducts(c *gin.Context) {
  id := c.Param("id")

  userId, err := strconv.ParseInt(id, 10, 64)
  if userId == 0 || err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
    return
  }

  var user models.User
  result := db.DB.First(&user, "id=?", userId)
  if result.Error == gorm.ErrRecordNotFound {
    c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
    return
  }

  c.IndentedJSON(http.StatusOK, gin.H{"data": user.Products})
}

func (ctrl UserController) GetBalance(c *gin.Context) {
  id := c.Param("id")
  userId, err := strconv.ParseInt(id, 10, 64)
  if userId == 0 || err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
    return
  }

  var user models.User
  result := db.DB.First(&user, "id = ?", userId)
  if result.Error == gorm.ErrRecordNotFound {
    c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
    return
  }

  c.IndentedJSON(http.StatusOK, gin.H{"data": user.Balance})
}

func (ctrl UserController) AddBalance(c *gin.Context) {
  id := c.Param("id")
  var user models.User
  var input models.BalanceInput

  userId, err := strconv.ParseInt(id, 10, 64)
  if userId == 0 || err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
    return
  }

  err = c.ShouldBind(&input)
  if err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  result := db.DB.First(&user, "id = ?", userId)
  if result.Error == gorm.ErrRecordNotFound {
    c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
    return
  }

  db.DB.Model(&user).Where("id = ?", userId).Update("balance", input.Balance+user.Balance)

  c.JSON(http.StatusOK, gin.H{"success": true, "balance": user.Balance})
}
