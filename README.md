# Go S2I Application for OpenShift

This is a simple Go application designed to be built and deployed using OpenShift's Source-to-Image (S2I) with the `s2i-go` ClusterTask in OpenShift Pipelines.

## 📁 Repository Structure

```
.
├── main.go              # Main application entry point
├── go.mod               # Go module definition
├── go.sum               # Go dependencies checksums
├── .gitignore           # Git ignore file
├── .s2i/
│   └── environment      # S2I build configuration
└── README.md            # This file
```

## 🚀 Quick Start

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/go-s2i-app.git
   cd go-s2i-app
   ```

2. **Run locally**
   ```bash
   go run main.go
   ```

3. **Build locally**
   ```bash
   go build -o app main.go
   ./app
   ```

4. **Test the application**
   ```bash
   # In another terminal
   curl http://localhost:8080
   curl http://localhost:8080/health
   curl http://localhost:8080/api/info
   ```

### Deploy to OpenShift using S2I Pipeline

1. **Create the pipeline** (see deployment guide)
   ```bash
   oc apply -f s2i-go-pipeline.yaml
   ```

2. **Run the pipeline**
   ```bash
   oc create -f s2i-go-pipeline.yaml
   ```

3. **Monitor the pipeline**
   ```bash
   tkn pipelinerun logs -f -L
   ```

## 🌐 API Endpoints

| Endpoint | Method | Description | Response Type |
|----------|--------|-------------|---------------|
| `/` | GET | Home page with HTML interface | HTML |
| `/health` | GET | Health check endpoint | JSON |
| `/ready` | GET | Readiness probe endpoint | JSON |
| `/api/info` | GET | Application information | JSON |

### Example Responses

**Health Check (`/health`)**
```json
{
  "status": "healthy",
  "timestamp": "2024-12-24T10:30:00Z",
  "version": "1.0.0",
  "uptime": "2h30m15s"
}
```

**Application Info (`/api/info`)**
```json
{
  "message": "Go application built with OpenShift S2I",
  "application": "s2i-go-demo",
  "environment": "production"
}
```

## 🔧 Configuration

### Environment Variables

The application supports the following environment variables:

- `PORT` - Server port (default: 8080)
- `ENVIRONMENT` - Application environment (default: development)

### S2I Build Configuration

The `.s2i/environment` file contains build-time configuration:

```bash
GO111MODULE=on
ENVIRONMENT=production
```

## 📦 Adding Dependencies

If you need to add external dependencies:

1. **Add the dependency**
   ```bash
   go get github.com/gin-gonic/gin
   ```

2. **Update your code** to use the dependency

3. **Commit changes**
   ```bash
   git add go.mod go.sum
   git commit -m "Add Gin dependency"
   git push
   ```

The S2I build will automatically download and include the dependencies.

## 🐳 S2I Build Process

The S2I builder performs these steps:

1. Clone the Git repository
2. Detect Go application (looks for `go.mod`)
3. Download dependencies (`go mod download`)
4. Build the application (`go build`)
5. Create container image with the binary
6. Set the default command to run the application

## 🔍 Troubleshooting

### Build Fails

**Issue: Module errors**
```bash
# Ensure go.mod is correctly formatted
go mod tidy
```

**Issue: Dependency download fails**
```bash
# Clear module cache and retry
go clean -modcache
go mod download
```

### Application Crashes

**Check logs:**
```bash
oc logs deployment/my-go-app
```

**Common issues:**
- Port already in use
- Missing environment variables
- Insufficient permissions

## 🧪 Testing

### Run Tests Locally
```bash
go test ./...
```

### Add Tests
Create a `main_test.go` file:

```go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(healthHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}
```

## 📝 License

This project is licensed under the MIT License.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📚 Additional Resources

- [OpenShift Pipelines Documentation](https://docs.openshift.com/container-platform/latest/cicd/pipelines/understanding-openshift-pipelines.html)
- [Source-to-Image (S2I)](https://github.com/openshift/source-to-image)
- [Go Documentation](https://go.dev/doc/)
- [Tekton Documentation](https://tekton.dev/docs/)

## 📧 Support

For issues and questions, please open an issue in the GitHub repository.
