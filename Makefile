# Variables
BINARY_NAME=stress-test
DOCKER_IMAGE_NAME=goexpert-httploadtester
DOCKER_COMPOSE_FILE=docker-compose.yaml

# Build the Go binary
build:
	@echo "🚧 Building the Go binary..."
	go build -o $(BINARY_NAME) .
	@echo "✅ Build completed."

# Run the binary locally with default parameters
run:
	@echo "🚀 Running the binary locally..."
	./$(BINARY_NAME) --help

# Clean up the binary
clean:
	@echo "🧹 Cleaning up the build files..."
	rm -f $(BINARY_NAME)
	@echo "✅ Cleanup completed."

# Build and run Docker Compose
up:
	@echo "🐳 Building and starting Docker Compose..."
	docker compose -f $(DOCKER_COMPOSE_FILE) up --build
	@echo "✅ Docker Compose is up and running."

# Stop Docker Compose
down:
	@echo "🛑 Stopping Docker Compose..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down
	@echo "✅ Docker Compose stopped."

# Restart Docker Compose
restart:
	@echo "🔄 Restarting Docker Compose..."
	$(MAKE) down
	$(MAKE) up
	@echo "✅ Docker Compose restarted."

# Help command to display all available targets
help:
	@echo "📜 Makefile Help:"
	@echo "  build          -> Build the Go binary"
	@echo "  run            -> Run the binary locally"
	@echo "  clean          -> Remove the binary"
	@echo "  docker-up      -> Build and start services using Docker Compose"
	@echo "  docker-down    -> Stop all services using Docker Compose"
	@echo "  docker-restart -> Restart all services using Docker Compose"
	@echo "  help           -> Display this help message"

.PHONY: build run clean docker-up docker-down docker-restart help
