package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jorge2751/GoAPI/internal/api/routes"
)

func TestArtHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/art", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.ArtHandler)

	// Call the handler with our request and response recorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the content type
	expectedContentType := "text/plain; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}

	// Verify that we got some content (should be the ASCII art)
	if len(rr.Body.String()) < 100 {
		t.Errorf("Response body too short, expected ASCII art")
	}

	// Verify that the response starts with the expected pattern
	expectedStart := "MMMMMMM"
	if rr.Body.String()[0:7] != expectedStart {
		t.Errorf("Response doesn't start with expected pattern: got %s, want %s",
			rr.Body.String()[0:7], expectedStart)
	}
}
