package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_quiz/service"
	"github.com/labstack/echo/v4"
)

type HistoryQuizHandler interface {
	FindById(ctx echo.Context) error
	Create(ctx echo.Context) error
	GetAll(ctx echo.Context) error
}

type HistoryQuizHandlereImpl struct {
	HistoryQuizService service.HistoryQuizService
}

func NewHistoryQuizHandler(service service.HistoryQuizService) HistoryQuizHandler {
	return &HistoryQuizHandlereImpl{
		HistoryQuizService: service,
	}
}