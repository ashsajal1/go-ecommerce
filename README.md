# Go E-commerce API

A modern e-commerce REST API built with Go, Gin, and GORM.

## Features

- User authentication and authorization
- Product management
- Category management
- Order processing
- Shopping cart functionality
- Payment integration (stripe)
- File upload for product images
- API documentation with Swagger

## Tech Stack

- Go 1.21+
- Gin Web Framework
- GORM (ORM)
- PostgreSQL
- JWT Authentication
- Docker

## Project Structure

```
.
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── api/              # API handlers
│   ├── config/           # Configuration
│   ├── middleware/       # Custom middleware
│   ├── models/           # Database models
│   ├── repository/       # Database operations
│   └── service/          # Business logic
├── pkg/                  # Public library code
├── docs/                 # Documentation
├── scripts/              # Build and deployment scripts
├── .env.example         # Example environment variables
├── go.mod               # Go module file
└── main.go              # Main application entry point
```

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and update the values
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

## API Documentation

API documentation is available at `/swagger/index.html` when running the server.

## License

MIT 