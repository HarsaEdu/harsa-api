package service

import (
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/recommendations"
	"github.com/go-playground/validator"
)

type RecommendationsService interface {
	GetRecommendations(request *web.GetRecommendationsRequest) (*web.GetRecommendationsResponse, error)
}

type RecommendationsServiceImpl struct {
	Validate           *validator.Validate
	RecommendationsApi recommendations.RecommendationsApi
	UserRepository     userRepositoryPkg.UserRepository
}

func NewRecommendationsService(validate *validator.Validate, recommendationsApi recommendations.RecommendationsApi, userRepositoryPkg userRepositoryPkg.UserRepository) RecommendationsService {
	return &RecommendationsServiceImpl{
		Validate:           validate,
		RecommendationsApi: recommendationsApi,
		UserRepository:     userRepositoryPkg,
	}
}