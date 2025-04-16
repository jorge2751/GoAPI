package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jorge2751/go-hello-api/internal/api/data"
)

// QuoteResponse is the response structure for quote endpoints
type QuoteResponse struct {
	Status string     `json:"status"`
	Data   data.Quote `json:"data"`
}

// RandomQuoteHandler returns a random quote
func RandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Create a new quote service
	quoteService := data.NewQuoteService()

	// Get a random quote
	randomQuote := quoteService.GetRandomQuote()

	// Create response
	response := QuoteResponse{
		Status: "success",
		Data:   randomQuote,
	}

	// Encode and send response
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
