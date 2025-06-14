.PHONY: build run clean test

# Build the Docker services
build:
	docker-compose build

# Run the Docker services
run:
	docker-compose up

# Clean up Docker resources
clean:
	docker-compose down

# Run tests
test:
	go test ./... 