package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/quizzes/service"
	"github.com/labstack/echo/v4"
)

type QuizzesHandler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	FindById(ctx echo.Context) error
	Delete(ctx echo.Context) error
	GetAll(ctx echo.Context) error
}

type QuizzesHandlereImpl struct {
	QuizzesService service.QuizzesService
}

func NewQuizzesHandler(service service.QuizzesService) QuizzesHandler {
	return &QuizzesHandlereImpl{
		QuizzesService: service,
	}
}