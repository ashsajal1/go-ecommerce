package service

import (
	"errors"

	"github.com/sajal/go-ecommerce/internal/models"
	"github.com/sajal/go-ecommerce/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	// Validate product data
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	return s.repo.Create(product)
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetProductsByCategory(categoryID uint) ([]models.Product, error) {
	return s.repo.FindByCategory(categoryID)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	// Check if product exists
	existingProduct, err := s.repo.FindByID(product.ID)
	if err != nil {
		return errors.New("product not found")
	}

	// Validate product data
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	// Preserve some fields
	product.CreatedAt = existingProduct.CreatedAt

	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	// Check if product exists
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	return s.repo.Delete(id)
}

func (s *ProductService) SearchProducts(query string) ([]models.Product, error) {
	return s.repo.Search(query)
}

func (s *ProductService) UpdateStock(id uint, quantity int) error {
	// Check if product exists
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	// Validate quantity
	if quantity < 0 {
		return errors.New("stock quantity cannot be negative")
	}

	return s.repo.UpdateStock(id, quantity)
}
