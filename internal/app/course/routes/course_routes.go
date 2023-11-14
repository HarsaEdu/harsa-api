package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (courseRoutes *CourseRoutesImpl) Course(apiGroup *echo.Group) {
	coursesGroup := apiGroup.Group("/courses")

	coursesGroup.POST("", courseRoutes.CourseHandler.Create, middleware.InstructorMiddleware)
}