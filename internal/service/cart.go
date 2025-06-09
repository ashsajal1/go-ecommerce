package service

import (
	"errors"

	"github.com/sajal/go-ecommerce/internal/models"
	"github.com/sajal/go-ecommerce/internal/repository"
)

type CartService struct {
	repo        *repository.CartRepository
	productRepo *repository.ProductRepository
}

func NewCartService(repo *repository.CartRepository, productRepo *repository.ProductRepository) *CartService {
	return &CartService{
		repo:        repo,
		productRepo: productRepo,
	}
}

func (s *CartService) GetCart(userID uint) (*models.Cart, error) {
	cart, err := s.repo.FindByUserID(userID)
	if err != nil {
		// Create new cart if not exists
		cart = &models.Cart{UserID: userID}
		if err := s.repo.Create(cart); err != nil {
			return nil, err
		}
	}
	return cart, nil
}

func (s *CartService) AddToCart(userID uint, productID uint, quantity int) error {
	// Check if product exists
	product, err := s.productRepo.FindByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	// Check stock
	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	// Get or create cart
	cart, err := s.GetCart(userID)
	if err != nil {
		return err
	}

	// Check if item already exists in cart
	existingItem, err := s.repo.FindCartItem(cart.ID, productID)
	if err == nil {
		// Update quantity if item exists
		existingItem.Quantity += quantity
		existingItem.Subtotal = float64(existingItem.Quantity) * existingItem.Price
		return s.repo.UpdateItem(existingItem)
	}

	// Add new item
	item := &models.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     product.Price,
		Subtotal:  float64(quantity) * product.Price,
	}

	return s.repo.AddItem(item)
}

func (s *CartService) UpdateCartItem(userID uint, itemID uint, quantity int) error {
	// Get cart
	cart, err := s.GetCart(userID)
	if err != nil {
		return err
	}

	// Get item
	item, err := s.repo.FindCartItemByID(itemID, cart.ID)
	if err != nil {
		return errors.New("item not found")
	}

	// Check stock
	product, err := s.productRepo.FindByID(item.ProductID)
	if err != nil {
		return errors.New("product not found")
	}
	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	// Update item
	item.Quantity = quantity
	item.Subtotal = float64(quantity) * item.Price
	return s.repo.UpdateItem(item)
}

func (s *CartService) RemoveFromCart(userID uint, itemID uint) error {
	// Get cart
	cart, err := s.GetCart(userID)
	if err != nil {
		return err
	}

	// Check if item exists in cart
	_, err = s.repo.FindCartItemByID(itemID, cart.ID)
	if err != nil {
		return errors.New("item not found")
	}

	return s.repo.RemoveItem(itemID)
}

func (s *CartService) ClearCart(userID uint) error {
	return s.repo.ClearCart(userID)
}
