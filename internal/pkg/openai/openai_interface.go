package openai

import (
	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/sashabaranov/go-openai"
)

type OpenAi interface {
	CreateThread(studentName string, topic string) (string, error)
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