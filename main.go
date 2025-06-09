package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sajal/go-ecommerce/internal/api"
	"github.com/sajal/go-ecommerce/internal/config"
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
