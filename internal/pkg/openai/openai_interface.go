package openai

import "github.com/HarsaEdu/harsa-api/configs"

type OpenAiClient interface {
	CreateCompletion(prompt string) (string, error)
	CreateChatCompletionStream(prompt string) (string, error)
}

type OpenAiClientImpl struct {
	OpenAIConfig *configs.OpenAI
}

func NewOpenAiClient(openAIConfig *configs.OpenAI) OpenAiClient {
	return &OpenAiClientImpl{
		OpenAIConfig: openAIConfig,
	}
}