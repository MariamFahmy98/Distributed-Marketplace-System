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
	StoreID   int64          `gorm:"store_id" json:"store_id"`
	CreatedAt time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Products  []*Product     `json:"-"`
	Store     Store          `json:"-"`
}

func (u User) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID,
		"email":      u.Email,
		"name":       u.Name,
		"balance":    u.Balance,
		"image_url":  u.ImageURL,
		"store_id":   u.StoreID,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
}

func (u User) PublicSerialize() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID,
		"email":      u.Email,
		"name":       u.Name,
		"image_url":  u.ImageURL,
		"store_id":   u.StoreID,
	}
}

type LoginInput struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignupInput struct {
	Name     string `form:"name" json:"name" binding:"required`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	ImageURL string `form:"image_url" json:"image_url" json:"image_url"`
}


type EditUserInput struct {
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password""`
	ImageURL string `form:"image_url" json:"image_url"`
}