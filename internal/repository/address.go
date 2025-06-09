package repository

import (
	"github.com/sajal/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{DB: db}
}

func (r *AddressRepository) Create(address *models.Address) error {
	return r.DB.Create(address).Error
}

func (r *AddressRepository) FindByID(id uint) (*models.Address, error) {
	var address models.Address
	err := r.DB.Preload("User").First(&address, id).Error
	return &address, err
}

func (r *AddressRepository) FindByUserID(userID uint) ([]models.Address, error) {
	var addresses []models.Address
	err := r.DB.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (r *AddressRepository) Update(address *models.Address) error {
	return r.DB.Save(address).Error
}

func (r *AddressRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Address{}, id).Error
}

func (r *AddressRepository) SetDefault(userID uint, addressID uint) error {
	// Reset all addresses to non-default
	if err := r.DB.Model(&models.Address{}).Where("user_id = ?", userID).Update("is_default", false).Error; err != nil {
		return err
	}

	// Set the specified address as default
	return r.DB.Model(&models.Address{}).Where("id = ? AND user_id = ?", addressID, userID).Update("is_default", true).Error
}
