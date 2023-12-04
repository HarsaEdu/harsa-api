package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/submission_answer/handler"
	"github.com/labstack/echo/v4"
)

type SubmissionAnswerRoutes interface {
	SubmissionAnswerWeb(apiGroup *echo.Group)
	SubmissionAnswerMobile(apiGroup *echo.Group)
}

type SubmissionAnswerRoutesImpl struct {
	SubmissionAnswerHandler handler.SubmissionAnswerHandler
}

func NewSubmissionAnswerRoutes(submissionAnswerHandler handler.SubmissionAnswerHandler) SubmissionAnswerRoutes {
	return &SubmissionAnswerRoutesImpl{
		SubmissionAnswerHandler: submissionAnswerHandler,
	}
}
