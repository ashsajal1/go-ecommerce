package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajal/go-ecommerce/internal/models"
)

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order from the user's cart
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body models.Order true "Order details"
// @Success 201 {object} Response
// @Router /orders [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	userID := c.GetUint("user_id")
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	// Get user's cart
	var cart models.Cart
	if err := h.db.Preload("Items.Product").First(&cart, "user_id = ?", userID).Error; err != nil {
		h.errorResponse(c, http.StatusNotFound, "Cart not found")
		return
	}

	// Create order items from cart items
	var orderItems []models.OrderItem
	for _, item := range cart.Items {
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Subtotal:  item.Subtotal,
		})
	}

	// Create order
	order.UserID = userID
	order.Items = orderItems
	order.Status = models.OrderStatusPending

	if err := h.db.Create(&order).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to create order")
		return
	}

	// Clear the cart
	if err := h.db.Delete(&cart).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to clear cart")
		return
	}

	h.createdResponse(c, order)
}

// GetOrders godoc
// @Summary List user's orders
// @Description Get all orders for the current user
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} Response
// @Router /orders [get]
func (h *Handler) GetOrders(c *gin.Context) {
	userID := c.GetUint("user_id")
	var orders []models.Order

	if err := h.db.Preload("Items.Product").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to fetch orders")
		return
	}

	h.successResponse(c, orders, "Orders retrieved successfully")
}

// GetOrder godoc
// @Summary Get order details
// @Description Get detailed information about a specific order
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} Response
// @Router /orders/{id} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	userID := c.GetUint("user_id")
	orderID := c.Param("id")
	var order models.Order

	if err := h.db.Preload("Items.Product").Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		h.errorResponse(c, http.StatusNotFound, "Order not found")
		return
	}

	h.successResponse(c, order, "Order retrieved successfully")
}

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Update the status of an order (admin only)
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Param status body string true "New order status"
// @Success 200 {object} Response
// @Router /orders/{id}/status [put]
func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	var input struct {
		Status models.OrderStatus `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", input.Status).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to update order status")
		return
	}

	h.successResponse(c, nil, "Order status updated successfully")
}

// CancelOrder godoc
// @Summary Cancel an order
// @Description Cancel an existing order
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} Response
// @Router /orders/{id}/cancel [post]
func (h *Handler) CancelOrder(c *gin.Context) {
	userID := c.GetUint("user_id")
	orderID := c.Param("id")

	if err := h.db.Model(&models.Order{}).Where("id = ? AND user_id = ?", orderID, userID).Update("status", models.OrderStatusCancelled).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to cancel order")
		return
	}

	h.successResponse(c, nil, "Order cancelled successfully")
}
