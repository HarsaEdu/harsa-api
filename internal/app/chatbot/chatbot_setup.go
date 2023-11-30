package chatbot

import (
	chatbotHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/chatbot/handler"
	chatbotRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/chatbot/repository"
	chatbotRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/chatbot/routes"
	chatbotServicePkg "github.com/HarsaEdu/harsa-api/internal/app/chatbot/service"
	userbotRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/pkg/openai"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func ChatbotSetup(db *gorm.DB, validate *validator.Validate, userbotRepositoryPkg userbotRepositoryPkg.UserRepository, openai openai.OpenAi) (chatbotRoutesPkg.ChatbotRoutes){
	chatbotRepository := chatbotRepositoryPkg.NewChatbotRepository(db)
	chatbotService := chatbotServicePkg.NewChatbotService(chatbotRepository, userbotRepositoryPkg, validate, openai)
	chatbotHandler := chatbotHandlerPkg.NewChatbotHandler(chatbotService)
	chatbotRoute := chatbotRoutesPkg.NewChatbotRoutes(chatbotHandler)

	return chatbotRoute
}
