# Makefile for Feasto Project Management

# Build and start all services
build:
	docker-compose up --build

# Start all services (without rebuilding)
up:
	docker-compose up

# Stop all services
down:
	docker-compose down

# Stop services and DELETE ALL DATA (for a fresh start)
clean:
	docker-compose down -v

# Follow the logs of all running services
logs:
	docker-compose logs -f

help:
	@echo "Available commands:"
	@echo "  make build   - Build images and start all services"
	@echo "  make up      - Start all services"
	@echo "  make down    - Stop all services"
	@echo "  make clean   - Stop services and delete all data"
	@echo "  make logs    - Follow container logs"