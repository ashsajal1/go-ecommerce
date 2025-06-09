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
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (s *ProductService) ListProducts(filters map[string]interface{}) ([]models.Product, error) {
	// Validate filters
	if minPrice, ok := filters["min_price"].(float64); ok && minPrice < 0 {
		return nil, errors.New("minimum price cannot be negative")
	}
	if maxPrice, ok := filters["max_price"].(float64); ok && maxPrice < 0 {
		return nil, errors.New("maximum price cannot be negative")
	}

	return s.repo.List(filters)
}

func (s *ProductService) UpdateProduct(id uint, updates *models.Product) error {
	// Check if product exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	// Apply updates
	if updates.Name != "" {
		existing.Name = updates.Name
	}
	if updates.Description != "" {
		existing.Description = updates.Description
	}
	if updates.Price > 0 {
		existing.Price = updates.Price
	}
	if updates.Stock >= 0 {
		existing.Stock = updates.Stock
	}
	if updates.CategoryID > 0 {
		existing.CategoryID = updates.CategoryID
	}

	return s.repo.Update(existing)
}

func (s *ProductService) DeleteProduct(id uint) error {
	// Check if product exists
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("product not found")
	}

	return s.repo.Delete(id)
}
