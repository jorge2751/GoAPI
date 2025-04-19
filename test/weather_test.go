package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jorge2751/GoAPI/internal/api/routes"
)

// Mock WeatherAPI server
func startMockWeatherAPIServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.URL.Query().Get("key")
		city := r.URL.Query().Get("q")

		if apiKey == "" || city == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, `{"error":{"message":"Missing key or query"}}`)
			return
		}

		if city == "errorcity" {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"error":{"message":"Internal API error simulation"}}`)
			return
		}

		// Simulate successful response for any other city
		w.WriteHeader(http.StatusOK)
		response := fmt.Sprintf(`{
            "location": {
                "name": "%s",
                "region": "Test Region",
                "country": "Test Country"
            },
            "current": {
                "temp_b": 15.0,
                "condition": {
                    "text": "Partly cloudy"
                }
            }
        }`, city)
		fmt.Fprintln(w, response)
	}))
}

func TestWeatherHandler(t *testing.T) {
	mockAPIServer := startMockWeatherAPIServer()
	defer mockAPIServer.Close()

	// Create a WeatherService instance pointing to the mock server
	// We override the HTTPClient to ensure it hits our mock server
	weatherService := routes.NewWeatherService("test-api-key")
	weatherService.HTTPClient = mockAPIServer.Client()
	weatherService.BaseURL = mockAPIServer.URL // Set BaseURL to mock server

	// --- Test Cases ---

	// Test Case 1: Successful request
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/weather?city=TestCity", nil)
		w := httptest.NewRecorder()
		weatherService.WeatherHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status OK; got %v", w.Code)
		}

		var resp routes.WeatherAPIResponse
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		if resp.Location.Name != "TestCity" {
			t.Errorf("Expected location name TestCity; got %s", resp.Location.Name)
		}
		if resp.Current.TempB != 15.0 {
			t.Errorf("Expected temp 15.0; got %f", resp.Current.TempB)
		}
	})

	// Test Case 2: Missing city parameter
	t.Run("MissingCityParam", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/weather", nil)
		w := httptest.NewRecorder()
		weatherService.WeatherHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status BadRequest; got %v", w.Code)
		}
		if !strings.Contains(w.Body.String(), "Query parameter 'city' is required") {
			t.Errorf("Expected error message about missing city; got %s", w.Body.String())
		}
	})

	// Test Case 3: Missing API Key (in service config)
	t.Run("MissingAPIKey", func(t *testing.T) {
		// Create a service instance without an API key
		badWeatherService := routes.NewWeatherService("")
		badWeatherService.HTTPClient = mockAPIServer.Client()

		req := httptest.NewRequest("GET", "/weather?city=SomeCity", nil)
		w := httptest.NewRecorder()
		badWeatherService.WeatherHandler(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status InternalServerError; got %v", w.Code)
		}
		if !strings.Contains(w.Body.String(), "WeatherAPI key not configured in service") {
			t.Errorf("Expected error message about missing API key; got %s", w.Body.String())
		}
	})

	// Test Case 4: Weather API returns an error
	t.Run("WeatherAPIError", func(t *testing.T) {
		// We need to modify the service's HTTPClient to point to the *correct* field
		// in the mock server URL struct for the fetch URL construction.
		// The Sprintf uses the raw URL string, not the client base URL.
		// So, we create a new service pointing directly to the mock server base URL for the *fetch URL*.
		// The actual client used will still be the mockAPIServer.Client().
		// This is a bit awkward due to how the fetch URL is constructed directly in the handler.
		// A further refactor could move URL construction into the service.

		// NOTE: Correcting the approach - the URL construction needs the API Key,
		// so we can't just point the client base. We need to use the mock server's URL
		// *within* the handler logic if we want to fully mock the external API.
		// However, the current refactor has the URL hardcoded. The mock client
		// handles the request redirection to the mock server correctly.
		// So, we just need to trigger the error condition in the mock server.

		req := httptest.NewRequest("GET", "/weather?city=errorcity", nil)
		w := httptest.NewRecorder()
		weatherService.WeatherHandler(w, req) // Use the original service with mock client

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status InternalServerError on API error; got %v", w.Code)
		}
		if !strings.Contains(w.Body.String(), "WeatherAPI request failed with status 500") {
			t.Errorf("Expected error message about API failure; got %s", w.Body.String())
		}
	})
}
