package controllers

import (
  "net/http"
  "github.com/distributed-marketplace-system/db"
  "github.com/distributed-marketplace-system/models"
  "github.com/gin-gonic/gin"
  _ "fmt"
)

func GetUsers(c *gin.Context) {
  var users []models.User
  db.DB.Find(&users)
  c.IndentedJSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
  // Validate input
  var input models.UserInput
  err := c.ShouldBindJSON(&input)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create new user
  user := models.User{ID: input.ID, Name: input.Name, Email: input.Email, Password: input.Password}
  db.DB.Create(&user)

  c.JSON(http.StatusOK, gin.H{"data": user})
}
