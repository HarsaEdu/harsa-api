package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_quiz/handler"
	"github.com/labstack/echo/v4"
)

type HistoryQuizRoutes interface {
	HistoryQuizWeb(apiGroup *echo.Group)
	HistoryQuizMobile(apiGroup *echo.Group)
}

type HistoryQuizRoutesImpl struct {
	HistoryQuizHandler   handler.HistoryQuizHandler
}

func NewHistoryQuizRoutes(quizzesHandler handler.HistoryQuizHandler) HistoryQuizRoutes {
	return &HistoryQuizRoutesImpl{
		HistoryQuizHandler: quizzesHandler,
	}
}