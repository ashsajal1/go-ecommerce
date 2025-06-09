package repository

import (
	"github.com/sajal/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	DB *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{DB: db}
}

func (r *ReviewRepository) Create(review *models.Review) error {
	return r.DB.Create(review).Error
}

func (r *ReviewRepository) FindByID(id uint) (*models.Review, error) {
	var review models.Review
	err := r.DB.Preload("User").Preload("Product").First(&review, id).Error
	return &review, err
}

func (r *ReviewRepository) FindByProductID(productID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.DB.Preload("User").Where("product_id = ?", productID).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepository) FindByUserID(userID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.DB.Preload("Product").Where("user_id = ?", userID).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepository) Update(review *models.Review) error {
	return r.DB.Save(review).Error
}

func (r *ReviewRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Review{}, id).Error
}

func (r *ReviewRepository) FindByUserAndProduct(userID, productID uint) (*models.Review, error) {
	var review models.Review
	err := r.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&review).Error
	return &review, err
}
