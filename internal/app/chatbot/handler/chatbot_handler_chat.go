package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (chatbotHandler *ChatbotHandlerImpl) ChatWithAssistant(ctx echo.Context) error {
	threadIdParam := ctx.Param("id")

	request := web.ChatWithAssistantRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, fmt.Sprintf("failed to bind request : %s", err.Error()), err)
	}

	response, err := chatbotHandler.ChatbotService.ChatWithAssistant(threadIdParam, &request)
	if err != nil {

		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "topic not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to chat with assistant, something happen", err)
	}

	return res.StatusOK(ctx, "success to chat with assistant", response, nil)
}