package models

import "time"

type Deposit struct {
	ID            int64 `db:"id, primarykey, autoincrement" json:"id"`
	UserID        int64 `db:"user_id" json:"user_id"`
	BalanceBefore int64 `db:"balance_before" json:"balance_before"`
	Amount        int64 `db:"amount" json:"amount"`
	BalanceAfter  int64 `db:"balance_after" json:"balance_after"`
	CreatedAt     time.Time
	UpdatesAt     time.Time
	User          User `json: "-"`
}

type DepositInput struct {
	Amount int64 `form:"amount" json:"amount" binding:"required"`
}
