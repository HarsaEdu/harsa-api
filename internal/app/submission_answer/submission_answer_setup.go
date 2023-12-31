package submissionanswer

import (
	submissionRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/submission/repository"
	submissionAnswerHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/submission_answer/handler"
	submissionAnswerRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/submission_answer/repository"
	subscriptionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
	submissionAnswerRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/submission_answer/routes"
	submissionAnswerServicePkg "github.com/HarsaEdu/harsa-api/internal/app/submission_answer/service"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func SubmissionAnswerSetup(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader, submissionRepo submissionRepositoryPkg.SubmissionRepository, subcriptionService subscriptionServicePkg.SubscriptionService) submissionAnswerRoutesPkg.SubmissionAnswerRoutes {
	submissionAnswerRepository := submissionAnswerRepositoryPkg.NewSubmissionAnswer(db)
	submissionAnswerService := submissionAnswerServicePkg.NewSubmissionAnswer(submissionAnswerRepository, validate, cloudinary, submissionRepo)
	submissionAnswerHandler := submissionAnswerHandlerPkg.NewSubmissionAnswer(submissionAnswerService, subcriptionService)
	submissionAnswerRoutes := submissionAnswerRoutesPkg.NewSubmissionAnswerRoutes(submissionAnswerHandler)

	return submissionAnswerRoutes
}
