package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Price       float64        `gorm:"not null" json:"price"`
	Stock       int            `gorm:"not null" json:"stock"`
	CategoryID  uint           `gorm:"not null" json:"category_id"`
	Category    Category       `json:"category"`
	Images      []Image        `json:"images"`
	SKU         string         `gorm:"uniqueIndex" json:"sku"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
}

type Category struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Products    []Product      `json:"products,omitempty"`
}

type Image struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	URL       string         `gorm:"not null" json:"url"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	IsPrimary bool           `gorm:"default:false" json:"is_primary"`
}
