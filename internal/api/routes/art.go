package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jorge2751/GoAPI/internal/api/data"
)

// ArtResponse is the response structure for art endpoints
type ArtResponse struct {
	Status string   `json:"status"`
	Data   data.Art `json:"data"`
}

// ArtHandler returns ASCII art as JSON
func ArtHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Create a new art service
	artService := data.NewArtService()

	// Get the art
	art := artService.GetArt()

	// Create response
	response := ArtResponse{
		Status: "success",
		Data:   art,
	}

	// Encode and send response
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// ArtTextHandler returns ASCII art directly as plain text
func ArtTextHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to plain text
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Create a new art service
	artService := data.NewArtService()

	// Get the art
	art := artService.GetArt()

	// Write the art content directly to the response
	_, err := w.Write([]byte(art.Content))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
