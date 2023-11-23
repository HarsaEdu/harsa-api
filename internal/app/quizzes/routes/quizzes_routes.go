package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (quizzesRoutes *QuizzesRoutesImpl) QuizzesWeb(apiGroup *echo.Group) {
	quizzesGroup := apiGroup.Group("/module/quizzes")
	quizzesModule := apiGroup.Group("/module/:module-id/quizzes")

	quizzesModule.POST("", quizzesRoutes.QuizzesHandler.Create, middleware.InstructorMiddleware)
	quizzesGroup.PUT("/:id", quizzesRoutes.QuizzesHandler.Update, middleware.InstructorMiddleware)
	quizzesGroup.GET("/:id", quizzesRoutes.QuizzesHandler.FindById, middleware.AllUserMiddleare)
	quizzesGroup.DELETE("/:id", quizzesRoutes.QuizzesHandler.Delete, middleware.InstructorMiddleware)
	quizzesModule.GET("", quizzesRoutes.QuizzesHandler.GetAll, middleware.AllUserMiddleare)

}

func (quizzesRoutes *QuizzesRoutesImpl) QuizzesMobile(apiGroup *echo.Group) {
	quizzesGroup := apiGroup.Group("/module/quizzes")
	quizzesModule := apiGroup.Group("/module/:module-id/quizzes")

	quizzesGroup.GET("/:id", quizzesRoutes.QuizzesHandler.FindById, middleware.StudentMiddleare)
	quizzesModule.GET("", quizzesRoutes.QuizzesHandler.GetAll, middleware.StudentMiddleare)

}