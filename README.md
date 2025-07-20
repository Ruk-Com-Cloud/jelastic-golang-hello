# Jelastic Golang Hello World (HTTP-Only)

A simple Go web application built with Fiber framework for standalone HTTP-only deployment. This version has been simplified to run without any database or Redis dependencies.

## Features

- Lightweight HTTP server with Fiber framework
- Simple configuration management with Viper
- Environment variable support
- Request logging middleware
- RESTful API endpoints
- Designed for Jelastic cloud platform deployment

## API Endpoints

### Health Check

- `GET /` - Returns a JSON health check message with optional query parameters
- `GET /api/info` - Returns service information
- `GET /api/health` - Returns health status with uptime
- `GET /api/echo` - Echo service with message parameter

**Example Responses:**

Health check:

```json
{
  "message": "Built with Love. Run with Ruk Com.",
  "environment": "development",
  "version": "1.0.0",
  "mode": "http-only"
}
```

Service info:

```json
{
  "service": "jelastic-golang-hello",
  "version": "1.0.0",
  "timestamp": "2025-07-20T10:30:00Z",
  "environment": "http-only"
}
```

## Configuration

**Environment Variables:**

- `PORT`: Server port (default: 3000)
- `HOST`: Server host (default: 0.0.0.0)
- `TEST_MSG`: Additional message to append to health check
- `ENVIRONMENT`: Environment name (default: development)

## Local Development

### Prerequisites

- Go 1.20 or later
- Docker and Docker Compose (for containerized deployment)

### Quick Start

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd jelastic-golang-hello
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Run the application**

   ```bash
   make run
   # or
   go run main.go
   ```

4. **Test the API**

   ```bash
   # Health check
   curl http://localhost:3000/

   # Service info
   curl http://localhost:3000/api/info

   # Health status
   curl http://localhost:3000/api/health

   # Echo service
   curl "http://localhost:3000/api/echo?message=Hello%20World"
   ```

### Development Commands

**Build and Run:**

```bash
make build    # Build the application
make run      # Run the application
make clean    # Clean build artifacts
```

**Code Quality:**

```bash
make test     # Run tests
make fmt      # Format code
make vet      # Vet code
make tidy     # Tidy modules
```

**Docker Commands:**
```bash
make docker-build  # Build Docker image
make docker-run    # Run with Docker Compose
make docker-stop   # Stop Docker containers
make docker-logs   # View container logs
make docker-clean  # Clean Docker resources
make docker-shell  # Access container shell
```

## Docker Deployment

### Using Docker Compose (Recommended)

1. **Build and run with Docker Compose**
   ```bash
   make docker-run
   # or
   docker-compose up -d
   ```

2. **View logs**
   ```bash
   make docker-logs
   # or
   docker-compose logs -f
   ```

3. **Stop the application**
   ```bash
   make docker-stop
   # or
   docker-compose down
   ```

### Using Docker directly

1. **Build the Docker image**
   ```bash
   make docker-build
   # or
   docker build -t jelastic-golang-hello:latest .
   ```

2. **Run the container**
   ```bash
   docker run -d \
     --name jelastic-golang-hello \
     -p 3000:3000 \
     -e TEST_MSG="Running in Docker!" \
     jelastic-golang-hello:latest
   ```

3. **View logs**
   ```bash
   docker logs -f jelastic-golang-hello
   ```

4. **Stop and remove**
   ```bash
   docker stop jelastic-golang-hello
   docker rm jelastic-golang-hello
   ```

### Docker Environment Variables

The Docker setup supports all the same environment variables:
- `PORT=3000` - Container port
- `HOST=0.0.0.0` - Bind address
- `TEST_MSG` - Custom message
- `ENVIRONMENT` - Environment name

## Jelastic Deployment

### One-Click Deploy

Deploy this application instantly to Jelastic cloud with our JPS manifest:

[![Deploy to Jelastic](https://github.com/Ruk-Com-Cloud/simple-jps/blob/main/deploy-to-ruk-com.png?raw=true)](https://app.manage.ruk-com.cloud/?jps=https://raw.githubusercontent.com/Ruk-Com-Cloud/jelastic-golang-hello/http-only/manifest.jps)

**Or manually import:**

1. Go to [Jelastic Import](https://app.ruk-com.cloud/import-template)
2. Use this URL: `https://raw.githubusercontent.com/Ruk-Com-Cloud/jelastic-golang-hello/http-only/manifest.jps`
3. Click "Import" and follow the installation wizard

### Build Configuration

The application uses standard Go build process:

```bash
go build -o main .
```

For Jelastic deployment, ensure your `go.mod` file is properly configured with Go 1.20+ and all dependencies are listed.

### Environment-Specific Notes

- Jelastic automatically sets the `PORT` environment variable
- The application listens on all interfaces (0.0.0.0)
- Logs are automatically collected by Jelastic monitoring
- SSL certificates are managed by Jelastic platform
- No database or external dependencies required

## Architecture

This is a simplified HTTP-only application with the following structure:

```text
internal/
├── config/     # Configuration management
└── handlers/   # HTTP handlers and routing
```

- **Web Framework**: Fiber v2 (Express-like framework for Go)
- **Configuration**: Viper for environment/config management
- **Architecture**: Simple HTTP service without external dependencies

## License

MIT License - see [LICENSE](LICENSE) file for details.
