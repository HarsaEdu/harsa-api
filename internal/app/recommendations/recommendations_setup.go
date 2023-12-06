package recommendations

import (
	recommendationsHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/recommendations/handler"
	recommendationsRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/recommendations/routes"
	recommendationsServicePkg "github.com/HarsaEdu/harsa-api/internal/app/recommendations/service"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/pkg/recommendations"
	"github.com/go-playground/validator"
)

func RecommendationsSetup(validate *validator.Validate, recommendationsApi recommendations.RecommendationsApi, userRepository userRepositoryPkg.UserRepository) recommendationsRoutesPkg.RecommendationsRoutes {
	recommendationsService := recommendationsServicePkg.NewRecommendationsService(validate, recommendationsApi, userRepository)
	recommendationsHandler := recommendationsHandlerPkg.NewRecommendationsHandler(recommendationsService)
	recommendationsRoutes := recommendationsRoutesPkg.NewRecommendationsRoutes(recommendationsHandler)

	return recommendationsRoutes
}