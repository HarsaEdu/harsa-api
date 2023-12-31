package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (feedbackRoutes *FeedbackRoutesImpl) FeedbackWeb(apiGroup *echo.Group) {
	feedbacksGroup := apiGroup.Group("/:courseId/feedbacks")

	feedbacksGroup.GET("", feedbackRoutes.FeedbackHandler.GetAllByCourseId)
	feedbacksGroup.GET("/:id", feedbackRoutes.FeedbackHandler.GetByIdAndCourseId)
	feedbacksGroup.DELETE("/:id", feedbackRoutes.FeedbackHandler.DeleteById)
	

}

func (feedbackRoutes *FeedbackRoutesImpl) FeedbackMobile(apiGroup *echo.Group) {
	feedbacksGroup := apiGroup.Group("/:courseId/feedbacks")

	feedbacksGroup.POST("", feedbackRoutes.FeedbackHandler.CreateByUserAndCourseId, middleware.StudentMiddleare)
	feedbacksGroup.GET("/my-feedback", feedbackRoutes.FeedbackHandler.GetByIdUserAndCourseId, middleware.StudentMiddleare)
	feedbacksGroup.GET("", feedbackRoutes.FeedbackHandler.GetAllByCourseId, middleware.StudentMiddleare)
	feedbacksGroup.PUT("", feedbackRoutes.FeedbackHandler.UpdateByUserAndCourseId, middleware.StudentMiddleare)
	feedbacksGroup.DELETE("", feedbackRoutes.FeedbackHandler.DeleteByUserAndCourseId, middleware.StudentMiddleare)
}
