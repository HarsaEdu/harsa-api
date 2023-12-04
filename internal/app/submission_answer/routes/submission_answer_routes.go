package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerRoutes *SubmissionAnswerRoutesImpl) SubmissionAnswerWeb(apiGroup *echo.Group) {
	submissionAnswerGroup := apiGroup.Group("/submissions")

	submissionAnswerGroup.GET("/:submission-id", submissionAnswerRoutes.SubmissionAnswerHandler.Get, middleware.StudentMiddleare)
}

func (submissionAnswerRoutes *SubmissionAnswerRoutesImpl) SubmissionAnswerMobile(apiGroup *echo.Group) {
	submissionAnswerGroup := apiGroup.Group("/submissions/:idSubmission/submission-answer")

	submissionAnswerGroup.POST("", submissionAnswerRoutes.SubmissionAnswerHandler.Create, middleware.StudentMiddleare)
	submissionAnswerGroup.PATCH("/:id", submissionAnswerRoutes.SubmissionAnswerHandler.Update, middleware.StudentMiddleare)
	submissionAnswerGroup.GET("/:id", submissionAnswerRoutes.SubmissionAnswerHandler.FindById, middleware.StudentMiddleare)
}
