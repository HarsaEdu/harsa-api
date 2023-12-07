package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (faqsRoutes *FaqsRoutesImpl) FaqsWeb(apiGroup *echo.Group) {
	categoriesGroup := apiGroup.Group("/faqs")

	categoriesGroup.POST("", faqsRoutes.FaqsHandler.Create, middleware.AdminMiddleware)
	categoriesGroup.GET("", faqsRoutes.FaqsHandler.GetAll)
	categoriesGroup.GET("/:id", faqsRoutes.FaqsHandler.GetById)
	categoriesGroup.DELETE("/:id", faqsRoutes.FaqsHandler.Delete, middleware.AdminMiddleware)
	categoriesGroup.PUT("/:id", faqsRoutes.FaqsHandler.Update, middleware.AdminMiddleware)

}

func (faqsRoutes *FaqsRoutesImpl) FaqsMobile(apiGroup *echo.Group) {
	categoriesGroup := apiGroup.Group("/faqs")

	categoriesGroup.GET("", faqsRoutes.FaqsHandler.GetAll, middleware.StudentMiddleare)

}
