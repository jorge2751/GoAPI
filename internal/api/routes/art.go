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

// ArtHandler returns ASCII art
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
