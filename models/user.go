package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	Email     string         `gorm:"email, uniqueIndex" json:"email"`
	Password  string         `gorm:"password" json:"-"`
	Name      string         `gorm:"name" json:"name"`
	Balance   float64        `gorm:"balance" json:"-"`
	ImageURL  string         `gorm:"image_url" json:"image_url"`
	CreatedAt time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Products  []*Product     `json:"-"`
}

type LoginInput struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignupInput struct {
	Name     string `form:"name" binding:"required,min=3,max=20"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
