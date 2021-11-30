package models

// import (
// 	_ "time"
// )

type User struct {
  ID       int64  `db:"id, primarykey, autoincrement" json:"id"`
  Email    string `db:"email, uniqueIndex" json:"email"`
  Password string `db:"password" json:"-"`
  Name     string `db:"name" json:"name"`
  Balance  int64  `db:"balance" json:"-"`
  ImageURL string `db:"image_url" json:"image_url"`

  Products []Product `json:"-"`
}

type BalanceInput struct {
  Balance int64 `form:"amount" json:"amount" binding:"required"`
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
