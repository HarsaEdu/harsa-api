package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/course_tracking/service"
	"github.com/labstack/echo/v4"
)

type CourseTrackingHandler interface {
	Create(ctx echo.Context) error
	GetById(ctx echo.Context) error
	FindSub(ctx echo.Context) error
}

type CourseTrackingHandlerImpl struct {
	CourseTrackingService service.CourseTrackingService
}

func NewCourseTrackingHandler(service service.CourseTrackingService) CourseTrackingHandler {
	return &CourseTrackingHandlerImpl{CourseTrackingService: service}
}