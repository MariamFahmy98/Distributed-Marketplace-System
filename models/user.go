package models

//import "github.com/jinzhu/gorm"

type User struct {
	ID       uint64 `json:"id"` //gorm:"primary_key"
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserInput struct {
	ID       uint64 `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
