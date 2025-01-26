package weather

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeatherActivity(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify query parameters
		assert.Contains(t, r.URL.String(), "q=London")
		assert.Contains(t, r.URL.String(), "units=metric")

		// Return mock weather data
		w.Write([]byte(`{
			"main": {
				"temp": 20.5,
				"feels_like": 19.8,
				"humidity": 65
			},
			"weather": [{"main": "Clouds", "description": "scattered clouds"}],
			"wind": {"speed": 4.1},
			"name": "London"
		}`))
	}))
	defer server.Close()

	// Override the weather URL for testing
	originalURL := weatherURL
	weatherURL = server.URL
	defer func() { weatherURL = originalURL }()

	// Test the activity
	result, err := GetWeatherActivity(context.Background(), "London")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 20.5, result.Temperature)
	assert.Equal(t, "Clouds", result.Conditions)
	assert.Equal(t, "London", result.City)
}
