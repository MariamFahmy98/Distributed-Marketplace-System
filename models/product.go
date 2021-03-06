package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	UserID    int64          `gorm:"user_id" json:"user_id"`
	Title     string         `gorm:"title" json:"title"`
	Content   string         `gorm:"content" json:"content"`
	ImageURL  string         `gorm:"image_url" json:"image_url"`
	Price     float64        `gorm:"price" json:"price"`
	Status    bool           `gorm:"status" json:"status"`
	CreatedAt time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Stores    []*Store       `gorm:"many2many:product_store;"json: "-"`
	User      User           `json: "-"`
}

func (p Product) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":         p.ID,
		"user":    		p.User.PublicSerialize(),
		"title":      p.Title,
		"content":    p.Content,
		"image_url":  p.ImageURL,
		"price":      p.Price,
		"status":     p.Status,
		"created_at": p.CreatedAt,
		"updated_at": p.UpdatedAt,
	}
}

type AddProductInput struct {
	Title    string  `form:"title" json:"title" binding:"required"`
	Content  string  `form:"content" json:"content" binding:"required"`
	Price    float64 `form:"price" json:"price"  binding:"required"`
	ImageURL string  `form:"image_url" json:"image_url"`
}

type EditProductInput struct {
	Title   string  `form:"title" json:"title"`
	Content string  `form:"content" json:"content"`
	Price   float64 `form:"price" json:"price"`
	ImageURL string  `form:"image_url" json:"image_url"`
}
