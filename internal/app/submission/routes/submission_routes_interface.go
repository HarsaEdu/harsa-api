package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/submission/handler"
	"github.com/labstack/echo/v4"
)

type SubmissionRoutes interface {
	SubmissionWeb(apiGroup *echo.Group)
	SubmissionMobile(apiGroup *echo.Group)
}

type SubmissionRoutesImpl struct {
	submissionHandler handler.SubmissionHandler
}

func NewSubmissionRoutes(submission handler.SubmissionHandler) SubmissionRoutes {
	return &SubmissionRoutesImpl{
		submissionHandler: submission,
	}
}
