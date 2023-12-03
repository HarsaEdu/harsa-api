package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseTrackingHandler *CourseTrackingHandlerImpl) FindSub(ctx echo.Context) error {

	id := ctx.Get("user_id").(uint)

	idParam := ctx.Param("module-id")
	courseId, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid course id", err)
	}
	
	allSub,err := courseTrackingHandler.CourseTrackingService.FindSubByIdMobile(uint(courseId), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "sub in module not found", err)
		}
		return res.StatusInternalServerError(ctx, "failed to get all sub in module, something happen", err)

	}

	return res.StatusOK(ctx, "success to get all sub in module", allSub, nil)
}

