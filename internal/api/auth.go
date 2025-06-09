package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sajal/go-ecommerce/internal/models"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterInput true "User registration details"
// @Success 201 {object} Response
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := h.db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		h.errorResponse(c, http.StatusConflict, "Email already registered")
		return
	}

	// Create user
	user := models.User{
		Email:    input.Email,
		Password: input.Password, // Will be hashed by BeforeSave hook
		Name:     input.Name,
		Role:     "user",
	}

	if err := h.db.Create(&user).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	h.createdResponse(c, user)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body LoginInput true "Login credentials"
// @Success 200 {object} Response
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	// Find user
	var user models.User
	if err := h.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		h.errorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Check password
	if !user.CheckPassword(input.Password) {
		h.errorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your-secret-key")) // Use config.JWTSecret in production
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	h.successResponse(c, gin.H{"token": tokenString}, "Login successful")
}

// AuthMiddleware is a middleware to check JWT token
func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			h.errorResponse(c, http.StatusUnauthorized, "Authorization header required")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil // Use config.JWTSecret in production
		})

		if err != nil || !token.Valid {
			h.errorResponse(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("user_role", claims["role"].(string))
		c.Next()
	}
}

// AdminMiddleware is a middleware to check if user is admin
func (h *Handler) AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("user_role")
		if role != "admin" {
			h.errorResponse(c, http.StatusForbidden, "Admin access required")
			c.Abort()
			return
		}
		c.Next()
	}
}

// GetCurrentUser godoc
// @Summary Get current user
// @Description Get details of the currently authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} Response
// @Router /users/me [get]
func (h *Handler) GetCurrentUser(c *gin.Context) {
	userID := c.GetUint("user_id")
	var user models.User

	if err := h.db.First(&user, userID).Error; err != nil {
		h.errorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	h.successResponse(c, user, "User retrieved successfully")
}

// UpdateUser godoc
// @Summary Update current user
// @Description Update details of the currently authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body models.User true "Updated user details"
// @Success 200 {object} Response
// @Router /users/me [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	userID := c.GetUint("user_id")
	var user models.User

	if err := h.db.First(&user, userID).Error; err != nil {
		h.errorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	// Prevent role update
	user.Role = "user"

	if err := h.db.Save(&user).Error; err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "Failed to update user")
		return
	}

	h.successResponse(c, user, "User updated successfully")
}
