# Jelastic Golang Hello World

A simple Go web application built with Fiber framework, designed for deployment on Jelastic cloud platform.

## Features

- HTTP server with Fiber framework
- JSON API responses
- Environment variable configuration
- Request logging middleware
- Query parameter support

## API Endpoints

### GET /

Returns a JSON message that can be customized via environment variables and query parameters.

**Response:**

```json
{
  "message": "Built with Love. Run with Ruk Com."
}
```

**Environment Variables:**

- `PORT`: Server port (default: 3000)
- `TEST_MSG`: Additional message to append

**Query Parameters:**

- `message`: Additional text to append to the response

## Local Development

1. Install Go 1.19 or later
2. Clone this repository
3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

5. Test the API:

   ```bash
   curl http://localhost:3000/
   curl "http://localhost:3000/?message=hello"
   ```

## Jelastic Deployment

### One-Click Deploy

Deploy this application instantly to Jelastic cloud with our JPS manifest:

[![Deploy to Jelastic](https://github.com/Ruk-Com-Cloud/simple-jps/blob/main/deploy-to-ruk-com.png?raw=true)](https://app.ruk-com.cloud/import-template?jps=https://raw.githubusercontent.com/Ruk-Com-Cloud/jelastic-golang-hello/main/manifest.jps)

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
