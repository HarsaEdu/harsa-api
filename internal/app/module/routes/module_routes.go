package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (moduleRoutes *ModuleRoutesImpl) ModuleWeb(coursesGruop *echo.Group) {
	modulesGroup := coursesGruop.Group("/section")
	sectionGroup := coursesGruop.Group("/:courseId/section")

	modulesGroup.POST("/:sectionId/module", moduleRoutes.ModuleHandler.CreateModule, middleware.InstructorMiddleware)
	sectionGroup.POST("", moduleRoutes.ModuleHandler.CreateSection, middleware.InstructorMiddleware)
	sectionGroup.GET("", moduleRoutes.ModuleHandler.GetAllSection, middleware.InstructorMiddleware)
	sectionGroup.GET("/:id", moduleRoutes.ModuleHandler.GetAllModule, middleware.InstructorMiddleware)
	modulesGroup.GET("/module/:id", moduleRoutes.ModuleHandler.GetModuleById, middleware.InstructorMiddleware)
	modulesGroup.PUT("/module/:id", moduleRoutes.ModuleHandler.UpdateModule, middleware.InstructorMiddleware)
	sectionGroup.PUT("/:id", moduleRoutes.ModuleHandler.UpdateSection, middleware.InstructorMiddleware)
	modulesGroup.DELETE("/module/:id", moduleRoutes.ModuleHandler.DeleteModule, middleware.InstructorMiddleware)
	sectionGroup.DELETE("/:id", moduleRoutes.ModuleHandler.DeleteSection, middleware.InstructorMiddleware)
	modulesGroup.PUT("/module/order", moduleRoutes.ModuleHandler.UpdateModuleOrder, middleware.InstructorMiddleware)
	sectionGroup.PUT("/order", moduleRoutes.ModuleHandler.UpdateSectionOrder, middleware.InstructorMiddleware)
}

func (moduleRoutes *ModuleRoutesImpl) ModuleMobile(coursesGruop *echo.Group) {
	modulesGroup := coursesGruop.Group("/:courseId/section")

	modulesGroup.GET("", moduleRoutes.ModuleHandler.GetAllSection, middleware.StudentMiddleare)
	// modulesGroup.GET("/:id", moduleRoutes.ModuleHandler.GetById)
	// modulesGroup.PUT("/:id", moduleRoutes.ModuleHandler.Update, middleware.InstructorMiddleware)
	// modulesGroup.PATCH("/:id", moduleRoutes.ModuleHandler.UpdateImage, middleware.InstructorMiddleware)
	// modulesGroup.DELETE("/:id", moduleRoutes.ModuleHandler.Delete, middleware.InstructorMiddleware)
}