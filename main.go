package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
	Uptime    string    `json:"uptime"`
}

type AppResponse struct {
	Message     string `json:"message"`
	Application string `json:"application"`
	Environment string `json:"environment"`
}

var startTime time.Time

func main() {
	startTime = time.Now()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Register HTTP handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ready", readyHandler)
	http.HandleFunc("/api/info", infoHandler)

	log.Printf("Starting server on port %s...", port)
	log.Printf("Endpoints available:")
	log.Printf("  - http://localhost:%s/", port)
	log.Printf("  - http://localhost:%s/health", port)
	log.Printf("  - http://localhost:%s/ready", port)
	log.Printf("  - http://localhost:%s/api/info", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Go S2I Application</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 50px auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background-color: white;
            padding: 30px;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        h1 {
            color: #333;
        }
        .endpoints {
            background-color: #f9f9f9;
            padding: 15px;
            border-radius: 5px;
            margin-top: 20px;
        }
        .endpoint {
            margin: 10px 0;
        }
        a {
            color: #0066cc;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
        .success {
            color: #28a745;
            font-weight: bold;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Go Application Running Successfully!</h1>
        <p class="success">✓ Your OpenShift S2I Go Pipeline deployment is working!</p>
        
        <h2>Available Endpoints:</h2>
        <div class="endpoints">
            <div class="endpoint">
                <strong>🏠 Home:</strong> <a href="/">/</a> - This page
            </div>
            <div class="endpoint">
                <strong>💚 Health:</strong> <a href="/health">/health</a> - Health check endpoint (JSON)
            </div>
            <div class="endpoint">
                <strong>✅ Ready:</strong> <a href="/ready">/ready</a> - Readiness probe endpoint (JSON)
            </div>
            <div class="endpoint">
                <strong>📊 Info:</strong> <a href="/api/info">/api/info</a> - Application information (JSON)
            </div>
        </div>

        <h2>About This Application:</h2>
        <p>This is a Go application built and deployed using:</p>
        <ul>
            <li>OpenShift Pipelines (Tekton)</li>
            <li>Source-to-Image (S2I) with s2i-go ClusterTask</li>
            <li>Standard Go HTTP server</li>
        </ul>
    </div>
</body>
</html>
`
	fmt.Fprint(w, html)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	uptime := time.Since(startTime).Round(time.Second)
	
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
		Uptime:    uptime.String(),
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Add any readiness checks here (database connection, etc.)
	ready := map[string]interface{}{
		"status": "ready",
		"checks": map[string]string{
			"server": "ok",
		},
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ready)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	
	response := AppResponse{
		Message:     "Go application built with OpenShift S2I",
		Application: "s2i-go-demo",
		Environment: env,
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
