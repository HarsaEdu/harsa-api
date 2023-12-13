package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (historyQuizRoutes *HistoryQuizRoutesImpl) HistoryQuizWeb(apiGroup *echo.Group) {
	historyQuizGroup := apiGroup.Group("/quizzes")

	historyQuizGroup.GET("/:quiz-id/history-quiz", historyQuizRoutes.HistoryQuizHandler.GetAll, middleware.InstructorMiddleware)

}

func (historyQuizRoutes *HistoryQuizRoutesImpl) HistoryQuizMobile(apiGroup *echo.Group) {
	historyQuizGroup := apiGroup.Group("/module/quizzes")

	historyQuizGroup.POST("/:quiz-id/answering", historyQuizRoutes.HistoryQuizHandler.Create, middleware.StudentMiddleare)
	historyQuizGroup.GET("/history-quiz/:id", historyQuizRoutes.HistoryQuizHandler.FindById, middleware.StudentMiddleare)
}