package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (subsPlanRoutes *SubsPlanRoutesImpl) SubsPlanWeb(apiGroup *echo.Group) {
	subsPlanGroup := apiGroup.Group("/subs-plan")

	subsPlanGroup.POST("", subsPlanRoutes.subsPlanHandler.Create, middleware.AdminMiddleware)
	subsPlanGroup.GET("", subsPlanRoutes.subsPlanHandler.GetAll, middleware.AllUserMiddleare)
	subsPlanGroup.PUT("/:id", subsPlanRoutes.subsPlanHandler.Update, middleware.AdminMiddleware)
	subsPlanGroup.PATCH("/:id", subsPlanRoutes.subsPlanHandler.UpdateImage, middleware.AdminMiddleware)
	subsPlanGroup.DELETE("/:id", subsPlanRoutes.subsPlanHandler.Delete, middleware.AdminMiddleware)
}

func (subsPlanRoutes *SubsPlanRoutesImpl) SubsPlanMobile(apiGroup *echo.Group) {
	subsPlanGroup := apiGroup.Group("/subs-plan")

	subsPlanGroup.GET("", subsPlanRoutes.subsPlanHandler.GetAll, middleware.AllUserMiddleare)
}
