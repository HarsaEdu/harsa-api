package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (chatbotRoutes *ChatbotRoutesImpl) Chatbot(apiGroup *echo.Group) {
	chatBotGroup := apiGroup.Group("/chatbot")

	chatBotGroup.POST("", chatbotRoutes.ChatbotHandler.CreateThread, middleware.StudentMiddleare)
	chatBotGroup.GET("", chatbotRoutes.ChatbotHandler.GetAllThreadByUserId, middleware.StudentMiddleare)
	chatBotGroup.GET("/:id", chatbotRoutes.ChatbotHandler.GetAllMessagesInThread, middleware.StudentMiddleare)
	// chatBotGroup.POST("", chatbotRoutes.ChatbotHandler.GetResponse, middleware.StudentMiddleare)
}