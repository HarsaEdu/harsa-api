package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/questions/handler"
	"github.com/labstack/echo/v4"
)

type QuestionsRoutes interface {
	QuestionsWeb(apiGroup *echo.Group)
}

type QuestionsRoutesImpl struct {
	QuestionsHandler   handler.QuestionsHandler
}

func NewQuestionsRoutes(questionsHandler handler.QuestionsHandler) QuestionsRoutes {
	return &QuestionsRoutesImpl{
		QuestionsHandler: questionsHandler,
	}
}