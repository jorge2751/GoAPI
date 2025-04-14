package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Response represents the API response structure
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Define HTTP server
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/hello_world", helloWorldHandler)

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