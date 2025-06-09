package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajal/go-ecommerce/internal/models"
)

// GetCart godoc
// @Summary Get user's cart
// @Description Get the current user's shopping cart with all items
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} Response
// @Router /cart [get]
func (h *Handler) GetCart(c *gin.Context) {
	userID := c.GetUint("user_id") // Set by auth middleware
	var cart models.Cart

	if err := h.db.Preload("Items.Product").First(&cart, "user_id = ?", userID).Error; err != nil {
		h.errorResponse(c, http.StatusNotFound, "Cart not found")
		return
	}

	h.successResponse(c, cart, "Cart retrieved successfully")
}

// AddToCart godoc
// @Summary Add item to cart
// @Description Add a product to the user's shopping cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param item body models.CartItem true "Cart item details"
// @Success 200 {object} Response
// @Router /cart/items [post]
func (h *Handler) AddToCart(c *gin.Context) {
	userID := c.GetUint("user_id")
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	// Get or create cart
	var cart models.Cart
	if err := h.db.FirstOrCreate(&cart, models.Cart{UserID: userID}).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to get/create cart")
		return
	}

	// Check if product exists
	var product models.Product
	if err := h.db.First(&product, cartItem.ProductID).Error; err != nil {
		h.errorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	// Check if item already exists in cart
	var existingItem models.CartItem
	if err := h.db.Where("cart_id = ? AND product_id = ?", cart.ID, cartItem.ProductID).First(&existingItem).Error; err == nil {
		// Update quantity if item exists
		existingItem.Quantity += cartItem.Quantity
		existingItem.Price = product.Price
		existingItem.Subtotal = float64(existingItem.Quantity) * existingItem.Price
		if err := h.db.Save(&existingItem).Error; err != nil {
			h.errorResponse(c, http.StatusInternalServerError, "Failed to update cart item")
			return
		}
		h.successResponse(c, existingItem, "Cart item updated successfully")
		return
	}

	// Add new item
	cartItem.CartID = cart.ID
	cartItem.Price = product.Price
	cartItem.Subtotal = float64(cartItem.Quantity) * cartItem.Price

	if err := h.db.Create(&cartItem).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to add item to cart")
		return
	}

	h.successResponse(c, cartItem, "Item added to cart successfully")
}

// UpdateCartItem godoc
// @Summary Update cart item
// @Description Update quantity of an item in the cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Cart Item ID"
// @Param item body models.CartItem true "Updated cart item details"
// @Success 200 {object} Response
// @Router /cart/items/{id} [put]
func (h *Handler) UpdateCartItem(c *gin.Context) {
	userID := c.GetUint("user_id")
	itemID := c.Param("id")

	var cartItem models.CartItem
	if err := h.db.Joins("Cart").Where("cart_items.id = ? AND cart.user_id = ?", itemID, userID).First(&cartItem).Error; err != nil {
		h.errorResponse(c, http.StatusNotFound, "Cart item not found")
		return
	}

	if err := c.ShouldBindJSON(&cartItem); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	cartItem.Subtotal = float64(cartItem.Quantity) * cartItem.Price
	if err := h.db.Save(&cartItem).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to update cart item")
		return
	}

	h.successResponse(c, cartItem, "Cart item updated successfully")
}

// RemoveFromCart godoc
// @Summary Remove item from cart
// @Description Remove an item from the shopping cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Cart Item ID"
// @Success 204 "No Content"
// @Router /cart/items/{id} [delete]
func (h *Handler) RemoveFromCart(c *gin.Context) {
	userID := c.GetUint("user_id")
	itemID := c.Param("id")

	if err := h.db.Joins("Cart").Where("cart_items.id = ? AND cart.user_id = ?", itemID, userID).Delete(&models.CartItem{}).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to remove item from cart")
		return
	}

	h.noContentResponse(c)
}

// ClearCart godoc
// @Summary Clear cart
// @Description Remove all items from the shopping cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 204 "No Content"
// @Router /cart [delete]
func (h *Handler) ClearCart(c *gin.Context) {
	userID := c.GetUint("user_id")

	if err := h.db.Exec("DELETE FROM cart_items WHERE cart_id IN (SELECT id FROM carts WHERE user_id = ?)", userID).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to clear cart")
		return
	}

	h.noContentResponse(c)
}
