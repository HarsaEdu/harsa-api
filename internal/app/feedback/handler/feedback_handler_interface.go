package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/feedback/service"
	"github.com/labstack/echo/v4"
)

type FeedbackHandler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	FindById(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetByIdUserAndCourseId(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type FeedbackHandlerImpl struct {
	FeedbackService service.FeedbackService
}

func NewFeedbackHandler(service service.FeedbackService) FeedbackHandler {
	return &FeedbackHandlerImpl{FeedbackService: service}
}
