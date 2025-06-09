package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `gorm:"uniqueIndex;not null" json:"user_id"`
	User      User           `json:"user"`
	Items     []CartItem     `json:"items"`
	Total     float64        `gorm:"default:0" json:"total"`
}

type CartItem struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CartID    uint           `gorm:"not null" json:"cart_id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Product   Product        `json:"product"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	Price     float64        `gorm:"not null" json:"price"`
	Subtotal  float64        `gorm:"not null" json:"subtotal"`
}
