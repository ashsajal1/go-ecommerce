package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajal/go-ecommerce/internal/models"
)

// ListProducts godoc
// @Summary List all products
// @Description Get all products with optional filtering
// @Tags products
// @Accept json
// @Produce json
// @Param category_id query int false "Filter by category ID"
// @Param min_price query float false "Minimum price"
// @Param max_price query float false "Maximum price"
// @Param search query string false "Search term"
// @Success 200 {object} Response
// @Router /products [get]
func (h *Handler) ListProducts(c *gin.Context) {
	var products []models.Product
	query := h.db.Model(&models.Product{})

	// Apply filters
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if minPrice := c.Query("min_price"); minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice := c.Query("max_price"); maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}
	if search := c.Query("search"); search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Preload("Category").Preload("Images").Find(&products).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	h.successResponse(c, products, "Products retrieved successfully")
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get detailed information about a specific product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} Response
// @Router /products/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := h.db.Preload("Category").Preload("Images").First(&product, id).Error; err != nil {
		h.errorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	h.successResponse(c, product, "Product retrieved successfully")
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product details"
// @Success 201 {object} Response
// @Router /products [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.db.Create(&product).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	h.createdResponse(c, product)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update an existing product's details
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Updated product details"
// @Success 200 {object} Response
// @Router /products/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := h.db.First(&product, id).Error; err != nil {
		h.errorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.db.Save(&product).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to update product")
		return
	}

	h.successResponse(c, product, "Product updated successfully")
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Router /products/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.db.Delete(&models.Product{}, id).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	h.noContentResponse(c)
}
