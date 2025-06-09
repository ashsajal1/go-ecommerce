package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sajal/go-ecommerce/internal/api"
	"github.com/sajal/go-ecommerce/internal/config"
	"github.com/sajal/go-ecommerce/internal/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db := config.InitDB(cfg)

	// Initialize router
	router := gin.Default()

	// Apply global middlewares
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Initialize API handler
	handler := api.NewHandler(db)

	// Setup routes
	handler.SetupRoutes(router)

	// Start server
	serverAddr := ":" + cfg.Port
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
