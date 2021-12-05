package models

import (
	"gorm.io/gorm"
	"time"
)

type Deposit struct {
	ID            int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	UserID        int64          `gorm:"user_id" json:"user_id"`
	BalanceBefore int64          `gorm:"balance_before" json:"balance_before"`
	Amount        int64          `gorm:"amount" json:"amount"`
	BalanceAfter  int64          `gorm:"balance_after" json:"balance_after"`
	CreatedAt     time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	User          User           `json: "-"`
}

type DepositInput struct {
	Amount int64 `form:"amount" json:"amount" binding:"required"`
}