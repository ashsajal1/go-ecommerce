package repository

import (
	"github.com/sajal/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (r *CartRepository) FindByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.DB.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	return &cart, err
}

func (r *CartRepository) Create(cart *models.Cart) error {
	return r.DB.Create(cart).Error
}

func (r *CartRepository) AddItem(item *models.CartItem) error {
	return r.DB.Create(item).Error
}

func (r *CartRepository) UpdateItem(item *models.CartItem) error {
	return r.DB.Save(item).Error
}

func (r *CartRepository) RemoveItem(id uint) error {
	return r.DB.Delete(&models.CartItem{}, id).Error
}

func (r *CartRepository) ClearCart(userID uint) error {
	return r.DB.Exec("DELETE FROM cart_items WHERE cart_id IN (SELECT id FROM carts WHERE user_id = ?)", userID).Error
}

func (r *CartRepository) FindCartItem(cartID uint, productID uint) (*models.CartItem, error) {
	var item models.CartItem
	err := r.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error
	return &item, err
}

func (r *CartRepository) FindCartItemByID(itemID uint, cartID uint) (*models.CartItem, error) {
	var item models.CartItem
	err := r.DB.Where("id = ? AND cart_id = ?", itemID, cartID).First(&item).Error
	return &item, err
}
