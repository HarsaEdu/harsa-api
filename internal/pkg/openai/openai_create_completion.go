package openai

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sashabaranov/go-openai"
)

func (openAiClient *OpenAiClientImpl) CreateCompletion(prompt string) (string, error) {
	client := openai.NewClient(openAiClient.OpenAIConfig.ApiKey)

	intTemperature, err := strconv.Atoi(openAiClient.OpenAIConfig.Temperature)
	if err != nil {
		return "", fmt.Errorf("error when convert : %s", err.Error())
	}

	float32Temperature := float32(intTemperature)

	intMaxTokens, err := strconv.Atoi(openAiClient.OpenAIConfig.MaxTokens)
	if err != nil {
		return "", fmt.Errorf("error when convert : %s", err.Error())
	}


	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openAiClient.OpenAIConfig.Model,
			Temperature: float32Temperature,
			MaxTokens: intMaxTokens,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: openAiClient.OpenAIConfig.SystemRole},
				{Role: openai.ChatMessageRoleUser, Content: prompt},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("error when create completion : %s", err.Error())
	}

	return response.Choices[0].Message.Content, nil
}