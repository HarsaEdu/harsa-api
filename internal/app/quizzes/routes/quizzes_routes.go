package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (quizzesRoutes *QuizzesRoutesImpl) Quizzes(apiGroup *echo.Group) {
	quizzesGroup := apiGroup.Group("/quizzes")

	quizzesGroup.POST("", quizzesRoutes.QuizzesHandler.Create, middleware.InstructorMiddleware)

}