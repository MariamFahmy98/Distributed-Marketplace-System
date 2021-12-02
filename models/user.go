package models

import "time"

type User struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email, uniqueIndex" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	Balance   int64  `db:"balance" json:"balance"`
	ImageURL  string `db:"image_url" json:"image_url"`
	CreatedAt time.Time
	UpdatesAt time.Time
	Products  []Product `json:"-"`
}

type LoginInput struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignupInput struct {
	Name     string `form:"name" binding:"required,min=3,max=20"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
