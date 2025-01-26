package weather

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func TestWeatherWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock the activity
	env.OnActivity(GetWeatherActivity, mock.Anything, "London").Return(&WeatherInfo{
		Temperature: 20.5,
		FeelsLike:   19.8,
		Humidity:    65,
		WindSpeed:   5.2,
		City:        "London",
	}, nil)

	// Execute workflow
	env.ExecuteWorkflow(WeatherWorkflow, "London")

	// Verify workflow completed successfully
	if !env.IsWorkflowCompleted() {
		t.Fatal("Workflow did not complete")
	}

	if env.GetWorkflowError() != nil {
		t.Fatalf("Workflow failed: %v", env.GetWorkflowError())
	}

	// Get workflow result
	var result *WeatherInfo
	if err := env.GetWorkflowResult(&result); err != nil {
		t.Fatalf("Failed to get workflow result: %v", err)
	}

	// Verify results
	expected := &WeatherInfo{
		Temperature: 20.5,
		FeelsLike:   19.8,
		Humidity:    65,
		WindSpeed:   5.2,
		City:        "London",
	}

	if result.Temperature != expected.Temperature {
		t.Errorf("Temperature = %v, want %v", result.Temperature, expected.Temperature)
	}
	if result.FeelsLike != expected.FeelsLike {
		t.Errorf("FeelsLike = %v, want %v", result.FeelsLike, expected.FeelsLike)
	}
	if result.Humidity != expected.Humidity {
		t.Errorf("Humidity = %v, want %v", result.Humidity, expected.Humidity)
	}
	if result.WindSpeed != expected.WindSpeed {
		t.Errorf("WindSpeed = %v, want %v", result.WindSpeed, expected.WindSpeed)
	}
	if result.City != expected.City {
		t.Errorf("City = %v, want %v", result.City, expected.City)
	}

	// Verify activity was called
	env.AssertExpectations(t)
}
