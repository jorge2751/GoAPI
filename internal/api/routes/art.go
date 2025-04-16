package routes

import (
	"net/http"

	"github.com/jorge2751/GoAPI/internal/api/data"
)

// ArtHandler returns ASCII art directly as plain text
func ArtHandler(w http.ResponseWriter, r *http.Request) {
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
