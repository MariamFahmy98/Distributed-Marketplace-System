package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	UserID    int64          `gorm:"user_id" json:"user_id"`
	StoreID   int64          `gorm:"store_id" json:"store_id"`
	Title     string         `gorm:"title" json:"title"`
	Content   string         `gorm:"content" json:"content"`
	ImageURL  string         `gorm:"image_url" json:"image_url"`
	Price     float64        `gorm:"price" json:"price"`
	Status    bool           `gorm:"status" json:"status"`
	CreatedAt time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Store     Store          `json: "-"`
	User      User           `json: "-"`
}

type AddProductInput struct {
	Title   string  `form:"title" json:"title" binding:"required"`
	Content string  `form:"content" json:"content" binding:"required"`
	Price   float64 `form:"price" json:"price"  binding:"required"`
}
type EditProductInput struct {
	Title   string  `form:"title" json:"title"`
	Content string  `form:"content" json:"content"`
	Price   float64 `form:"price" json:"price"`
}
