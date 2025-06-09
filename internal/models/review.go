package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	UserID     uint           `gorm:"not null" json:"user_id"`
	User       User           `json:"user"`
	ProductID  uint           `gorm:"not null" json:"product_id"`
	Product    Product        `json:"product"`
	Rating     int            `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Title      string         `gorm:"size:100" json:"title"`
	Comment    string         `gorm:"type:text" json:"comment"`
	IsVerified bool           `gorm:"default:false" json:"is_verified"` // Whether the user purchased the product
}
