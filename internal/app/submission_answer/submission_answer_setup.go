package submissionanswer

import (
	submissionAnswerHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/submission_answer/handler"
	submissionAnswerRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/submission_answer/repository"
	submissionAnswerRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/submission_answer/routes"
	submissionAnswerServicePkg "github.com/HarsaEdu/harsa-api/internal/app/submission_answer/service"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func SubmissionAnswerSetup(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader) submissionAnswerRoutesPkg.SubmissionAnswerRoutes {
	submissionAnswerRepository := submissionAnswerRepositoryPkg.NewSubmissionAnswer(db)
	submissionAnswerService := submissionAnswerServicePkg.NewSubmissionAnswer(submissionAnswerRepository, validate, cloudinary)
	submissionAnswerHandler := submissionAnswerHandlerPkg.NewSubmissionAnswer(submissionAnswerService)
	submissionAnswerRoutes := submissionAnswerRoutesPkg.NewSubmissionAnswerRoutes(submissionAnswerHandler)

	return submissionAnswerRoutes
}
