package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/quizzes/handler"
	"github.com/labstack/echo/v4"
)

type QuizzesRoutes interface {
	Quizzes(apiGroup *echo.Group)
}

type QuizzesRoutesImpl struct {
	Echo            *echo.Echo
	QuizzesHandler   handler.QuizzesHandler
}

func NewQuizzesRoutes(e *echo.Echo, quizzesHandler handler.QuizzesHandler) QuizzesRoutes {
	return &QuizzesRoutesImpl{
		Echo:            e,
		QuizzesHandler: quizzesHandler,
	}
}