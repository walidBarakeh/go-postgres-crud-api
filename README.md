// README.md
# Go PostgreSQL CRUD API with Migrations

A complete Go REST API with PostgreSQL, schema migrations, and CRUD operations using GORM and Gin.

[![CI/CD Pipeline](https://github.com/yourusername/go-postgres-crud-api/actions/workflows/ci.yml/badge.svg)](https://github.com/yourusername/go-postgres-crud-api/actions/workflows/ci.yml)

## Features

- REST API with CRUD operations for Users and Posts
- PostgreSQL database with GORM ORM
- Database schema migrations using golang-migrate
- Docker setup with PostgreSQL
- GitHub Actions CI/CD pipeline
- Structured project layout
- Environment configuration
- Foreign key relationships
- Input validation
- Error handling

## Quick Start with Docker

The fastest way to get started is using Docker:

```bash
# Clone the repository
git clone https://github.com/yourusername/go-postgres-crud-api.git
cd go-postgres-crud-api

# Start PostgreSQL with Docker
make docker-up

# Run migrations and start the API
make setup
```

That's it! The API will be available at `http://localhost:8080`.

## Manual Setup

If you prefer to install PostgreSQL manually:

1. **Install dependencies:**
```bash
go mod tidy
make install-migrate
```

2. **Setup PostgreSQL database:**
```bash
createdb crud_db
# Or using make:
make setup-db
```

3. **Configure environment:**
```bash
cp .env.example .env
# Update database credentials if needed
```

4. **Run migrations:**
```bash
make migrate-up
```

5. **Start the server:**
```bash
make run
```

## Docker Options

### Option 1: PostgreSQL Only (Recommended for Development)
```bash
# Start PostgreSQL container
make docker-up

# Run API locally
make run

# View PostgreSQL logs
make docker-logs

# Stop PostgreSQL
make docker-down
```

### Option 2: Full Docker Setup (API + PostgreSQL)
```bash
# Build and run everything in Docker
docker-compose -f docker-compose.full.yml up --build

# With pgAdmin for database management
docker-compose -f docker-compose.full.yml --profile tools up
```

### Option 3: pgAdmin Database Management
Access pgAdmin at `http://localhost:5050`:
- Email: `admin@example.com`
- Password: `admin`

Add server connection:
- Host: `postgres` (or `localhost` if running locally)
- Port: `5432`
- Username: `postgres`
- Password: `postgres`

## API Endpoints

### Users
- `GET /api/v1/users` - Get all users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user
- `GET /api/v1/users/:userId/posts` - Get user's posts

### Posts
- `GET /api/v1/posts` - Get all posts
- `POST /api/v1/posts` - Create post
- `GET /api/v1/posts/:id` - Get post by ID
- `PUT /api/v1/posts/:id` - Update post
- `DELETE /api/v1/posts/:id` - Delete post

## Migration Commands

```bash
# Run all migrations up
make migrate-up

# Run all migrations down
make migrate-down

# Run specific number of migrations
make migrate-up-steps steps=1
make migrate-down-steps steps=1

# Check current version
make migrate-version

# Force to specific version (if stuck)
make migrate-force version=1

# Create new migration
make migrate-create name=add_users_avatar
```

## Example API Usage

### Create User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","age":30}'
```

### Create Post
```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"My First Post","content":"Hello World!","user_id":1,"published":true}'
```

### Get All Users
```bash
curl http://localhost:8080/api/v1/users
```

## Project Structure

```
├── .github/workflows/  # GitHub Actions CI/CD
├── cmd/migrate/        # Migration CLI tool
├── internal/
│   ├── database/      # Database connection
│   ├── handlers/      # HTTP handlers
│   ├── models/        # Data models
│   └── routes/        # Route definitions
├── migrations/        # SQL migration files
├── bin/               # Compiled binaries (gitignored)
├── docker-compose.yml # PostgreSQL container
├── docker-compose.full.yml # Full stack with API
├── Dockerfile         # API container image
├── Makefile          # Build and migration commands
├── .env.example      # Environment template
├── .env              # Environment configuration (gitignored)
├── .gitignore        # Git ignore rules
└── main.go           # Application entry point
```

## GitHub Setup

1. **Create a new repository on GitHub**

2. **Initialize and push your code:**
```bash
# Initialize git repository
git init

# Add all files
git add .

# Create initial commit
git commit -m "Initial commit: Go PostgreSQL CRUD API with Docker"

# Add GitHub remote (replace with your repository URL)
git remote add origin https://github.com/yourusername/go-postgres-crud-api.git

# Push to GitHub
git push -u origin main
```

3. **GitHub Actions will automatically:**
   - Run tests on every push/PR
   - Set up PostgreSQL for testing
   - Build the application
   - Build Docker image on main branch

## Docker Commands

```bash
# Start PostgreSQL container
make docker-up

# Stop PostgreSQL container
make docker-down

# View PostgreSQL logs
make docker-logs

# Full setup (start PostgreSQL + run migrations)
make setup

# Build Docker image locally
docker build -t crud-api .

# Run full stack with Docker Compose
docker-compose -f docker-compose.full.yml up --build
```

## Environment Variables

The application uses these environment variables (see `.env.example`):

```env
DB_HOST=localhost      # Database host
DB_PORT=5432          # Database port
DB_USER=postgres      # Database user
DB_PASSWORD=postgres  # Database password
DB_NAME=crud_db       # Database name
DB_SSL_MODE=disable   # SSL mode
PORT=8080            # API server port
```

## Development Workflow

1. **Start development environment:**
```bash
make docker-up    # Start PostgreSQL
make migrate-up   # Apply migrations
make run         # Start API server
```

2. **Create new migration:**
```bash
make migrate-create name=add_user_avatar
# Edit the generated files in migrations/
make migrate-up  # Apply the migration
```

3. **Reset database:**
```bash
make migrate-down  # Rollback all migrations
make migrate-up    # Reapply migrations
```

## Production Deployment

The project includes a `Dockerfile` for production deployment:

```bash
# Build production image
docker build -t crud-api:latest .

# Run with external PostgreSQL
docker run -p 8080:8080 \
  -e DB_HOST=your-postgres-host \
  -e DB_PASSWORD=your-password \
  crud-api:latest
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes
4. Run tests: `go test ./...`
5. Commit changes: `git commit -am 'Add feature'`
6. Push to branch: `git push origin feature-name`
7. Create a Pull Request

## License

This project is licensed under the MIT License.