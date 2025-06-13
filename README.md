# Expense API

A Go-based expense tracking API with PostgreSQL, structured logging, and Docker support.

## Getting Started

1. Clone this repository:

```bash
git clone https://github.com/your-username/expense-api.git
cd expense-api
```

2. Start the application using Docker Compose:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`.

## API Endpoints

- `GET /health` - Health check endpoint
- `GET /expense` - Expense API endpoint

## Features

- Structured logging with `log/slog`
- PostgreSQL database with Docker Compose
- Configuration management with environment variables
- HTTP/REST API with Chi router
- Testing setup with Go's testing package
- Docker containerization for both API and database

## Prerequisites

- Docker and Docker Compose
- Go 1.21 or later (for local development)

## Environment Variables

The following environment variables are configured in the Docker Compose file:

- `PORT` - Server port (default: "8080")
- `DB_HOST` - Database host (default: "postgres")
- `DB_PORT` - Database port (default: "5432")
- `DB_NAME` - Database name (default: "expense-api")
- `DB_USER` - Database user (default: "postgres")
- `DB_PASSWORD` - Database password (default: "postgres")
- `DB_SSLMODE` - Database SSL mode (default: "disable")

## Development

### Local Development

1. Start the database:

```bash
docker-compose up -d postgres
```

2. Run the application locally:

```bash
export DB_HOST=localhost DB_PORT=5432 DB_USER=postgres DB_PASSWORD=postgres DB_NAME=expense-api DB_SSLMODE=disable
make run
```

### Testing

Run the tests with:

```bash
make test
```

## Project Structure

```
.
├── config/           # Configuration management
├── internal/         # Internal packages
│   ├── database/    # Database connection and queries
│   └── handlers/    # HTTP handlers and routing
├── main.go          # Application entry point
├── docker-compose.yml # Docker Compose configuration
├── Dockerfile       # Go application container configuration
├── Makefile         # Build and run commands
└── go.mod           # Go module file
```

## Docker Setup

The project uses two containers:

1. `postgres:15-alpine` - PostgreSQL database
2. Custom Go application container

The containers are configured to:

- Run in the same Docker network
- Use persistent volume for database storage
- Expose necessary ports (8080 for API, 5432 for database)
- Include health checks for the database
