package service

import (
	"errors"

	"github.com/sajal/go-ecommerce/internal/models"
	"github.com/sajal/go-ecommerce/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	// Validate category data
	if category.Name == "" {
		return errors.New("category name is required")
	}

	// Check if category with same name exists
	_, err := s.repo.FindByName(category.Name)
	if err == nil {
		return errors.New("category with this name already exists")
	}

	return s.repo.Create(category)
}

func (s *CategoryService) GetCategory(id uint) (*models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {
	// Check if category exists
	existingCategory, err := s.repo.FindByID(category.ID)
	if err != nil {
		return errors.New("category not found")
	}

	// Validate category data
	if category.Name == "" {
		return errors.New("category name is required")
	}

	// Check if another category with same name exists
	if category.Name != existingCategory.Name {
		duplicateCategory, err := s.repo.FindByName(category.Name)
		if err == nil && duplicateCategory.ID != category.ID {
			return errors.New("category with this name already exists")
		}
	}

	// Preserve some fields
	category.CreatedAt = existingCategory.CreatedAt

	return s.repo.Update(category)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	// Check if category exists
	category, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("category not found")
	}

	// Check if category has products
	if len(category.Products) > 0 {
		return errors.New("cannot delete category with associated products")
	}

	return s.repo.Delete(id)
}
