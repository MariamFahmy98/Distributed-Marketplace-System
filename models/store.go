package models

import (
	"gorm.io/gorm"
	"time"
)

type Store struct {
	ID        int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	Title     string         `gorm:"title" json:"title"`
	CreatedAt time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Products  []*Product     `gorm:"many2many:product_store;"json: "products"`
}

func (s Store) Serialize() map[string]interface{} {
	products := make([]map[string]interface{}, len(s.Products))
	for i, product := range s.Products {
		products[i] = product.Serialize()
	}

	return map[string]interface{}{
		"id":        s.ID,
		"title":     s.Title,
		"createdAt": s.CreatedAt,
		"updatedAt": s.UpdatedAt,
		"products":  products,
	}
}
