package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (questionsRoutes *QuestionsRoutesImpl) QuestionsWeb(apiGroup *echo.Group) {
	questionsGroup := apiGroup.Group("/module/quizzes/questions")

	questionsGroup.DELETE("/:id", questionsRoutes.QuestionsHandler.Delete, middleware.InstructorMiddleware)

}
