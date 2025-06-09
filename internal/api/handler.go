package api

import (
	"github.com/sajal/go-ecommerce/internal/repository"
	"github.com/sajal/go-ecommerce/internal/service"
	"gorm.io/gorm"
)

type Handler struct {
	db             *gorm.DB
	productHandler *ProductHandler
}

func NewHandler(db *gorm.DB) *Handler {
	// Initialize repositories
	productRepo := repository.NewProductRepository(db)

	// Initialize services
	productService := service.NewProductService(productRepo)

	// Create base handler
	handler := &Handler{
		db: db,
	}

	// Initialize specific handlers
	handler.productHandler = NewProductHandler(handler, productService)

	return handler
}
