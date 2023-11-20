package openai

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

func (openAiClient *OpenAiClientImpl) CreateChatCompletion(prompt string) (string, error) {
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

func (openAiClient *OpenAiClientImpl) CreateChatCompletionStream(prompt string) (string, error) {
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

	response, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openAiClient.OpenAIConfig.Model,
			Temperature: float32Temperature,
			MaxTokens:   intMaxTokens,
			Stream:      true,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: openAiClient.OpenAIConfig.SystemRole},
				{Role: openai.ChatMessageRoleUser, Content: prompt},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("error when create completion : %s", err.Error())
	}
	defer response.Close()

	var result string
	var responseMessage string

	for {
			response, err := response.Recv()
			if errors.Is(err, io.EOF) {
					fmt.Println("Stream finished")
					break
			}
	
			if err != nil {
					fmt.Printf("Stream error: %v\n", err)
					return "", fmt.Errorf("error when create completion : %s", err.Error())
			}
	
			// Assuming Choices is a slice of structs with a Text field
			if len(response.Choices) > 0 {
				result = response.Choices[0].Delta.Content
				logrus.Info(result)

				responseMessage += result 
			}
	}
	
	return responseMessage, nil
	
}
