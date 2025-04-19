package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// WeatherService holds dependencies for the weather handler
type WeatherService struct {
	APIKey     string
	HTTPClient *http.Client
	BaseURL    string
}

// NewWeatherService creates a new WeatherService instance
func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL:    "http://api.weatherapi.com/v1",
	}
}

// WeatherAPIResponse defines the structure for the relevant parts of the WeatherAPI response
type WeatherAPIResponse struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempF     float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

// WeatherHandler fetches weather data for a given city
func (s *WeatherService) WeatherHandler(w http.ResponseWriter, r *http.Request) {
	// Get city from query parameters
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "Query parameter 'city' is required", http.StatusBadRequest)
		return
	}

	// Use API key from the service struct
	if s.APIKey == "" {
		http.Error(w, "WeatherAPI key not configured in service", http.StatusInternalServerError)
		fmt.Println("Error: WeatherAPI key not configured in WeatherService.") // Log for server admin
		return
	}

	// Construct WeatherAPI URL using BaseURL
	apiURL := fmt.Sprintf("%s/current.json?key=%s&q=%s&aqi=no", s.BaseURL, s.APIKey, city)

	// Make GET request using the service's HTTP client
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		http.Error(w, "Failed to create weather API request", http.StatusInternalServerError)
		fmt.Printf("Error creating weather request: %v\n", err)
		return
	}

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		fmt.Printf("Error fetching weather data: %v\n", err) // Log error
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body) // Read body for more info if possible
		errorMsg := fmt.Sprintf("WeatherAPI request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
		http.Error(w, errorMsg, http.StatusInternalServerError)
		fmt.Println(errorMsg) // Log error
		return
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read weather data response", http.StatusInternalServerError)
		fmt.Printf("Error reading weather data response: %v\n", err) // Log error
		return
	}

	// Parse JSON response
	var weatherData WeatherAPIResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		http.Error(w, "Failed to parse weather data", http.StatusInternalServerError)
		fmt.Printf("Error parsing weather data JSON: %v\n", err) // Log error
		return
	}

	// Set content type and encode response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(weatherData)
	if err != nil {
		// Don't use http.Error here as headers might have been written
		fmt.Printf("Error encoding weather response: %v\n", err)
	}
}
