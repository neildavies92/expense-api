# Go Microservice Template

A template repository for creating Go microservices with PostgreSQL, structured logging, and configuration management.

## Getting Started

1. Clone this repository:
```bash
git clone https://github.com/bryx/go-microservice-template.git <your-service-name>
cd <your-service-name>
```

2. Run the setup script to replace all instances of `<REPLACE>` with your service name:
```bash
./bin/setup.sh --service-name your-service-name
```

The script will:
- Validate the service name format (lowercase letters, numbers, and hyphens only)
- Replace all instances of `<REPLACE>` in the codebase
- Update module paths to `github.com/bryx/your-service-name`
- Update binary names and database names

3. Start the PostgreSQL database:
```bash
make up
```

4. Install dependencies:
```bash
go mod tidy
```

5. Run the application:
```bash
make run
```

The server will start on port 8080 by default. You can change this by setting the `PORT` environment variable.

## Features

- Structured logging with `log/slog`
- PostgreSQL database with Docker Compose
- Configuration management with environment variables
- HTTP/REST API with Chi router
- Testing setup with Go's testing package

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- Make (optional, for using Makefile commands)

## Environment Variables

- `PORT` - Server port (default: "8080")
- `DB_HOST` - Database host (default: "localhost")
- `DB_PORT` - Database port (default: "5432")
- `DB_USER` - Database user (default: "postgres")
- `DB_PASSWORD` - Database password (default: "postgres")
- `DB_NAME` - Database name (default: "microservice")
- `DB_SSLMODE` - Database SSL mode (default: "disable")

## Testing

Run the tests with:
```bash
make test
```

## Project Structure

```
.
├── bin/              # Setup and utility scripts
├── config/           # Configuration management
├── internal/         # Internal packages
│   ├── database/    # Database connection and queries
│   └── handlers/    # HTTP handlers and routing
├── main.go          # Application entry point
├── docker-compose.yml # Docker Compose configuration
├── Makefile         # Build and run commands
└── go.mod           # Go module file
```
