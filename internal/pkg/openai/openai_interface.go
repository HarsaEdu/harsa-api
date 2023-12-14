package openai

import (
	"context"

	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/sashabaranov/go-openai"
)

type OpenAi interface {
	CreateThread(studentName string, topic string) (string, error)
	GetAllMessagesInThread(threadId string, limit int, after string, before string) (openai.MessagesList, error)
	CreateMessageInThread(ctx context.Context, threadId string, message string) (openai.Message, error)
	RunMessageInThread(ctx context.Context, assistant openai.Assistant, threadId string) (openai.Run, error)
	GetResponseMessageInThread(ctx context.Context, threadId string) (openai.Message, error)
	ChatWithAssistant(threadId string, message string) (string, error)
	GetChatCompletion(message, systemInstruction string) (string, error)
}

type OpenAiImpl struct {
	OpenAIConfig *configs.OpenAI
	Client *openai.Client
}

func NewOpenAi(openAIConfig *configs.OpenAI) OpenAi {
	return &OpenAiImpl{
		OpenAIConfig: openAIConfig,
		Client: openai.NewClient(openAIConfig.ApiKey),
	}
}