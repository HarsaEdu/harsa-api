package handler

import (
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseTrackingHandler *DashboardHandlerImpl) CountAll(ctx echo.Context) error {

	result, err := courseTrackingHandler.DashboardService.CountAll()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "category not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get all count dashboard, something happen", err)
	}
	return res.StatusOK(ctx, "success to get all count dashboard", result, nil)
}