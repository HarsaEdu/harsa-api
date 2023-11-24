package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/chatbot/handler"
	"github.com/labstack/echo/v4"
)

type ChatbotRoutes interface {
	ChatbotMobile(apiGroup *echo.Group)
}

type ChatbotRoutesImpl struct {
	ChatbotHandler handler.ChatbotHandler
}

func NewChatbotRoutes(ChatbotHandler handler.ChatbotHandler) ChatbotRoutes {
	return &ChatbotRoutesImpl{
		ChatbotHandler: ChatbotHandler,
	}
}