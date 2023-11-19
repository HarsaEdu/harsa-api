package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (quizzesRoutes *QuizzesRoutesImpl) Quizzes(apiGroup *echo.Group) {
	quizzesGroup := apiGroup.Group("/quizzes")

	quizzesGroup.POST("", quizzesRoutes.QuizzesHandler.Create, middleware.InstructorMiddleware)
	quizzesGroup.PUT("/:id", quizzesRoutes.QuizzesHandler.Update, middleware.InstructorMiddleware)
	quizzesGroup.GET("/:id", quizzesRoutes.QuizzesHandler.FindById, middleware.AllUserMiddleare)
	quizzesGroup.DELETE("/:id", quizzesRoutes.QuizzesHandler.Delete, middleware.InstructorMiddleware)
	quizzesGroup.GET("/module/:id", quizzesRoutes.QuizzesHandler.GetAll, middleware.AllUserMiddleare)

}