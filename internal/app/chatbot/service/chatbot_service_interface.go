package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/chatbot/repository"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/openai"
	"github.com/go-playground/validator"
)

type ChatbotService interface {
	CreateThread(request *web.CreateThreadRequest, userId uint) error
	GetAllMessagesInThread(threadId string, limit int, after string, before string) ([]web.GetMessageInThreadResponse, error)
	// GetResponse(request *web.GetResponseRequest) (string, error)
}

type ChatbotServiceImpl struct {
	ChatbotRepository repository.ChatbotRepository
	UserRepository userRepositoryPkg.UserRepository
	Validate *validator.Validate
	OpenAi openai.OpenAi
}

func NewChatbotService(ChatbotRepository repository.ChatbotRepository, UserRepository userRepositoryPkg.UserRepository, Validate *validator.Validate, OpenAi openai.OpenAi) ChatbotService {
	return &ChatbotServiceImpl{
		ChatbotRepository: ChatbotRepository,
		UserRepository: UserRepository,
		Validate: Validate,
		OpenAi: OpenAi,
	}
}