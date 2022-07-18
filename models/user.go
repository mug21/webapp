package models

type User struct {
	UserID   string  `gorm:"primaryKey" json:"user_Id" validate:"required"`
	Username string  `json:"username" validate:"required"`
	EmailID  string  `json:"email_Id" validate:"required,email"`
	Address  string  `json:"address" validate:"required"`
	Orders   []Order `gorm:"constraint:OnDelete:CASCADE" json:"orders,omitempty"`
}
