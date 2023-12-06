package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (courseTrackingRoutes *CourseTrackingRoutesImpl) CourseTrackingMobile(apiGroup *echo.Group) {
	courseTrackingsGroup := apiGroup.Group("/users/course/:course-id")
	courseTrackingsGroupGet := apiGroup.Group("/users/course")
	moduleTrackingGroup := apiGroup.Group("/users/course/module/:module-id")

	courseTrackingsGroup.POST("/enrolled", courseTrackingRoutes.CourseTrackingHandler.Create, middleware.StudentMiddleare)
	courseTrackingsGroupGet.GET("/trackings", courseTrackingRoutes.CourseTrackingHandler.GetAllTracking, middleware.StudentMiddleare)
	courseTrackingsGroup.GET("", courseTrackingRoutes.CourseTrackingHandler.GetByUserIdAndCourseID, middleware.StudentMiddleare)
	courseTrackingsGroupGet.GET("/module/:module-id", courseTrackingRoutes.CourseTrackingHandler.FindSub, middleware.StudentMiddleare)

	moduleTrackingGroup.GET("/tracking", courseTrackingRoutes.CourseTrackingHandler.FindModuleHistory, middleware.StudentMiddleare)
	moduleTrackingGroup.GET("/sub-module/:sub-module-id", courseTrackingRoutes.CourseTrackingHandler.FindSubModuleByID, middleware.StudentMiddleare)
	moduleTrackingGroup.GET("/submission/:submission-id", courseTrackingRoutes.CourseTrackingHandler.FindSubmissionByID, middleware.StudentMiddleare)
	moduleTrackingGroup.GET("/quizz/:quizz-id", courseTrackingRoutes.CourseTrackingHandler.FindQuizzByID, middleware.StudentMiddleare)
}
	

func (courseTrackingRoutes *CourseTrackingRoutesImpl) CourseTrackingWeb(apiGroup *echo.Group)  {
	profileTrackingsGroup := apiGroup.Group("/profile/tracking/user")
	userTrackingGroup     := apiGroup.Group("/course/:course-id/user")
	
	profileTrackingsGroup.GET("/:id", courseTrackingRoutes.CourseTrackingHandler.GetAllTrackingWeb, middleware.InstructorMiddleware)
	userTrackingGroup.GET("", courseTrackingRoutes.CourseTrackingHandler.GetAllTrackingUserWeb, middleware.InstructorMiddleware)
	userTrackingGroup.DELETE("/tracking/:tracking-id", courseTrackingRoutes.CourseTrackingHandler.DeleteEnrolled, middleware.InstructorMiddleware)
	userTrackingGroup.POST("/:user-id", courseTrackingRoutes.CourseTrackingHandler.Create, middleware.InstructorMiddleware)
	
}


