package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (notificationHandler *NotificationHandlerImpl) ReadNotification(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind notification update read request", err)
	}

	err = notificationHandler.NotificationService.ReadNotification(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "notification not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed update notification, something happen", err)

	}
	return res.StatusOK(ctx, "success to update notification", nil, nil)

}
func (notificationHandler *NotificationHandlerImpl) ArsipNotification(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	req := web.ArsipNotification{}
	req.ID = uint(id)
	err = ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind notification update read request", err)
	}

	err = notificationHandler.NotificationService.ArsipNotification(req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "notification not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed update notification, something happen", err)

	}
	return res.StatusOK(ctx, "success to update notification", nil, nil)

}
