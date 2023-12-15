package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerRoutes *SubmissionAnswerRoutesImpl) SubmissionAnswerWeb(apiGroup *echo.Group) {
	submissionAnswerGroup := apiGroup.Group("/submissions/:idSubmission/submission-answer")

	submissionAnswerGroup.GET("", submissionAnswerRoutes.SubmissionAnswerHandler.Get, middleware.InstructorMiddleware)
	submissionAnswerGroup.GET("/:subsAnsId", submissionAnswerRoutes.SubmissionAnswerHandler.FindByIdWeb, middleware.InstructorMiddleware)
	submissionAnswerGroup.PATCH("/:subsAnsId", submissionAnswerRoutes.SubmissionAnswerHandler.UpdateWeb, middleware.InstructorMiddleware)

}

func (submissionAnswerRoutes *SubmissionAnswerRoutesImpl) SubmissionAnswerMobile(apiGroup *echo.Group) {
	submissionAnswerGroup := apiGroup.Group("/submissions/:idSubmission/submission-answer")

	submissionAnswerGroup.POST("", submissionAnswerRoutes.SubmissionAnswerHandler.Create, middleware.StudentMiddleare)
	submissionAnswerGroup.PATCH("/:subsAnsId", submissionAnswerRoutes.SubmissionAnswerHandler.Update, middleware.StudentMiddleare)
	submissionAnswerGroup.GET("/:subsAnsId", submissionAnswerRoutes.SubmissionAnswerHandler.FindById, middleware.StudentMiddleare)
}
