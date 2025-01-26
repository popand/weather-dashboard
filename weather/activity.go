package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

var weatherURL = "https://api.openweathermap.org/data/2.5/weather"

type Config struct {
	APIKey string
}

type weatherResponse struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
}

type WeatherResult struct {
	Temperature  float64 `json:"temperature"`
	Conditions   string  `json:"conditions"`
	City         string  `json:"city"`
	AICommentary string  `json:"aiCommentary"`
}

func GetWeatherActivity(ctx context.Context, city string) (*WeatherResult, error) {
	// Read config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	apiKey := viper.GetString("weather.api_key")

	escapedCity := url.QueryEscape(city)
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", weatherURL, escapedCity, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API returned status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var weather weatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %w", err)
	}

	// Get the weather conditions
	conditions := ""
	if len(weather.Weather) > 0 {
		conditions = weather.Weather[0].Main
	}

	// Replace the anonymous struct with WeatherResult
	result := &WeatherResult{
		Temperature: weather.Main.Temp,
		Conditions:  conditions,
		City:        weather.Name,
	}

	return result, nil
}
