package handler

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (chatbotHandler *ChatbotHandlerImpl) GetResponse(ctx echo.Context) error {
	getResponseRequest := web.GetResponseRequest{}
	err := ctx.Bind(&getResponseRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, fmt.Sprintf("failed to bind request : %s", err.Error()), err)
	}

	response, err := chatbotHandler.ChatbotService.GetResponse(&getResponseRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		return res.StatusInternalServerError(ctx, "failed to get message response, something happen", err)
	}

	return res.StatusOK(ctx, "success to get response", response, nil)
}