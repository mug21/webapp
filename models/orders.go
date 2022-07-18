package models

type Order struct {
	OrderID   string  `gorm:"primaryKey" json:"order_Id" validate:"required"`
	UserID    string  `json:"user_Id" validate:"required"`
	OrderName string  `json:"orderName" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required,numeric"`
	Amount    float32 `json:"amount" validate:"required,numeric"`
}
