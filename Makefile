.PHONY: run build test clean docker-up docker-down

# Binary name
BINARY_NAME=<REPLACE>

# Build the application
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run:
	go run main.go

# Run tests
test:
	go test -v ./...

# Clean build files
clean:
	go clean
	rm -f $(BINARY_NAME)

# Start docker containers
up:
	docker-compose up -d

# Stop docker containers
down:
	docker-compose down