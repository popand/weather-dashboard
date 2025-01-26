package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func TestWeatherWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock the GetWeatherActivity
	env.OnActivity(GetWeatherActivity, mock.Anything, "London").Return(&WeatherResult{
		Temperature: 20.5,
		Conditions:  "Sunny",
		City:        "London",
	}, nil)

	// Mock the GetAICommentaryActivity
	env.OnActivity(GetAICommentaryActivity, mock.Anything, mock.Anything).Return(
		"It's a beautiful sunny day in London!", nil)

	// Execute workflow
	var result WeatherResult
	env.ExecuteWorkflow(WeatherWorkflow, "London")

	// Verify workflow completed successfully
	assert.True(t, env.IsWorkflowCompleted())
	assert.NoError(t, env.GetWorkflowError())

	// Get workflow result
	assert.NoError(t, env.GetWorkflowResult(&result))

	// Verify results
	assert.Equal(t, 20.5, result.Temperature)
	assert.Equal(t, "Sunny", result.Conditions)
	assert.Equal(t, "London", result.City)
	assert.Equal(t, "It's a beautiful sunny day in London!", result.AICommentary)
}
