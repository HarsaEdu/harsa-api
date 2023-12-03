package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (courseTrackingHandler *CourseTrackingHandlerImpl) DeleteEnrolled(ctx echo.Context) error {

	courseParam := ctx.Param("id")
	
	courseId, err := strconv.Atoi(courseParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid course id", err)
	}

	trackingParam := ctx.Param("tracking-id")
	
	trackingId, err := strconv.Atoi(trackingParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid tracking id", err)
	}

	user_id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	err = courseTrackingHandler.CourseTrackingService.Delete(uint(trackingId), uint(courseId),uint(user_id), roleString)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "enrolled not found", err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot delete this enrolled" ,err)
		}
		
		return res.StatusInternalServerError(ctx, "failed to delete enrolled, something happen", err)
	}

	return res.StatusOK(ctx, "success to delete enrolled", nil, nil)
}