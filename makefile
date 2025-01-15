# Run the application
run:
	go mod tidy
	go run cmd/wascherei-go/main.go

# Build the application binary
build:
	go build -o bin/wascherei-go ./cmd/wascherei-go

# Run migrations
migrate:
	go run cmd/migrate/migrate.go

# Clean the build (remove binaries and build artifacts)
clean:
	rm -f bin/wascherei-go

# Change environment to production
env-prod:
	rm .env
	cp .env.prod .env

# Change environment to development
env-dev:
	rm .env
	cp .env.dev .env