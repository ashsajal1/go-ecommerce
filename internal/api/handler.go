package api

import (
	"github.com/sajal/go-ecommerce/internal/repository"
	"github.com/sajal/go-ecommerce/internal/service"
	"gorm.io/gorm"
)

type Handler struct {
	db             *gorm.DB
	productHandler *ProductHandler
	userHandler    *UserHandler
	cartHandler    *CartHandler
	orderHandler   *OrderHandler
	reviewHandler  *ReviewHandler
	addressHandler *AddressHandler
}

type UserHandler struct {
	*Handler
	service *service.UserService
}

type CartHandler struct {
	*Handler
	service *service.CartService
}

type OrderHandler struct {
	*Handler
	service *service.OrderService
}

type ReviewHandler struct {
	*Handler
	service *service.ReviewService
}

type AddressHandler struct {
	*Handler
	service *service.AddressService
}

func NewUserHandler(handler *Handler, service *service.UserService) *UserHandler {
	return &UserHandler{
		Handler: handler,
		service: service,
	}
}

func NewCartHandler(handler *Handler, service *service.CartService) *CartHandler {
	return &CartHandler{
		Handler: handler,
		service: service,
	}
}

func NewOrderHandler(handler *Handler, service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		Handler: handler,
		service: service,
	}
}

func NewReviewHandler(handler *Handler, service *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		Handler: handler,
		service: service,
	}
}

func NewAddressHandler(handler *Handler, service *service.AddressService) *AddressHandler {
	return &AddressHandler{
		Handler: handler,
		service: service,
	}
}

func NewHandler(db *gorm.DB) *Handler {
	// Initialize repositories
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)
	cartRepo := repository.NewCartRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	addressRepo := repository.NewAddressRepository(db)

	// Initialize services
	productService := service.NewProductService(productRepo)
	userService := service.NewUserService(userRepo)
	cartService := service.NewCartService(cartRepo, productRepo)
	orderService := service.NewOrderService(orderRepo, cartRepo)
	reviewService := service.NewReviewService(reviewRepo, productRepo, orderRepo)
	addressService := service.NewAddressService(addressRepo)

	// Create base handler
	handler := &Handler{
		db: db,
	}

	// Initialize specific handlers
	handler.productHandler = NewProductHandler(handler, productService)
	handler.userHandler = NewUserHandler(handler, userService)
	handler.cartHandler = NewCartHandler(handler, cartService)
	handler.orderHandler = NewOrderHandler(handler, orderService)
	handler.reviewHandler = NewReviewHandler(handler, reviewService)
	handler.addressHandler = NewAddressHandler(handler, addressService)

	return handler
}
