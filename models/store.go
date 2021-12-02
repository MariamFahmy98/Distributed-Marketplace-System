package models

import "time"

type Store struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Title     string `db:"title" json:"title"`
	UserID    int64  `db:"user_id" json:"user_id"`
	CreatedAt time.Time
	UpdatesAt time.Time
	User      User `json: "-"`
}

type CreateStoreInput struct {
	Title string `form:"title" json:"title" binding:"required"`
}
