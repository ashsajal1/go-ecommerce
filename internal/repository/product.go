package repository

import (
	"github.com/sajal/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.DB.Preload("Category").Preload("Reviews").First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Category").Find(&products).Error
	return products, err
}

func (r *ProductRepository) FindByCategory(categoryID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Category").Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}

func (r *ProductRepository) Search(query string) ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Category").
		Where("name LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").
		Find(&products).Error
	return products, err
}

func (r *ProductRepository) UpdateStock(id uint, quantity int) error {
	return r.DB.Model(&models.Product{}).Where("id = ?", id).Update("stock", quantity).Error
}
