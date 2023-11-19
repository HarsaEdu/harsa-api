package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/course/handler"
	"github.com/labstack/echo/v4"
)

type CourseRoutes interface {
	Course(apiGroup *echo.Group) *echo.Group
}

type CourseRoutesImpl struct {
	CourseHandler handler.CourseHandler
}

func NewCourseRoutes(CourseHandler handler.CourseHandler) CourseRoutes {
	return &CourseRoutesImpl{
		CourseHandler: CourseHandler,
	}
}