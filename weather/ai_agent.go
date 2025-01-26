package weather

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

const aiPrompt = `You are a friendly weather assistant. Given this weather forecast: "%s", 
provide a brief, friendly message that includes both the forecast information and some cheerful advice 
or observation about the weather. Keep the response concise and engaging.`

func GetAICommentaryActivity(ctx context.Context, forecast string) (string, error) {
	client := openai.NewClient("sk-svcacct-9oS4j4096NFiMt3cRmWj_07LNPQcAmEHkhaYA0ePLGLN-IUwFrb2ajNBhCzd57P_RfvT3BlbkFJAPddv_8EnZGSkIman0CYcLxaFXctZ88PFuJdqAKliuwj9LwQUISRKlq5P33yrT0F8lgA")

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
