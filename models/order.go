package models

type Order struct {
  ID         int64  `db:"id, primarykey, autoincrement" json:"id"`
  BuyerID    int64  `db:"buyer_id" json:"buyer_id"`
  SellerID   int64  `db:"seller_id" json:"seller_id"`
  ProductID  int64  `db:"product_id" json:"product_id"`

  Product Product `json: "-"`
}
