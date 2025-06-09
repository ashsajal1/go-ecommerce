package repository

import (
	"github.com/sajal/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").Preload("Images").First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) List(filters map[string]interface{}) ([]models.Product, error) {
	var products []models.Product
	query := r.db.Model(&models.Product{})

	if categoryID, ok := filters["category_id"]; ok {
		query = query.Where("category_id = ?", categoryID)
	}
	if minPrice, ok := filters["min_price"]; ok {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice, ok := filters["max_price"]; ok {
		query = query.Where("price <= ?", maxPrice)
	}
	if search, ok := filters["search"]; ok {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search.(string)+"%", "%"+search.(string)+"%")
	}

	err := query.Preload("Category").Preload("Images").Find(&products).Error
	return products, err
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
