package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (courseTrackingRoutes *CourseTrackingRoutesImpl) CourseTrackingMobile(apiGroup *echo.Group)  {
	courseTrackingsGroup := apiGroup.Group("/users/course/:course-id")
	courseTrackingsGroupGet := apiGroup.Group("/users/course")

	courseTrackingsGroup.POST("/enrolled", courseTrackingRoutes.CourseTrackingHandler.Create, middleware.StudentMiddleare)
	courseTrackingsGroupGet.GET("/tracking/:id", courseTrackingRoutes.CourseTrackingHandler.GetById, middleware.StudentMiddleare)
	courseTrackingsGroupGet.GET("/module/:module-id", courseTrackingRoutes.CourseTrackingHandler.FindSub, middleware.StudentMiddleare)
	
}