package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseTrackingHandler *CourseTrackingHandlerImpl) FindSubModuleByID(ctx echo.Context) error {
	moduleID, _ := strconv.Atoi(ctx.Param("module-id"))
	subModuleID, _ := strconv.Atoi(ctx.Param("sub-module-id"))
	id := ctx.Get("user_id").(uint)

	result, err := courseTrackingHandler.CourseTrackingService.FindSubModuleByID(uint(moduleID), uint(subModuleID), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "sub module not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get course tracking, something happen", err)
	}
	return res.StatusOK(ctx, "success to get sub module tracking", result, nil)
}
