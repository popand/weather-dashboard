package weather

import (
	"encoding/json"
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type WeatherInfo struct {
	Temperature float64 `json:"temp"`
	FeelsLike   float64 `json:"feels_like"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	City        string  `json:"city"`
}

type WeatherResult struct {
	Temperature  float64 `json:"temperature"`
	Conditions   string  `json:"conditions"`
	AICommentary string  `json:"aiCommentary"`
}

func WeatherWorkflow(ctx workflow.Context, city string) (*WeatherResult, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var weatherData string
	err := workflow.ExecuteActivity(ctx, GetWeatherActivity, city).Get(ctx, &weatherData)
	if err != nil {
		return nil, err
	}

	// Parse the weather data
	var result WeatherResult
	if err := json.Unmarshal([]byte(weatherData), &result); err != nil {
		return nil, err
	}

	// Get AI commentary
	var aiMessage string
	weatherDesc := fmt.Sprintf("Temperature: %.1f°C, Conditions: %s", result.Temperature, result.Conditions)
	err = workflow.ExecuteActivity(ctx, GetAICommentaryActivity, weatherDesc).Get(ctx, &aiMessage)
	if err != nil {
		return nil, err
	}

	result.AICommentary = aiMessage
	return &result, nil
}
