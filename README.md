# Jelastic Golang Hello World

A Go web application built with Fiber framework and hexagonal architecture, featuring a PostgreSQL database and RESTful API. Designed for deployment on Jelastic cloud platform.

## Features

- HTTP server with Fiber framework
- Hexagonal architecture (ports and adapters)
- PostgreSQL database with GORM
- RESTful API for user management
- Database seeding system with sample data
- Environment variable configuration
- Request logging middleware
- Docker Compose for local development
- Makefile for easy development workflow

## API Endpoints

### Health Check

- `GET /` - Returns a JSON health check message

### User Management

- `POST /users` - Create a new user
- `GET /users` - Get all users
- `GET /users/:id` - Get user by ID
- `PUT /users/:id` - Update user by ID
- `DELETE /users/:id` - Delete user by ID

**User JSON Structure:**

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com"
}
```

**Environment Variables:**

- `PORT`: Server port (default: 3000)
- `TEST_MSG`: Additional message to append to health check
- `DB_HOST`: PostgreSQL host (default: localhost)
- `DB_USER`: Database user (default: postgres)
- `DB_PASSWORD`: Database password (default: password)
- `DB_NAME`: Database name (default: testdb)
- `DB_PORT`: Database port (default: 5432)
- `DB_SSLMODE`: SSL mode (default: disable)

## Local Development

### Prerequisites

- Go 1.20 or later
- Docker and Docker Compose

### Quick Start

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd jelastic-golang-hello
   ```

2. **Start PostgreSQL with Docker Compose**

   ```bash
   docker-compose up -d
   ```

3. **Set up environment variables**

   ```bash
   cp .env.example .env
   # Edit .env if needed
   ```

4. **Install dependencies**

   ```bash
   go mod tidy
   ```

5. **Run the application**

   ```bash
   go run main.go
   ```

6. **Test the API**

   ```bash
   # Health check
   curl http://localhost:3000/

   # Create a user
   curl -X POST http://localhost:3000/users \
     -H "Content-Type: application/json" \
     -d '{"name":"John Doe","email":"john@example.com"}'

   # Get all users
   curl http://localhost:3000/users

   # Get user by ID
   curl http://localhost:3000/users/1

   # Update user
   curl -X PUT http://localhost:3000/users/1 \
     -H "Content-Type: application/json" \
     -d '{"name":"Jane Doe","email":"jane@example.com"}'

   # Delete user
   curl -X DELETE http://localhost:3000/users/1
   ```

### Database Management

**Start PostgreSQL:**

```bash
docker-compose up -d
```

**Stop PostgreSQL:**

```bash
docker-compose down
```

**View PostgreSQL logs:**

```bash
docker-compose logs postgres
```

**Connect to PostgreSQL:**

```bash
docker exec -it jelastic-postgres psql -U postgres -d testdb
```

**Reset database (remove volume):**

```bash
docker-compose down -v
docker-compose up -d
```

### Database Seeding

The application includes a comprehensive seeding system to populate the database with sample data.

**Using Make commands (recommended):**

```bash
# Run all seeders
make seed

# Run only user seeder
make seed-users

# Rollback all seeders
make rollback

# List available seeders
make list-seeders

# Setup development environment (starts DB + runs seeders)
make dev-setup

# Reset development environment (resets DB + runs seeders)
make dev-reset
```

**Using Go commands directly:**

```bash
# Run all seeders
go run cmd/seeder/main.go -action=seed

# Run specific seeder
go run cmd/seeder/main.go -action=seed -seeder=UserSeeder

# Rollback all seeders
go run cmd/seeder/main.go -action=rollback

# List available seeders
go run cmd/seeder/main.go -action=list

# Show help
go run cmd/seeder/main.go -help
```

**Using application flags:**

```bash
# Run application with seeders
go run main.go -seed

# Run seeders only (don't start server)
go run main.go -seed-only
```

**Sample Data:**
The user seeder creates 10 sample users with realistic names and email addresses:

- John Doe (<john.doe@example.com>)
- Jane Smith (<jane.smith@example.com>)
- Bob Johnson (<bob.johnson@example.com>)
- And 7 more...

**Creating Custom Seeders:**

1. Create a new seeder in `internal/seeder/`
2. Implement the `Seeder` interface
3. Register it in `internal/seeder/registry.go`

## Jelastic Deployment

### One-Click Deploy

Deploy this application instantly to Jelastic cloud with our JPS manifest:

[![Deploy to Jelastic](https://github.com/Ruk-Com-Cloud/simple-jps/blob/main/deploy-to-ruk-com.png?raw=true)](https://app.manage.ruk-com.cloud/?jps=https://raw.githubusercontent.com/Ruk-Com-Cloud/jelastic-golang-hello/main/manifest.jps)

**Or manually import:**

1. Go to [Jelastic Import](https://app.ruk-com.cloud/import-template)
2. Use this URL: `https://raw.githubusercontent.com/Ruk-Com-Cloud/jelastic-golang-hello/main/manifest.jps`
3. Click "Import" and follow the installation wizard

### Prerequisites

- Jelastic account
- Access to Jelastic dashboard

### Deployment Steps

1. **Create New Environment**
   - Log into your Jelastic dashboard
   - Click "New Environment"
   - Select "Go" as the programming language
   - Choose Go version 1.19 or later
   - Set environment name (e.g., "golang-hello")
   - Configure topology as needed
   - Click "Create"

2. **Deploy Application**
   - Option A: Git deployment (recommended)
     - In Jelastic dashboard, go to your environment
     - Click "Deployment Manager"
     - Add your Git repository URL
     - Click "Deploy to..."
     - Select your Go application server

   - Option B: Archive upload
     - Create a ZIP archive of your project
     - Upload via Jelastic deployment manager
     - Deploy to your Go server

3. **Configure Environment Variables** (Optional)
   - Go to environment Settings
   - Add environment variables:
     - `PORT`: Will be automatically set by Jelastic
     - `TEST_MSG`: Your custom message

4. **Access Your Application**
   - Once deployed, Jelastic will provide a URL
   - Test your endpoints:
     - `https://your-env-name.app.ruk-com.cloud/`
     - `https://your-env-name.app.ruk-com.cloud/?message=test`

### Build Configuration

The application uses standard Go build process:

```bash
go build -o main .
```

For Jelastic deployment, ensure your `go.mod` file is properly configured with Go 1.19+ and all dependencies are listed.

### Environment-Specific Notes

- Jelastic automatically sets the `PORT` environment variable
- The application listens on all interfaces (0.0.0.0)
- Logs are automatically collected by Jelastic monitoring
- SSL certificates are managed by Jelastic platform

## License

MIT License - see [LICENSE](LICENSE) file for details.
