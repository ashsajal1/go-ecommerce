package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `json:"user"`
	Type      string         `gorm:"type:varchar(20);not null" json:"type"` // shipping or billing
	Street    string         `gorm:"not null" json:"street"`
	City      string         `gorm:"not null" json:"city"`
	State     string         `gorm:"not null" json:"state"`
	Country   string         `gorm:"not null" json:"country"`
	ZipCode   string         `gorm:"not null" json:"zip_code"`
	IsDefault bool           `gorm:"default:false" json:"is_default"`
}
