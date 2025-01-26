package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWeatherActivity(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request parameters
		if r.URL.Path != "/data/2.5/weather" {
			t.Errorf("Expected path '/data/2.5/weather', got %s", r.URL.Path)
		}

		city := r.URL.Query().Get("q")
		if city != "London" {
			t.Errorf("Expected city 'London', got %s", city)
		}

		// Return mock weather data
		mockResponse := `{
			"main": {
				"temp": 20.5,
				"feels_like": 19.8,
				"humidity": 65
			},
			"wind": {
				"speed": 5.2
			},
			"weather": [
				{
					"main": "Clouds",
					"description": "scattered clouds"
				}
			],
			"name": "London"
		}`
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, mockResponse)
	}))
	defer server.Close()

	// Override the weather URL for testing
	weatherURL = server.URL

	// Test the activity
	result, err := GetWeatherActivity(context.Background(), "London")
	if err != nil {
		t.Fatalf("GetWeatherActivity failed: %v", err)
	}

	// Parse the returned JSON string
	var weatherResult struct {
		Temperature float64 `json:"temperature"`
		Conditions  string  `json:"conditions"`
	}
	if err := json.Unmarshal([]byte(result), &weatherResult); err != nil {
		t.Fatalf("Failed to parse result JSON: %v", err)
	}

	// Verify the results
	if weatherResult.Temperature != 20.5 {
		t.Errorf("Expected temperature 20.5, got %f", weatherResult.Temperature)
	}
	if weatherResult.Conditions != "Clouds" {
		t.Errorf("Expected conditions 'Clouds', got %s", weatherResult.Conditions)
	}
}
