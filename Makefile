.PHONY: run build migrate-up migrate-down migrate-create migrate-force migrate-version

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o bin/api main.go
	go build -o bin/migrate cmd/migrate/main.go

# Run migrations up
migrate-up:
	go run cmd/migrate/main.go -action=up

# Run migrations down
migrate-down:
	go run cmd/migrate/main.go -action=down

# Run specific number of migrations up
migrate-up-steps:
	go run cmd/migrate/main.go -action=up -steps=$(steps)

# Run specific number of migrations down
migrate-down-steps:
	go run cmd/migrate/main.go -action=down -steps=$(steps)

# Create a new migration file
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

# Force migration to specific version
migrate-force:
	go run cmd/migrate/main.go -action=force -version=$(version)

# Check migration version
migrate-version:
	go run cmd/migrate/main.go -action=version

# Install dependencies
deps:
	go mod tidy
	go mod download

# Install golang-migrate CLI
install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Setup database (create database)
setup-db:
	createdb crud_db

# Drop database
drop-db:
	dropdb crud_db

# Docker commands
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f postgres

# Full setup with Docker
setup: docker-up
	sleep 5
	make migrate-up
