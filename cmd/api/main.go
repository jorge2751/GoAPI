package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jorge2751/GoAPI/internal/api/middleware"
	"github.com/jorge2751/GoAPI/internal/api/routes"
)

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Define HTTP server
	mux := http.NewServeMux()

	// Register routes with middleware
	routes.RegisterRoutes(mux, middleware.LoggingMiddleware)

	// Start server
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(addr, mux))
}
