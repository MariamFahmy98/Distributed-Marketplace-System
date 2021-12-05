package models

import (
	"gorm.io/gorm"
	"time"
)

type Store struct {
	ID        int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	Title     string         `gorm:"title" json:"title"`
	UserID    int64          `gorm:"user_id" json:"user_id"`
	CreatedAt time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	User      User           `json: "-"`
	Products  []Product      `json:"-"`
}

type CreateStoreInput struct {
	Title string `form:"title" json:"title" binding:"required"`
}
