package models

type User struct {
	ID       int64  `db:"id, primarykey, autoincrement" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	Name     string `db:"name" json:"name"`
}

type LoginInput struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//RegisterForm ...
type RegisterInput struct {
	// Name     string `form:"name" json:"name" binding:"required,min=3,max=20"` // rules
	// Email    string `form:"email" json:"email" binding:"required"`
	// Password string `form:"password" json:"password" binding:"required,min=3,max=50"`
	Name     string `form:"name" binding:"required,min=3,max=20"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
