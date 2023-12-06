package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/course_tracking/service"
	"github.com/labstack/echo/v4"
)

type CourseTrackingHandler interface {
	Create(ctx echo.Context) error
	GetById(ctx echo.Context) error
	FindSub(ctx echo.Context) error
	FindSubModuleByID(ctx echo.Context) error
	FindSubmissionByID(ctx echo.Context) error
	FindQuizzByID(ctx echo.Context) error
	FindModuleHistory(ctx echo.Context) error
	GetAllTracking(ctx echo.Context) error
	GetAllTrackingWeb(ctx echo.Context) error
	GetAllTrackingUserWeb(ctx echo.Context) error
	DeleteEnrolled(ctx echo.Context) error
	CreateWeb(ctx echo.Context) error
	GetByUserIdAndCourseID(ctx echo.Context) error
}

type CourseTrackingHandlerImpl struct {
	CourseTrackingService service.CourseTrackingService
}

func NewCourseTrackingHandler(service service.CourseTrackingService) CourseTrackingHandler {
	return &CourseTrackingHandlerImpl{CourseTrackingService: service}
}
