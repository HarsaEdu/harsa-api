package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (submissionRoutes *SubmissionRoutesImpl) SubmissionWeb(apiGroup *echo.Group) {
	submissionsGroup := apiGroup.Group("/modules/:moduleId/submissions")

	submissionsGroup.POST("", submissionRoutes.submissionHandler.Create, middleware.InstructorMiddleware)
	submissionsGroup.GET("", submissionRoutes.submissionHandler.GetAllWeb, middleware.InstructorMiddleware)
	submissionsGroup.PUT("/:id", submissionRoutes.submissionHandler.Update, middleware.InstructorMiddleware)
	submissionsGroup.DELETE("/:id", submissionRoutes.submissionHandler.Delete, middleware.InstructorMiddleware)
	submissionsGroup.GET("/:id", submissionRoutes.submissionHandler.FindById, middleware.InstructorMiddleware)
}

func (submissionRoutes *SubmissionRoutesImpl) SubmissionMobile(apiGroup *echo.Group) {
	submissionsGroup := apiGroup.Group("/modules/:moduleId/submissions")

	submissionsGroup.GET("", submissionRoutes.submissionHandler.GetAllMobile, middleware.StudentMiddleare)
	submissionsGroup.GET("/:id", submissionRoutes.submissionHandler.FindById, middleware.StudentMiddleare)

}
