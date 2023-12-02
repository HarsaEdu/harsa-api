package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerRoutes *SubmissionAnswerRoutesImpl) SubmissionAnswerWeb(apiGroup *echo.Group) {
	submissionAnswerGroup := apiGroup.Group("/submissions/:idSubmission/submission-answer")

	submissionAnswerGroup.POST("", submissionAnswerRoutes.SubmissionAnswerHandler.Create, middleware.StudentMiddleare)
}
