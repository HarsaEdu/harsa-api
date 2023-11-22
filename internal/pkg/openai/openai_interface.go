package openai

import (
	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/sashabaranov/go-openai"
)

type OpenAi interface {
	CreateThread(studentName string, topic string) (string, error)
	GetAllMessagesInThread(threadId string, limit int, after string, before string) (openai.MessagesList, error)
	// CreateMessageInThread(threadId string, message string) (string, error)
	// runMessageInThread(ctx context.Context, threadId string) (string, error)
	// getResponseMessageInThread(ctx context.Context, threadId string) (string, error)
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