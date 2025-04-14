package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Response represents the API response structure
type Response struct {
	Message string `json:"message"`
}

// loggingMiddleware logs the incoming HTTP request and its duration
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Start timer
		startTime := time.Now()

		// Create a custom response writer to capture the status code
		crw := &customResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Default to 200 OK
		}

		// Log request details
		log.Printf("Request: %s %s", r.Method, r.URL.Path)

		// Call the actual handler
		next(crw, r)

		// Calculate duration
		duration := time.Since(startTime)

		// Log response details
		log.Printf("Response: %s %s - Status: %d - Duration: %v",
			r.Method, r.URL.Path, crw.statusCode, duration)
	}
}

// customResponseWriter is a wrapper for http.ResponseWriter that captures the status code
type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code before calling the underlying ResponseWriter
func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Define HTTP server
	mux := http.NewServeMux()

	// Register routes with middleware
	mux.HandleFunc("/hello_world", loggingMiddleware(helloWorldHandler))

	// Start server
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Create response
	response := Response{
		Message: "Hello World from Go API!",
	}

	// Encode and send response
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
