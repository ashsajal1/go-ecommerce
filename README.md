# Go E-commerce API

A modern e-commerce REST API built with Go, Gin, and GORM.

## Features

- User authentication and authorization (JWT)
- Product management with categories
- Shopping cart functionality
- Order processing and management
- Admin dashboard for product and order management
- RESTful API design
- Clean architecture with separation of concerns

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
│   ├── api/              # API handlers and routes
│   │   ├── auth.go       # Authentication handlers
│   │   ├── cart.go       # Cart handlers
│   │   ├── handler.go    # Base handler and initialization
│   │   ├── order.go      # Order handlers
│   │   ├── product.go    # Product handlers
│   │   └── routes.go     # Route definitions
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

## API Endpoints

### Public Routes
- `GET /api/v1/products` - List all products
- `GET /api/v1/products/:id` - Get product details
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - User login

### Protected Routes
- `GET /api/v1/users/me` - Get current user
- `PUT /api/v1/users/me` - Update user profile

### Cart Routes
- `GET /api/v1/cart` - Get cart
- `POST /api/v1/cart/items` - Add item to cart
- `PUT /api/v1/cart/items/:id` - Update cart item
- `DELETE /api/v1/cart/items/:id` - Remove item from cart
- `DELETE /api/v1/cart` - Clear cart

### Order Routes
- `POST /api/v1/orders` - Create order
- `GET /api/v1/orders` - List user orders
- `GET /api/v1/orders/:id` - Get order details
- `POST /api/v1/orders/:id/cancel` - Cancel order

### Admin Routes
- `POST /api/v1/admin/products` - Create product
- `PUT /api/v1/admin/products/:id` - Update product
- `DELETE /api/v1/admin/products/:id` - Delete product
- `PUT /api/v1/admin/orders/:id/status` - Update order status

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

## Development

The project follows clean architecture principles with clear separation of concerns:
- Handlers: Handle HTTP requests and responses
- Services: Implement business logic
- Repositories: Handle database operations
- Models: Define data structures

## License

MIT 