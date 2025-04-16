package data

import (
	"math/rand"
	"time"
)

// Quote represents a quote with its text and author
type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

// QuoteService provides quote-related functionality
type QuoteService struct {
	quotes []Quote
	r      *rand.Rand
}

// NewQuoteService creates a new QuoteService with predefined quotes
func NewQuoteService() *QuoteService {
	// Create a random source with time seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Define our list of quotes
	quotes := []Quote{
		{Text: "Life is what happens when you're busy making other plans.", Author: "John Lennon"},
		{Text: "The way to get started is to quit talking and begin doing.", Author: "Walt Disney"},
		{Text: "Your time is limited, so don't waste it living someone else's life.", Author: "Steve Jobs"},
		{Text: "The future belongs to those who believe in the beauty of their dreams.", Author: "Eleanor Roosevelt"},
		{Text: "The purpose of our lives is to be happy.", Author: "Dalai Lama"},
		{Text: "Get busy living or get busy dying.", Author: "Stephen King"},
		{Text: "You only live once, but if you do it right, once is enough.", Author: "Mae West"},
		{Text: "Many of life's failures are people who did not realize how close they were to success when they gave up.", Author: "Thomas A. Edison"},
		{Text: "The secret of success is to do the common thing uncommonly well.", Author: "John D. Rockefeller Jr."},
		{Text: "The best time to plant a tree was 20 years ago. The second best time is now.", Author: "Chinese Proverb"},
	}

	return &QuoteService{
		quotes: quotes,
		r:      r,
	}
}

// GetRandomQuote returns a random quote from the collection
func (qs *QuoteService) GetRandomQuote() Quote {
	// Get a random index from our quotes slice
	randomIndex := qs.r.Intn(len(qs.quotes))
	return qs.quotes[randomIndex]
}
