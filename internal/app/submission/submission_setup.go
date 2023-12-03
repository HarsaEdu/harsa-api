package submission

import (
	submissionHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/submission/handler"
	submissionRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/submission/repository"
	submissionRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/submission/routes"
	submissionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/submission/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func SubmissionSetup(db *gorm.DB, validate *validator.Validate) submissionRoutesPkg.SubmissionRoutes {
	submissionRepository := submissionRepositoryPkg.NewSubmissionRepository(db)
	submissionService := submissionServicePkg.NewSubmissionService(submissionRepository, *validate)
	submissionHandler := submissionHandlerPkg.NewSubmissionHandler(submissionService)
	submissionRoutes := submissionRoutesPkg.NewSubmissionRoutes(submissionHandler)

	return submissionRoutes
}
