package service

import (
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/openai"
	"github.com/go-playground/validator"
)

type ChatbotService interface {
	GetResponse(request *web.GetResponseRequest) (string, error)
}

type ChatbotServiceImpl struct {
	Validate *validator.Validate
	OpenAiClient openai.OpenAiClient
}

func NewChatbotService(Validate *validator.Validate, OpenAiClient openai.OpenAiClient) ChatbotService {
	return &ChatbotServiceImpl{
		Validate: Validate,
		OpenAiClient: OpenAiClient,
	}
}