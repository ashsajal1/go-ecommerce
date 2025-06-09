package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sajal/go-ecommerce/internal/models"
	"github.com/sajal/go-ecommerce/internal/service"
)

type ProductHandler struct {
	*Handler
	service *service.ProductService
}

func NewProductHandler(handler *Handler, service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		Handler: handler,
		service: service,
	}
}

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
func (h *ProductHandler) ListProducts(c *gin.Context) {
	filters := make(map[string]interface{})

	if categoryID := c.Query("category_id"); categoryID != "" {
		if id, err := strconv.ParseUint(categoryID, 10, 32); err == nil {
			filters["category_id"] = uint(id)
		}
	}
	if minPrice := c.Query("min_price"); minPrice != "" {
		if price, err := strconv.ParseFloat(minPrice, 64); err == nil {
			filters["min_price"] = price
		}
	}
	if maxPrice := c.Query("max_price"); maxPrice != "" {
		if price, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			filters["max_price"] = price
		}
	}
	if search := c.Query("search"); search != "" {
		filters["search"] = search
	}

	products, err := h.service.ListProducts(filters)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
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
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.service.GetProduct(uint(id))
	if err != nil {
		h.errorResponse(c, http.StatusNotFound, err.Error())
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
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
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
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.service.UpdateProduct(uint(id), &product); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
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
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.noContentResponse(c)
}
