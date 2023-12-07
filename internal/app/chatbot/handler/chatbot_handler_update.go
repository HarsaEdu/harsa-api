package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (chatbotHandler *ChatbotHandlerImpl) Update(ctx echo.Context) error {
	threadIdParam := ctx.Param("id")

	request := web.CreateThreadRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, fmt.Sprintf("failed to bind request : %s", err.Error()), err)
	}

	userId := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	err = chatbotHandler.ChatbotService.Update(threadIdParam, userId, &request, roleString)
	if err != nil {

		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "topic not found", err)
		}

		if strings.Contains(err.Error(), "not allowed") {
			return res.StatusUnauthorized(ctx, "not allowed to update this topic", err)
		}

		return res.StatusInternalServerError(ctx, "failed to update this topic, something happen", err)
	}

	return res.StatusOK(ctx, "success to update this topic", nil, nil)
}