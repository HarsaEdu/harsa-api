package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (subsPlanRoutes *SubmissionRoutesImpl) SubmissionWeb(apiGroup *echo.Group) {
	submissionsGroup := apiGroup.Group("/modules/:moduleId/submissions")

	submissionsGroup.POST("", subsPlanRoutes.submissionHandler.Create, middleware.InstructorMiddleware)
	submissionsGroup.GET("", subsPlanRoutes.submissionHandler.GetAllWeb, middleware.InstructorMiddleware)
	submissionsGroup.PUT("/:id", subsPlanRoutes.submissionHandler.Update, middleware.InstructorMiddleware)
	submissionsGroup.DELETE("/:id", subsPlanRoutes.submissionHandler.Delete, middleware.InstructorMiddleware)
}

func (subsPlanRoutes *SubmissionRoutesImpl) SubmissionMobile(apiGroup *echo.Group) {
	submissionsGroup := apiGroup.Group("/modules/:moduleId/submissions")

	submissionsGroup.GET("", subsPlanRoutes.submissionHandler.GetAllMobile, middleware.StudentMiddleare)

}
