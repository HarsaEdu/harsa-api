package interest

import (
	interestHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/interest/handler"
	interestRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/interest/repository"
	interestRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/interest/routes"
	interestServicePkg "github.com/HarsaEdu/harsa-api/internal/app/interest/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func InterestSetup(db *gorm.DB, validate *validator.Validate) interestRoutesPkg.InterestRoutes {
	interestRepository := interestRepositoryPkg.NewInterestRepository(db)
	interestService := interestServicePkg.NewInterestService(interestRepository, validate)
	interestHandler := interestHandlerPkg.NewInterestHandler(interestService)
	interestRoute := interestRoutesPkg.NewInterestRoutes(interestHandler)

	return interestRoute
}
