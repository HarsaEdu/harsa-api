package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/questions/service"
	"github.com/labstack/echo/v4"
)

type QuestionsHandler interface {
	Delete(ctx echo.Context) error
}

type QuestionsHandlereImpl struct {
	QuestionsService service.QuestionsService
}

func NewQuestionsHandler(service service.QuestionsService) QuestionsHandler {
	return &QuestionsHandlereImpl{
		QuestionsService: service,
	}
}
