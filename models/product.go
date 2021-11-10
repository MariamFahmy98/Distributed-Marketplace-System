package models

type Product struct {
	ID       int64   `db:"id, primarykey, autoincrement" json:"id"`
	UserID   int64   `db:"user_id" json:"user_id"` // to eliminate return specific field put "-"
	Title    string  `db:"title" json:"title"`
	Content  string  `db:"content" json:"content"`
	ImageURL string  `db:"image_url" json:"image_url"`
	Price    float64 `db:"price" json:"price"`
}

type AddProductInput struct {
	UserID    int64   `form:"user_id" json:"user_id" binding:"required"`
	Title     string  `form:"title" json:"title" binding:"required"`
	ImagePath string  `form:"image_path" json:"image_path" binding:"required"`
	Content   string  `form:"content" json:"content" binding:"required"`
	Price     float64 `form:"price" json:"price"`
}
