package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (chatbotHandler *ChatbotHandlerImpl) CreateThread(ctx echo.Context) error {
	request := web.CreateThreadRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, fmt.Sprintf("failed to bind request : %s", err.Error()), err)
	}

	userId := ctx.Get("user_id").(uint)

	err = chatbotHandler.ChatbotService.CreateThread(&request, userId)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "user not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get create thread, something happen", err)
	}

	return res.StatusCreated(ctx, "success to create response", nil, nil)
}