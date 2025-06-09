package service

import (
	"errors"

	"github.com/sajal/go-ecommerce/internal/models"
	"github.com/sajal/go-ecommerce/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *models.User) error {
	// Check if email already exists
	existing, err := s.repo.FindByEmail(user.Email)
	if err == nil && existing != nil {
		return errors.New("email already registered")
	}

	// Set default role if not specified
	if user.Role == "" {
		user.Role = "user"
	}

	return s.repo.Create(user)
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) UpdateUser(id uint, updates *models.User) error {
	// Check if user exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	// Prevent role update
	updates.Role = existing.Role

	// Apply updates
	if updates.Name != "" {
		existing.Name = updates.Name
	}
	if updates.Email != "" {
		existing.Email = updates.Email
	}
	if updates.Password != "" {
		existing.Password = updates.Password
	}

	return s.repo.Update(existing)
}

func (s *UserService) DeleteUser(id uint) error {
	// Check if user exists
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("user not found")
	}

	return s.repo.Delete(id)
}
