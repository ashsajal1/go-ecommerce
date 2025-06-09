package service

import (
	"errors"

	"github.com/sajal/go-ecommerce/internal/models"
	"github.com/sajal/go-ecommerce/internal/repository"
)

type OrderService struct {
	repo     *repository.OrderRepository
	cartRepo *repository.CartRepository
}

func NewOrderService(repo *repository.OrderRepository, cartRepo *repository.CartRepository) *OrderService {
	return &OrderService{
		repo:     repo,
		cartRepo: cartRepo,
	}
}

func (s *OrderService) CreateOrder(userID uint) (*models.Order, error) {
	// Get user's cart
	cart, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("cart not found")
	}

	if len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	// Create order items from cart items
	var orderItems []models.OrderItem
	var total float64

	for _, item := range cart.Items {
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Subtotal:  item.Subtotal,
		})
		total += item.Subtotal
	}

	// Create order
	order := &models.Order{
		UserID:      userID,
		Items:       orderItems,
		TotalAmount: total,
		Status:      models.OrderStatusPending,
	}

	if err := s.repo.Create(order); err != nil {
		return nil, err
	}

	// Clear the cart
	if err := s.cartRepo.ClearCart(userID); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetOrder(id uint, userID uint) (*models.Order, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("order not found")
	}

	// Check if order belongs to user
	if order.UserID != userID {
		return nil, errors.New("unauthorized access")
	}

	return order, nil
}

func (s *OrderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.repo.FindByUserID(userID)
}

func (s *OrderService) UpdateOrderStatus(id uint, status models.OrderStatus) error {
	// Validate status
	switch status {
	case models.OrderStatusPending,
		models.OrderStatusProcessing,
		models.OrderStatusShipped,
		models.OrderStatusDelivered,
		models.OrderStatusCancelled:
	default:
		return errors.New("invalid order status")
	}

	return s.repo.UpdateStatus(id, status)
}

func (s *OrderService) CancelOrder(id uint, userID uint) error {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("order not found")
	}

	// Check if order belongs to user
	if order.UserID != userID {
		return errors.New("unauthorized access")
	}

	// Check if order can be cancelled
	if order.Status != models.OrderStatusPending {
		return errors.New("order cannot be cancelled")
	}

	return s.repo.UpdateStatus(id, models.OrderStatusCancelled)
}
