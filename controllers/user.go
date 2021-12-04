package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"distributed-marketplace-system/db"
	"distributed-marketplace-system/errors"
	"distributed-marketplace-system/models"
	"distributed-marketplace-system/sqlc/sqlc"
	"distributed-marketplace-system/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type UserController struct{}

func (ctrl UserController) Signup(c *gin.Context) {
  var input models.SignupInput
  err := c.ShouldBind(&input)
  if err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  
	ctx := context.Background()
  // Change later to use env variables 😅
  db, err := sql.Open("postgres" , "host=localhost user=postgres password=1700455 dbname=ds_db sslmode=disable")
  if err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}
	queries := sqlc.New(db)

  existingUser , err := queries.GetUserByEmail(ctx, input.Email)
  if err == nil {
    // User doesnt exist
  }
  if existingUser.ID > 0  {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":errors.ErrEmailAlreadyRegistered})
    return
  }
  
  bytePassword := []byte(input.Password)
  hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
  if err != nil {
    c.AbortWithStatusJSON(422, errors.ErrUnprocessable)
    return
  }

  user, err := queries.CreateUser(ctx, sqlc.CreateUserParams{
    Name: input.Name,
    Email: input.Email,
    Password: string(hashedPassword),
  })
  if err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
    return
	}
  userId := strconv.FormatInt(int64(user.ID), 10)
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
