# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the applications
RUN go build -o bin/api main.go
RUN go build -o bin/migrate cmd/migrate/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Copy binaries from builder
COPY --from=builder /app/bin/ ./bin/
COPY --from=builder /app/migrations/ ./migrations/
COPY --from=builder /app/.env.docker ./.env

# Expose port
EXPOSE 8080

# Run the application
CMD ["./bin/api"]
