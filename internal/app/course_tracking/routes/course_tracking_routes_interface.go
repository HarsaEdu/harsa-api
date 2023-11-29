package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/course_tracking/handler"
	"github.com/labstack/echo/v4"
)

type CourseTrackingRoutes interface {
	CourseTrackingMobile(apiGroup *echo.Group)
}

type CourseTrackingRoutesImpl struct {
	CourseTrackingHandler handler.CourseTrackingHandler
}

func NewCourseTrackingRoutes(categoryHandler handler.CourseTrackingHandler) CourseTrackingRoutes {
	return &CourseTrackingRoutesImpl{
		CourseTrackingHandler: categoryHandler,
	}
}