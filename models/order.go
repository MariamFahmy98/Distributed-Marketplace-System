package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	BuyerID   int64          `gorm:"buyer_id" json:"buyer_id"`
	SellerID  int64          `gorm:"seller_id" json:"seller_id"`
	ProductID int64          `gorm:"product_id" json:"product_id"`
	CreatedAt time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Product   Product        `json: "-"`
	Seller    User           `json: "-"`
	Buyer     User           `json: "-"`
}
