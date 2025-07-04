# Blog Management API

A comprehensive blog management system built with Go Fiber and PostgreSQL, following clean architecture principles with controller, service, and repository patterns.

## 🚀 Features

- **CRUD Operations**: Create, Read, Update, and Delete blog posts
- **Clean Architecture**: Controller, Service, Repository pattern
- **PostgreSQL Database**: Robust data persistence
- **RESTful API**: Standard HTTP methods and status codes
- **Input Validation**: Comprehensive request validation
- **Error Handling**: Centralized error handling and logging
- **Unit Tests**: High test coverage with mocking
- **CORS Support**: Cross-origin resource sharing enabled
- **Environment Configuration**: Flexible configuration management

## 📋 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/blog-post` | Create a new blog post |
| GET | `/api/blog-post` | Get all blog posts |
| GET | `/api/blog-post/:id` | Get a specific blog post |
| PATCH | `/api/blog-post/:id` | Update a blog post |
| DELETE | `/api/blog-post/:id` | Delete a blog post |
| GET | `/health` | Health check endpoint |

## 🏗️ Project Structure

```
BlogManagment/
├── internal/                 # Private application code
│   ├── config/              # Database and app configuration
│   ├── controller/          # HTTP handlers (API endpoints)
│   ├── middleware/          # HTTP middleware (logging, error handling)
│   ├── models/              # Data structures and DTOs
│   ├── repository/          # Data access layer
│   ├── routes/              # Route definitions
│   └── service/             # Business logic layer
├── docs/                    # API documentation
├── main.go                  # Application entry point
├── go.mod                   # Go module definition
├── config.env               # Environment variables
└── README.md                # Project documentation
```

## 🛠️ Technology Stack

- **Framework**: [Go Fiber](https://gofiber.io/) - Fast HTTP framework
- **Database**: PostgreSQL - Relational database
- **ORM**: GORM - Go ORM library
- **Testing**: Testify - Testing framework
- **Environment**: Godotenv - Environment variable management

## 📦 Installation

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher

### Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd BlogManagment
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up PostgreSQL database**
   ```sql
   CREATE DATABASE blog_management;
   ```

4. **Configure environment variables**
   ```bash
   cp config.env.example config.env
   # Edit config.env with your database credentials
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

## 🔧 Configuration

Update the `config.env` file with your database settings:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=blog_management
SERVER_PORT=8080
```

## 🧪 Testing

Run all tests with coverage:

```bash
go test ./... -cover
```

Run specific test files:

```bash
go test ./internal/service -v
go test ./internal/controller -v
```

## 📚 API Documentation

For detailed API documentation, see [docs/API_DOCUMENTATION.md](docs/API_DOCUMENTATION.md)

### Quick Examples

#### Create a blog post
```bash
curl -X POST http://localhost:8080/api/blog-post \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Blog Post",
    "description": "This is a brief description",
    "body": "This is the main content of my blog post..."
  }'
```

#### Get all blog posts
```bash
curl -X GET http://localhost:8080/api/blog-post
```

#### Get a specific blog post
```bash
curl -X GET http://localhost:8080/api/blog-post/{id}
```

#### Update a blog post
```bash
curl -X PATCH http://localhost:8080/api/blog-post/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title"
  }'
```

#### Delete a blog post
```bash
curl -X DELETE http://localhost:8080/api/blog-post/{id}
```

## 🏛️ Architecture

### Clean Architecture Layers

1. **Controller Layer** (`internal/controller/`)
   - Handles HTTP requests and responses
   - Input validation and error handling
   - Delegates business logic to service layer

2. **Service Layer** (`internal/service/`)
   - Contains business logic
   - Orchestrates between controller and repository
   - Handles data transformation and validation

3. **Repository Layer** (`internal/repository/`)
   - Data access abstraction
   - Database operations
   - Implements repository pattern for testability

4. **Model Layer** (`internal/models/`)
   - Data structures and DTOs
   - Database schema definitions
   - Request/Response models

### Design Patterns

- **Repository Pattern**: Abstracts data access layer
- **Dependency Injection**: Loose coupling between layers
- **Interface Segregation**: Clean interfaces for each layer
- **Single Responsibility**: Each component has one reason to change

## 🔒 Error Handling

The API provides consistent error responses:

```json
{
  "error": "Error title",
  "message": "Detailed error message"
}
```

Common HTTP status codes:
- `200` - Success
- `201` - Created
- `400` - Bad Request (validation errors)
- `404` - Not Found
- `500` - Internal Server Error

## 🚀 Deployment

### Docker (Optional)

Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

Build and run:

```bash
docker build -t blog-management-api .
docker run -p 8080:8080 blog-management-api
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For support and questions:
- Create an issue in the repository
- Check the API documentation
- Review the test examples

## 🔄 Version History

- **v1.0.0** - Initial release with CRUD operations
  - PostgreSQL integration
  - Clean architecture implementation
  - Comprehensive test coverage
  - API documentation 