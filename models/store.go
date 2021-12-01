package models

type Store struct {
  ID         int64  `db:"id, primarykey, autoincrement" json:"id"`
	Title      string `db:"title" json:"title"`
  UserID     int64  `db:"user_id" json:"user_id"`

	User  User  `json: "-"`
}


type CreateStoreInput struct {
  Title   string  `form:"title" json:"title" binding:"required"`
}