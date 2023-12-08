package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (chatbotHandler *ChatbotHandlerImpl) GetAllMessagesInThread(ctx echo.Context) error {
	threadIdParam := ctx.Param("id")

	limitParam := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit data request", err)
	}

	afterParam := ctx.QueryParam("after")

	beforeParam := ctx.QueryParam("before")

	response, err := chatbotHandler.ChatbotService.GetAllMessagesInThread(threadIdParam, limit, afterParam, beforeParam)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "messages not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get all message, something happen", err)
	}

	return res.StatusOK(ctx, "success to get all message", response, nil)
}

func (chatbotHandler *ChatbotHandlerImpl) GetAllThreadByUserId(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)

	response, err := chatbotHandler.ChatbotService.GetAllThreadByUserId(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "thread not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get all thread, something happen", err)
	}

	return res.StatusOK(ctx, "success to get all thread", response, nil)
}