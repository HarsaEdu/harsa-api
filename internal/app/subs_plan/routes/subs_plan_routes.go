package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (subsPlanRoutes *SubsPlanRoutesImpl) SubsPlan(apiGroup *echo.Group) {
	subsPlanGroup := apiGroup.Group("/subs-plan")

	subsPlanGroup.POST("", subsPlanRoutes.subsPlanHandler.Create, middleware.AdminMiddleware)
	subsPlanGroup.GET("", subsPlanRoutes.subsPlanHandler.GetAll, middleware.AllUserMiddleare)
	subsPlanGroup.GET("/mobile", subsPlanRoutes.subsPlanHandler.GetAll, middleware.StudentMiddleare)
	subsPlanGroup.PUT("/:id", subsPlanRoutes.subsPlanHandler.Update, middleware.AdminMiddleware)
	subsPlanGroup.PATCH("/:id", subsPlanRoutes.subsPlanHandler.UpdateImage, middleware.AdminMiddleware)
	subsPlanGroup.DELETE("/:id", subsPlanRoutes.subsPlanHandler.Delete, middleware.AdminMiddleware)
}
