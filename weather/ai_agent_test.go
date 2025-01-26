package weather

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAICommentaryActivity(t *testing.T) {
	weatherDesc := "Temperature: 20.5°C, Conditions: Sunny"

	result, err := GetAICommentaryActivity(context.Background(), weatherDesc)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	// Note: Can't test exact response as AI output varies
	assert.Contains(t, result, "°C")
}
