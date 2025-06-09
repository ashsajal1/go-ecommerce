package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type Order struct {
	ID                uint           `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	UserID            uint           `gorm:"not null" json:"user_id"`
	User              User           `json:"user"`
	Status            OrderStatus    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	TotalAmount       float64        `gorm:"not null" json:"total_amount"`
	ShippingCost      float64        `json:"shipping_cost"`
	TaxAmount         float64        `json:"tax_amount"`
	Discount          float64        `json:"discount"`
	Items             []OrderItem    `json:"items"`
	ShippingAddressID uint           `gorm:"not null" json:"shipping_address_id"`
	ShippingAddress   Address        `gorm:"foreignKey:ShippingAddressID" json:"shipping_address"`
	PaymentID         string         `json:"payment_id"`
	TrackingNumber    string         `json:"tracking_number"`
	Notes             string         `json:"notes"`
}

type OrderItem struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	OrderID   uint           `gorm:"not null" json:"order_id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Product   Product        `json:"product"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	Price     float64        `gorm:"not null" json:"price"` // Price at time of purchase
	Subtotal  float64        `gorm:"not null" json:"subtotal"`
}
