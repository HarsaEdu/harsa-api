package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (courseTrackingHandler *CourseTrackingHandlerImpl) Create(ctx echo.Context) error {

	req := web.CourseTrackingRequest{}

	id := ctx.Get("user_id").(uint)

	idParam := ctx.Param("course-id")
	courseId, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid course id", err)
	}

	req.CourseID = uint(courseId)
	req.UserID = id
	req.Status = "in progress"
	err = courseTrackingHandler.CourseTrackingService.Create(ctx, req)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"subscribe to enrolled this course" ,err)
		}
		if strings.Contains(err.Error(), "already exist") {
			return res.StatusBadRequest(ctx, "You are already enrolled in this course", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create course tracking, something happen", err)

	}

	return res.StatusCreated(ctx, "success to create course tracking", nil, nil)
}

func (courseTrackingHandler *CourseTrackingHandlerImpl) CreateWeb(ctx echo.Context) error {

	req := web.CourseTrackingRequest{}

	idUserParam := ctx.Param("user-id")
	userId, err := strconv.Atoi(idUserParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid user id", err)
	}

	idParam := ctx.Param("course-id")
	courseId, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid course id", err)
	}
	

	req.CourseID = uint(courseId)
	req.UserID = uint(userId)
	req.Status = "in progress"
	err = courseTrackingHandler.CourseTrackingService.CreateWeb(req)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"subscribe to enrolled this course" ,err)
		}
		if strings.Contains(err.Error(), "already exist") {
			return res.StatusBadRequest(ctx, " already enrolled in this course", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create course tracking, something happen", err)

	}

	return res.StatusCreated(ctx, "success to create course tracking", nil, nil)
}