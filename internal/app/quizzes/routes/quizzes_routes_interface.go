package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/quizzes/handler"
	"github.com/labstack/echo/v4"
)

type QuizzesRoutes interface {
	QuizzesWeb(apiGroup *echo.Group)
	QuizzesMobile(apiGroup *echo.Group)
}

type QuizzesRoutesImpl struct {
	QuizzesHandler   handler.QuizzesHandler
}

func NewQuizzesRoutes(quizzesHandler handler.QuizzesHandler) QuizzesRoutes {
	return &QuizzesRoutesImpl{
		QuizzesHandler: quizzesHandler,
	}
}