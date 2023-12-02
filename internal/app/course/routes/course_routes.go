package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (courseRoutes *CourseRoutesImpl) CourseWeb(apiGroup *echo.Group) *echo.Group {
	coursesGroup := apiGroup.Group("/courses")
	dashboardsGroup := apiGroup.Group("/dashboard")

	coursesGroup.POST("", courseRoutes.CourseHandler.Create, middleware.InstructorMiddleware)
	dashboardsGroup.GET("/instructur", courseRoutes.CourseHandler.GetAllByUserId, middleware.InstructorMiddleware)
	coursesGroup.GET("", courseRoutes.CourseHandler.GetAll, middleware.AdminMiddleware)
	coursesGroup.GET("/:id", courseRoutes.CourseHandler.GetById, middleware.InstructorMiddleware)
	coursesGroup.PUT("/:id", courseRoutes.CourseHandler.Update, middleware.InstructorMiddleware)
	coursesGroup.PATCH("/:id", courseRoutes.CourseHandler.UpdateImage, middleware.InstructorMiddleware)
	coursesGroup.DELETE("/:id", courseRoutes.CourseHandler.Delete, middleware.InstructorMiddleware)

	return coursesGroup
}

func (courseRoutes *CourseRoutesImpl) CourseMobile(apiGroup *echo.Group) *echo.Group {
	coursesGroup := apiGroup.Group("/courses")

	coursesGroup.GET("", courseRoutes.CourseHandler.GetAll)
	coursesGroup.GET("/:id", courseRoutes.CourseHandler.GetById)

	return coursesGroup
}