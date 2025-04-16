package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jorge2751/go-hello-api/internal/api/routes"
)

func TestHelloWorldHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/hello_world", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.HelloWorldHandler)

	// Call the handler with our request and response recorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the content type
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}

	// Check the response body
	var response routes.Response
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	// Verify the message field
	expectedMessage := "Hello World from Go API!"
	if response.Message != expectedMessage {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expectedMessage)
	}
}
