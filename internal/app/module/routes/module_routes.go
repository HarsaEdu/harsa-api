package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (moduleRoutes *ModuleRoutesImpl) ModuleWeb(coursesGruop *echo.Group) {
	modulesGroup := coursesGruop.Group("/:courseId/modules")

	modulesGroup.POST("", moduleRoutes.ModuleHandler.Create, middleware.InstructorMiddleware)
	modulesGroup.GET("", moduleRoutes.ModuleHandler.GetAll, middleware.InstructorMiddleware)
	// modulesGroup.GET("/:id", moduleRoutes.ModuleHandler.GetById)
	// modulesGroup.PUT("/:id", moduleRoutes.ModuleHandler.Update, middleware.InstructorMiddleware)
	// modulesGroup.PATCH("/:id", moduleRoutes.ModuleHandler.UpdateImage, middleware.InstructorMiddleware)
	// modulesGroup.DELETE("/:id", moduleRoutes.ModuleHandler.Delete, middleware.InstructorMiddleware)
}

func (moduleRoutes *ModuleRoutesImpl) ModuleMobile(coursesGruop *echo.Group) {
	modulesGroup := coursesGruop.Group("/:courseId/modules")

	modulesGroup.POST("", moduleRoutes.ModuleHandler.Create, middleware.InstructorMiddleware)
	modulesGroup.GET("", moduleRoutes.ModuleHandler.GetAll, middleware.StudentMiddleare)
	// modulesGroup.GET("/:id", moduleRoutes.ModuleHandler.GetById)
	// modulesGroup.PUT("/:id", moduleRoutes.ModuleHandler.Update, middleware.InstructorMiddleware)
	// modulesGroup.PATCH("/:id", moduleRoutes.ModuleHandler.UpdateImage, middleware.InstructorMiddleware)
	// modulesGroup.DELETE("/:id", moduleRoutes.ModuleHandler.Delete, middleware.InstructorMiddleware)
}