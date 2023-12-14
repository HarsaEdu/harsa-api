package recommendations

import (
	interestRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/interest/repository"
	recommendationsHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/recommendations/handler"
	recommendationsRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/recommendations/routes"
	recommendationsServicePkg "github.com/HarsaEdu/harsa-api/internal/app/recommendations/service"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/pkg/firebase"
	"github.com/HarsaEdu/harsa-api/internal/pkg/openai"
	"github.com/HarsaEdu/harsa-api/internal/pkg/recommendations"
	"github.com/go-playground/validator"
)

func RecommendationsSetup(validate *validator.Validate, recommendationsApi recommendations.RecommendationsApi, openAi openai.OpenAi, firebase firebase.Firebase, userRepositoryPkg userRepositoryPkg.UserRepository, interestRepositoryPkg interestRepositoryPkg.InterestRepository) recommendationsRoutesPkg.RecommendationsRoutes {
	recommendationsService := recommendationsServicePkg.NewRecommendationsService(validate, recommendationsApi, openAi, firebase, userRepositoryPkg, interestRepositoryPkg)
	recommendationsHandler := recommendationsHandlerPkg.NewRecommendationsHandler(recommendationsService)
	recommendationsRoutes := recommendationsRoutesPkg.NewRecommendationsRoutes(recommendationsHandler)

	return recommendationsRoutes
}