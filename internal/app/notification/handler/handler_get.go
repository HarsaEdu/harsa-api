package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (notificationHandler *NotificationHandlerImpl) GetAll(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)

	params := ctx.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params limit not valid", err)
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return res.StatusBadRequest(ctx, "params offset not valid", err)
	}

	response, pagiantion, err := notificationHandler.NotificationService.GetAll(offset, limit, userId)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "notification not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get all notification, something happen", err)
	}

	return res.StatusOK(ctx, "success to get notification", response, pagiantion)
}

func (notificationHandler *NotificationHandlerImpl) GetById(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	response, err := notificationHandler.NotificationService.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "notification not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get notification, something happen", err)
	}

	return res.StatusOK(ctx, "success to get notification", response, err)
}
