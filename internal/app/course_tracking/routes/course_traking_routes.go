package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (courseTrackingRoutes *CourseTrackingRoutesImpl) CourseTrackingMobile(apiGroup *echo.Group)  {
	courseTrackingsGroup := apiGroup.Group("/course/:course-id/tracking")
	courseTrackingsGroupGet := apiGroup.Group("/course/tracking")

	courseTrackingsGroup.POST("", courseTrackingRoutes.CourseTrackingHandler.Create, middleware.StudentMiddleare)
	courseTrackingsGroupGet.GET("/:id", courseTrackingRoutes.CourseTrackingHandler.GetById, middleware.StudentMiddleare)
}