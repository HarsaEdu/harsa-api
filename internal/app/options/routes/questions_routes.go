package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (optionsRoutes *OptionsRoutesImpl) OptionsWeb(apiGroup *echo.Group) {
	optionsGroup := apiGroup.Group("/module/quizzes/questions/options")

	optionsGroup.DELETE("/:id", optionsRoutes.OptionsHandler.Delete, middleware.InstructorMiddleware)

}
