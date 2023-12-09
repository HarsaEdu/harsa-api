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
	GetAllThreadByUserId(userId uint) ([]web.GetAllThreadByUserIdResponse, error)
	GetAllMessagesInThread(threadId string, limit int, after string, before string) ([]web.GetMessageInThreadResponse, error)
	ChatWithAssistant(threadId string, request *web.ChatWithAssistantRequest) (string, error)
	Delete(threadId string, userID uint, role string)  error
	Update(threadId string, userID uint, update *web.CreateThreadRequest , role string) error
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