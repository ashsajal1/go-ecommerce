package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port            string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	DBSSLMode       string
	JWTSecret       string
	JWTExpiration   string
	StripeSecretKey string
	UploadDir       string
	MaxFileSize     int64
}

func LoadConfig() *Config {
	return &Config{
		Port:            getEnv("PORT", "8080"),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPassword:      getEnv("DB_PASSWORD", "postgres"),
		DBName:          getEnv("DB_NAME", "ecommerce"),
		DBSSLMode:       getEnv("DB_SSL_MODE", "disable"),
		JWTSecret:       getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiration:   getEnv("JWT_EXPIRATION", "24h"),
		StripeSecretKey: getEnv("STRIPE_SECRET_KEY", ""),
		UploadDir:       getEnv("UPLOAD_DIR", "uploads"),
		MaxFileSize:     getEnvAsInt64("MAX_FILE_SIZE", 5242880), // 5MB default
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}
