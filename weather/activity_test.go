package weather

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetWeatherActivity(t *testing.T) {
	// Create a mock HTTP server
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

	// Save original URL and restore it after test
	originalURL := weatherURL
	weatherURL = server.URL + "/data/2.5/weather"
	defer func() {
		weatherURL = originalURL
	}()

	// Configure Viper for testing - Fix the config path
	viper.Reset()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")    // Look in parent directory's config folder
	viper.AddConfigPath("../../config") // Also look two directories up
	viper.AddConfigPath("config")       // Also look in current directory
	viper.AddConfigPath("testdata")     // Add this line after the other AddConfigPath calls

	// Set a default API key in case config file is not found
	viper.SetDefault("weather.api_key", "test-api-key")

	if err := viper.ReadInConfig(); err != nil {
		t.Logf("Warning: Could not read config file: %v", err)
		// Continue with default values
	}

	// Test the activity
	result, err := GetWeatherActivity(context.Background(), "London")
	if err != nil {
		t.Fatalf("GetWeatherActivity failed: %v", err)
	}

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 20.5, result.Temperature)
	assert.Equal(t, "Clouds", result.Conditions)
	assert.Equal(t, "London", result.City)
}
