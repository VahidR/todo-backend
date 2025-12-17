# Todo Backend

**Status:** Active Development  
**Author:** @VahidR  
**Last Updated:** December 2025

---

## Overview

A production-ready RESTful API backend for managing todo items, built with Go. This service provides CRUD operations for todos with a clean architecture pattern separating concerns across handlers, services, and repositories.

### Goals

- Provide a simple, reliable API for todo management
- Follow clean architecture principles for maintainability
- Support horizontal scaling with stateless design
- Enable easy local development and deployment

### Non-Goals

- User authentication (to be implemented separately)
- Real-time updates via WebSockets
- Multi-tenant support

---

## Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         HTTP Layer                               │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                    Gin Router                            │    │
│  │    /api/todos/* ──► CORS ──► Logger ──► Recovery        │    │
│  └─────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                       Handler Layer                              │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                   todo.Handler                           │    │
│  │  • Request validation    • Response formatting           │    │
│  │  • HTTP status codes     • Error handling                │    │
│  └─────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                       Service Layer                              │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                   todo.Service                           │    │
│  │  • Business logic        • Input validation              │    │
│  │  • Domain errors         • Context propagation           │    │
│  └─────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Repository Layer                            │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                  todo.Repository                         │    │
│  │  • Data persistence      • GORM operations               │    │
│  │  • Query building        • Auto-migration                │    │
│  └─────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                         MySQL                                    │
└─────────────────────────────────────────────────────────────────┘
```

### Project Structure

```
todo-backend/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go         # Environment configuration
│   ├── database/
│   │   └── database.go       # Database connection & migrations
│   ├── router/
│   │   └── router.go         # HTTP router setup & middleware
│   └── todo/
│       ├── handler.go        # HTTP request handlers
│       ├── model.go          # Domain model (Todo struct)
│       ├── repository.go     # Data access layer
│       └── service.go        # Business logic layer
├── bin/                      # Compiled binaries
├── Makefile                  # Build automation
├── go.mod                    # Go module definition
└── go.sum                    # Dependency checksums
```

---

## Getting Started

### Prerequisites

| Dependency | Version | Purpose |
|------------|---------|---------|
| Go | 1.25+ | Runtime |
| MySQL | 8.0+ | Database |
| Make | 4.0+ | Build automation |

### Quick Start

```bash
# 1. Clone the repository
git clone https://github.com/VahidR/todo-backend.git
cd todo-backend

# 2. Set up environment variables
export DB_DSN="user:password@tcp(localhost:3306)/todo_db?parseTime=true&charset=utf8mb4&loc=Local"
export PORT="8080"
export ENV="development"

# 3. Download dependencies
make deps

# 4. Run the server
make run
```

### Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `DB_DSN` | ✅ | - | MySQL connection string |
| `PORT` | ❌ | `8080` | Server port |
| `ENV` | ❌ | `development` | Environment (`development`, `production`) |

**Example DB_DSN format:**
```
user:password@tcp(localhost:3306)/todo_db?parseTime=true&charset=utf8mb4&loc=Local
```

---

## API Reference

### Base URL

```
http://localhost:8080/api
```

### Endpoints

#### List All Todos

```http
GET /api/todos/
```

**Response:** `200 OK`
```json
[
  {
    "id": 1,
    "title": "Buy groceries",
    "completed": false,
    "created_at": "2025-12-16T10:00:00Z",
    "updated_at": "2025-12-16T10:00:00Z"
  }
]
```

---

#### Get Todo by ID

```http
GET /api/todos/:id
```

**Response:** `200 OK`
```json
{
  "id": 1,
  "title": "Buy groceries",
  "completed": false,
  "created_at": "2025-12-16T10:00:00Z",
  "updated_at": "2025-12-16T10:00:00Z"
}
```

**Errors:**
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - Todo not found

---

#### Create Todo

```http
POST /api/todos/
Content-Type: application/json
```

**Request Body:**
```json
{
  "title": "Buy groceries"
}
```

**Response:** `201 Created`
```json
{
  "id": 1,
  "title": "Buy groceries",
  "completed": false,
  "created_at": "2025-12-16T10:00:00Z",
  "updated_at": "2025-12-16T10:00:00Z"
}
```

**Errors:**
- `400 Bad Request` - Title is required

---

#### Update Todo

```http
PUT /api/todos/:id
Content-Type: application/json
```

**Request Body:**
```json
{
  "title": "Buy groceries and cook dinner",
  "completed": true
}
```

**Response:** `200 OK`
```json
{
  "id": 1,
  "title": "Buy groceries and cook dinner",
  "completed": true,
  "created_at": "2025-12-16T10:00:00Z",
  "updated_at": "2025-12-16T12:00:00Z"
}
```

**Errors:**
- `400 Bad Request` - Invalid ID or missing title
- `404 Not Found` - Todo not found

---

#### Delete Todo

```http
DELETE /api/todos/:id
```

**Response:** `204 No Content`

**Errors:**
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - Todo not found

---

### Error Response Format

All errors follow this structure:

```json
{
  "error": "error message here"
}
```

---

## Development

### Available Make Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the binary to `bin/todo` |
| `make run` | Run the application |
| `make dev` | Run with hot reload (requires [air](https://github.com/air-verse/air)) |
| `make test` | Run all tests |
| `make test-coverage` | Run tests with coverage report |
| `make fmt` | Format code |
| `make vet` | Vet code for issues |
| `make lint` | Run golangci-lint |
| `make tidy` | Tidy and verify dependencies |
| `make deps` | Download dependencies |
| `make clean` | Clean build artifacts |
| `make build-all` | Cross-compile for Linux, macOS, Windows |
| `make help` | Show all available commands |

### Running Tests

```bash
# Run all tests
make test

# Run with coverage report
make test-coverage
# Opens coverage.html in browser
```

### Code Quality

```bash
# Format code
make fmt

# Run static analysis
make vet

# Run linter (install: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
make lint
```

### Hot Reload Development

```bash
# Install air (first time only)
go install github.com/air-verse/air@latest

# Run with hot reload
make dev
```

---

## Database

### Schema

The `todos` table is auto-migrated on startup:

| Column | Type | Constraints |
|--------|------|-------------|
| `id` | `BIGINT UNSIGNED` | PRIMARY KEY, AUTO_INCREMENT |
| `title` | `VARCHAR(255)` | NOT NULL |
| `completed` | `TINYINT(1)` | NOT NULL, DEFAULT 0 |
| `created_at` | `DATETIME(3)` | |
| `updated_at` | `DATETIME(3)` | |

### Setting Up MySQL

```bash
# Create database
mysql -u root -p -e "CREATE DATABASE todo_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# Create user (optional)
mysql -u root -p -e "CREATE USER 'todouser'@'localhost' IDENTIFIED BY 'password';"
mysql -u root -p -e "GRANT ALL PRIVILEGES ON todo_db.* TO 'todouser'@'localhost';"
mysql -u root -p -e "FLUSH PRIVILEGES;"
```

---

## Deployment

### Build Binary

```bash
# Build for current platform
make build

# Cross-compile for all platforms
make build-all
```

### Docker

```bash
# Build image
make docker-build

# Run container
make docker-run
```

> **Note:** Create a `Dockerfile` and `.env` file for Docker deployment.

### Production Checklist

- [ ] Set `ENV=production`
- [ ] Use secure MySQL credentials
- [ ] Configure proper CORS origins in `router.go`
- [ ] Set up health check endpoint
- [ ] Configure logging for production
- [ ] Set up reverse proxy (nginx/traefik)
- [ ] Enable TLS/HTTPS

---

## Tech Stack

| Component | Technology |
|-----------|------------|
| Language | Go 1.25 |
| Web Framework | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io/) |
| Database | MySQL 8.0 |
| CORS | [gin-contrib/cors](https://github.com/gin-contrib/cors) |

---

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Run `make fmt` before committing
- Ensure `make vet` and `make lint` pass
- Write tests for new features

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Appendix

### Related Documents

- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Gin Documentation](https://gin-gonic.com/en/docs/)
- [GORM Documentation](https://gorm.io/docs/)

### Changelog

| Version | Date | Changes |
|---------|------|---------|
| 0.1.0 | Dec 2025 | Initial release with CRUD operations |

