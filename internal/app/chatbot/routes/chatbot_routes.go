package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (chatbotRoutes *ChatbotRoutesImpl) Chatbot(apiGroup *echo.Group) {
	chatBotGroup := apiGroup.Group("/chatbot")

	chatBotGroup.POST("", chatbotRoutes.ChatbotHandler.GetResponse, middleware.StudentMiddleare)
}