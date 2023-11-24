package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (feedbackRoutes *FeedbackRoutesImpl) FeedbackWeb(apiGroup *echo.Group) {
	feedbackGroup := apiGroup.Group("/feedback")

	feedbackGroup.POST("", feedbackRoutes.FeedbackHandler.Create, middleware.StudentMiddleare)
	feedbackGroup.GET("", feedbackRoutes.FeedbackHandler.GetAll)
	feedbackGroup.GET("/:id", feedbackRoutes.FeedbackHandler.FindById)
	feedbackGroup.PUT("/:id", feedbackRoutes.FeedbackHandler.Update, middleware.StudentMiddleare)
	feedbackGroup.DELETE("/:id", feedbackRoutes.FeedbackHandler.Delete, middleware.StudentMiddleare)

}

func (feedbackRoutes *FeedbackRoutesImpl) FeedbackMobile(apiGroup *echo.Group) {
	feedbackGroup := apiGroup.Group("/feedback")

	feedbackGroup.POST("", feedbackRoutes.FeedbackHandler.Create, middleware.StudentMiddleare)
	feedbackGroup.GET("", feedbackRoutes.FeedbackHandler.GetAll)
	feedbackGroup.GET("/:id", feedbackRoutes.FeedbackHandler.FindById)
	feedbackGroup.PUT("/:id", feedbackRoutes.FeedbackHandler.Update, middleware.StudentMiddleare)
	feedbackGroup.DELETE("/:id", feedbackRoutes.FeedbackHandler.Delete, middleware.StudentMiddleare)

}
