package service

import (
	"errors"

	"github.com/sajal/go-ecommerce/internal/models"
	"github.com/sajal/go-ecommerce/internal/repository"
)

type AddressService struct {
	repo *repository.AddressRepository
}

func NewAddressService(repo *repository.AddressRepository) *AddressService {
	return &AddressService{repo: repo}
}

func (s *AddressService) CreateAddress(address *models.Address) error {
	// Validate address data
	if address.Street == "" {
		return errors.New("street is required")
	}
	if address.City == "" {
		return errors.New("city is required")
	}
	if address.State == "" {
		return errors.New("state is required")
	}
	if address.Country == "" {
		return errors.New("country is required")
	}
	if address.ZipCode == "" {
		return errors.New("zip code is required")
	}
	if address.Type == "" {
		return errors.New("address type is required")
	}
	if address.Type != "shipping" && address.Type != "billing" {
		return errors.New("address type must be either shipping or billing")
	}

	// If this is the first address, set it as default
	addresses, err := s.repo.FindByUserID(address.UserID)
	if err != nil {
		return err
	}
	if len(addresses) == 0 {
		address.IsDefault = true
	}

	return s.repo.Create(address)
}

func (s *AddressService) GetAddress(id uint) (*models.Address, error) {
	return s.repo.FindByID(id)
}

func (s *AddressService) GetUserAddresses(userID uint) ([]models.Address, error) {
	return s.repo.FindByUserID(userID)
}

func (s *AddressService) UpdateAddress(address *models.Address) error {
	// Check if address exists
	existingAddress, err := s.repo.FindByID(address.ID)
	if err != nil {
		return errors.New("address not found")
	}

	// Validate address data
	if address.Street == "" {
		return errors.New("street is required")
	}
	if address.City == "" {
		return errors.New("city is required")
	}
	if address.State == "" {
		return errors.New("state is required")
	}
	if address.Country == "" {
		return errors.New("country is required")
	}
	if address.ZipCode == "" {
		return errors.New("zip code is required")
	}
	if address.Type == "" {
		return errors.New("address type is required")
	}
	if address.Type != "shipping" && address.Type != "billing" {
		return errors.New("address type must be either shipping or billing")
	}

	// Preserve some fields
	address.UserID = existingAddress.UserID
	address.CreatedAt = existingAddress.CreatedAt
	address.IsDefault = existingAddress.IsDefault

	return s.repo.Update(address)
}

func (s *AddressService) DeleteAddress(id uint) error {
	// Check if address exists
	address, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("address not found")
	}

	// If this is the default address, make another address default
	if address.IsDefault {
		addresses, err := s.repo.FindByUserID(address.UserID)
		if err != nil {
			return err
		}

		// Find another address to make default
		for _, addr := range addresses {
			if addr.ID != id {
				if err := s.repo.SetDefault(address.UserID, addr.ID); err != nil {
					return err
				}
				break
			}
		}
	}

	return s.repo.Delete(id)
}

func (s *AddressService) SetDefaultAddress(userID uint, addressID uint) error {
	// Check if address exists and belongs to user
	address, err := s.repo.FindByID(addressID)
	if err != nil {
		return errors.New("address not found")
	}
	if address.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.repo.SetDefault(userID, addressID)
}
