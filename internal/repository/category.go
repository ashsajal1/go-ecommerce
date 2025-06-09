package repository

import (
	"github.com/sajal/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.DB.Preload("Products").First(&category, id).Error
	return &category, err
}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Preload("Products").Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) Update(category *models.Category) error {
	return r.DB.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Category{}, id).Error
}

func (r *CategoryRepository) FindByName(name string) (*models.Category, error) {
	var category models.Category
	err := r.DB.Where("name = ?", name).First(&category).Error
	return &category, err
}
