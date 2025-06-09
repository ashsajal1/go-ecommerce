package api

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutes(r *gin.Engine) {
	// Public routes
	public := r.Group("/api/v1")
	{
		// Product routes
		products := public.Group("/products")
		{
			products.GET("", h.ListProducts)
			products.GET("/:id", h.GetProduct)
		}

		// Auth routes
		auth := public.Group("/auth")
		{
			auth.POST("/register", h.Register)
			auth.POST("/login", h.Login)
		}
	}

	// Protected routes
	protected := r.Group("/api/v1")
	protected.Use(h.AuthMiddleware())
	{
		// User routes
		users := protected.Group("/users")
		{
			users.GET("/me", h.GetCurrentUser)
			users.PUT("/me", h.UpdateUser)
		}

		// Cart routes
		cart := protected.Group("/cart")
		{
			cart.GET("", h.GetCart)
			cart.DELETE("", h.ClearCart)
			cart.POST("/items", h.AddToCart)
			cart.PUT("/items/:id", h.UpdateCartItem)
			cart.DELETE("/items/:id", h.RemoveFromCart)
		}

		// Order routes
		orders := protected.Group("/orders")
		{
			orders.POST("", h.CreateOrder)
			orders.GET("", h.GetOrders)
			orders.GET("/:id", h.GetOrder)
			orders.POST("/:id/cancel", h.CancelOrder)
		}

		// Admin routes
		admin := protected.Group("/admin")
		admin.Use(h.AdminMiddleware())
		{
			// Product management
			admin.POST("/products", h.CreateProduct)
			admin.PUT("/products/:id", h.UpdateProduct)
			admin.DELETE("/products/:id", h.DeleteProduct)

			// Order management
			admin.PUT("/orders/:id/status", h.UpdateOrderStatus)
		}
	}
}
