package middleware

import (
	"log"
	"net/http"
	"time"
)

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

// LoggingMiddleware logs the incoming HTTP request and its duration
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
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
