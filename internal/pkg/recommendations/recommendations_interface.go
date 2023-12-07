package recommendations

import (
	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

type RecommendationsApi interface {
	GetRecommendations(request *web.GetRecommendationsRequest) (*web.GetRecommendationsResponse, error)
}

type recommendationsApiImpl struct {
	RecommendationsApiConfig *configs.RecommendationsApiConfig
}

func NewRecommendationsApi(recommendationsApiConfig *configs.RecommendationsApiConfig) RecommendationsApi {
	return &recommendationsApiImpl{RecommendationsApiConfig: recommendationsApiConfig}
}