package routes

import (
	"encoding/json"
	"net/http"
)

// Response represents the API response structure
type Response struct {
	Message string `json:"message"`
}

// HelloWorldHandler returns a simple hello world JSON response
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
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

// RegisterRoutes sets up all API routes with the given mux
func RegisterRoutes(mux *http.ServeMux, middleware func(http.HandlerFunc) http.HandlerFunc) {
	// Register routes with middleware
	mux.HandleFunc("/hello_world", middleware(HelloWorldHandler))
	mux.HandleFunc("/quotes/random", middleware(RandomQuoteHandler))
	mux.HandleFunc("/art", middleware(ArtHandler))
	mux.HandleFunc("/art/text", middleware(ArtTextHandler))
}
