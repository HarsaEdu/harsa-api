package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (chatbotHandler *ChatbotHandlerImpl) Delete(ctx echo.Context) error {
	threadIdParam := ctx.Param("id")

	userId := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	err := chatbotHandler.ChatbotService.Delete(threadIdParam, userId, roleString)
	if err != nil {

		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "topic not found", err)
		}

		if strings.Contains(err.Error(), "not allowed") {
			return res.StatusUnauthorized(ctx, "not allowed to delete this topic", err)
		}

		return res.StatusInternalServerError(ctx, "failed to delete topic, something happen", err)
	}

	return res.StatusOK(ctx, "success to delete topic", nil, nil)
}