package weather

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

const aiPrompt = `You are a friendly weather assistant. Given this weather forecast: "%s", 
provide a brief, friendly message that includes both the forecast information and some cheerful advice 
or observation about the weather. Keep the response concise and engaging.`

func GetAICommentaryActivity(ctx context.Context, forecast string) (string, error) {
	apiKey := viper.GetString("openai.api_key")
	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(aiPrompt, forecast),
				},
			},
			MaxTokens: 150,
		},
	)

	if err != nil {
		return "", fmt.Errorf("error getting AI commentary: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from AI")
	}

	return resp.Choices[0].Message.Content, nil
}
