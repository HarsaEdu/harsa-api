package openai

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func (openai *OpenAiImpl) ChatWithAssistant(threadId string, message string) (string, error) {
	ctx := context.Background()

	assistant, err := openai.Client.RetrieveAssistant(ctx, openai.OpenAIConfig.AssistantID)
	if err != nil {
		return "", err
	}

	_, err = openai.CreateMessageInThread(ctx, threadId, message)
	if err != nil {
		return "", err
	}

	_, err = openai.RunMessageInThread(ctx, assistant, threadId)
	if err != nil {
		return "", err
	}

	responseMessage, err := openai.GetResponseMessageInThread(ctx, threadId)
	if err != nil {
		return "", err
	}

	return responseMessage.Content[0].Text.Value, nil
}

func (openAi *OpenAiImpl) GetChatCompletion(message, systemInstruction string) (string, error) {
	ctx := context.Background()

	request := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			openai.ChatCompletionMessage{
				Role: openai.ChatMessageRoleSystem,
				Content: systemInstruction,
			},
			openai.ChatCompletionMessage{
				Role: openai.ChatMessageRoleUser,
				Content: message,
			},
		},
	}

	response, err := openAi.Client.CreateChatCompletion(ctx, request)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}