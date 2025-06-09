package repository

import (
	"github.com/sajal/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.DB.Preload("User").Preload("Items").Preload("Items.Product").First(&order, id).Error
	return &order, err
}

func (r *OrderRepository) FindByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Preload("Items").Preload("Items.Product").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) Update(order *models.Order) error {
	return r.DB.Save(order).Error
}

func (r *OrderRepository) UpdateStatus(id uint, status models.OrderStatus) error {
	return r.DB.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *OrderRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Order{}, id).Error
}

func (r *OrderRepository) FindAll() ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Preload("User").Preload("Items").Preload("Items.Product").Find(&orders).Error
	return orders, err
}
