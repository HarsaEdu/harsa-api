package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/chatbot/service"
	"github.com/labstack/echo/v4"
)

type ChatbotHandler interface {
	CreateThread(ctx echo.Context) error
	GetAllThreadByUserId(ctx echo.Context) error
	GetAllMessagesInThread(ctx echo.Context) error
	// GetResponse(ctx echo.Context) error
}

type ChatbotHandlerImpl struct {
	ChatbotService service.ChatbotService
}

func NewChatbotHandler(ChatbotService service.ChatbotService) ChatbotHandler {
	return &ChatbotHandlerImpl{
		ChatbotService: ChatbotService,
	}
}