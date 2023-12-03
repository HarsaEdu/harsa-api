package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/feedback/service"
	"github.com/labstack/echo/v4"
)

type FeedbackHandler interface {
	CreateByUserAndCourseId(ctx echo.Context) error
	UpdateByUserAndCourseId(ctx echo.Context) error
	FindById(ctx echo.Context) error
	GetAllByCourseId(ctx echo.Context) error
	GetByIdUserAndCourseId(ctx echo.Context) error
	DeleteByUserAndCourseId(ctx echo.Context) error
}

type FeedbackHandlerImpl struct {
	FeedbackService service.FeedbackService
}

func NewFeedbackHandler(service service.FeedbackService) FeedbackHandler {
	return &FeedbackHandlerImpl{FeedbackService: service}
}
