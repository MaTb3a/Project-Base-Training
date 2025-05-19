# WikiDocify

A modern document management system built with Go, featuring a RESTful API powered by Gin, GORM for database operations, and PostgreSQL for data storage.

## 📋 Features

- RESTful API endpoints for document management
- PostgreSQL database with GORM ORM
- Generic repository pattern implementation
- Docker containerization
- Health checks and graceful shutdown
- Environment-based configuration
- Structured logging
- API documentation with Swagger

## 🛠️ Technology Stack

- **Backend:** Go 1.24+
- **Web Framework:** Gin
- **ORM:** GORM
- **Database:** PostgreSQL 17
- **Containerization:** Docker & Docker Compose
- **Documentation:** Swagger/OpenAPI

## 📦 Prerequisites

- Docker (20.10+)
- Docker Compose (v2.0+)
- Go 1.24+ (optional, for local development)
- Make (optional, for using Makefile commands)

## 🚀 Getting Started

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/MaTb3aa/WikiDocify.git
   cd WikiDocify
   ```

2. **Start the application**
   ```bash
   docker compose up --build
   ```

3. **Verify the installation**
   ```bash
   curl http://localhost:8888/ping
   ```

### Local Development

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Run PostgreSQL**
   ```bash
   docker compose up db -d
   ```

4. **Start the application**
   ```bash
   go run main.go
   ```

## 🔄 API Endpoints

### Document Management

- `GET /documents` - List all documents
- `POST /documents` - Create a new document
- `GET /documents/:id` - Get document by ID
- `PUT /documents/:id` - Update document
- `DELETE /documents/:id` - Delete document

### Health Check

- `GET /ping` - Service health check

## 📝 Example Requests

### Create Document
```bash
curl -X POST http://localhost:8888/documents \
  -H "Content-Type: application/json" \
  -d '{"title":"Sample","content":"Content","author":"John Doe"}'
```

### List Documents
```bash
curl http://localhost:8888/documents
```

## 🛠️ Project Structure

```
WikiDocify/
├── config/         # Configuration management
├── handlers/       # HTTP request handlers
├── models/         # Database models
├── repository/     # Data access layer
│   ├── IGenericRepository.go
│   └── GenericRepository.go
├── routes/         # API route definitions
├── docker/         # Docker configuration files
├── docs/          # Documentation
└── main.go        # Application entry point
```

## 🔧 Configuration

Environment variables (can be set in `.env`):

| Variable | Description | Default |
|----------|-------------|---------|
| DB_HOST | Database host | localhost |
| DB_PORT | Database port | 5432 |
| DB_USER | Database user | postgres |
| DB_PASSWORD | Database password | postgres |
| DB_NAME | Database name | document_db |
| PORT | API server port | 8888 |

## 🏗️ Development

### Running Tests
```bash
go test ./... -v
```

### Generate Swagger Documentation
```bash
swag init
```

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📞 Support

For support, please open an issue in the GitHub repository.
