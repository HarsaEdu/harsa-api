package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/submission_answer/service"
	"github.com/labstack/echo/v4"
)

type SubmissionAnswerHandler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	UpdateWeb(ctx echo.Context) error
	FindById(ctx echo.Context) error
	Get(ctx echo.Context) error
}

type SubmissionAnswerHandlerImpl struct {
	SubmissionAnswerservice service.SubmissionAnswerService
}

func NewSubmissionAnswer(service service.SubmissionAnswerService) SubmissionAnswerHandler {
	return &SubmissionAnswerHandlerImpl{SubmissionAnswerservice: service}
}
