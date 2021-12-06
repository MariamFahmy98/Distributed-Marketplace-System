package models

import (
	"gorm.io/gorm"
	"time"
)

type Deposit struct {
	ID            int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	UserID        int64          `gorm:"user_id" json:"user_id"`
	BalanceBefore float64        `gorm:"balance_before" json:"balance_before"`
	Amount        float64        `gorm:"amount" json:"amount"`
	CreatedAt     time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	User          User           `json: "-"`
}

func (d Deposit) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":             d.ID,
		"balance_before": d.BalanceBefore,
		"amount":         d.Amount,
		"created_at":     d.CreatedAt,
		"updated_at":     d.UpdatedAt,
	}
}

type DepositInput struct {
	Amount float64 `form:"amount" json:"amount" binding:"required"`
}
