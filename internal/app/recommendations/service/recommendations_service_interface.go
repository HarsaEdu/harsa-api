package service

import (
	interestRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/interest/repository"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/firebase"
	"github.com/HarsaEdu/harsa-api/internal/pkg/openai"
	"github.com/HarsaEdu/harsa-api/internal/pkg/recommendations"
	"github.com/go-playground/validator"
)

type RecommendationsService interface {
	GetRecommendations(request *web.GetRecommendationsRequest) (*web.GetRecommendationsResponse, error)
	GetRecommendationsForInstructor() error
}

type RecommendationsServiceImpl struct {
	Validate           *validator.Validate
	RecommendationsApi recommendations.RecommendationsApi
	OpenAi             openai.OpenAi
	Firebase           firebase.Firebase
	UserRepository     userRepositoryPkg.UserRepository
	InterestRepository interestRepositoryPkg.InterestRepository
}

func NewRecommendationsService(validate *validator.Validate, recommendationsApi recommendations.RecommendationsApi, openAi openai.OpenAi, firebase firebase.Firebase, userRepositoryPkg userRepositoryPkg.UserRepository, interestRepositoryPkg interestRepositoryPkg.InterestRepository) RecommendationsService {
	return &RecommendationsServiceImpl{
		Validate:           validate,
		RecommendationsApi: recommendationsApi,
		OpenAi:             openAi,
		Firebase:           firebase,
		UserRepository:     userRepositoryPkg,
		InterestRepository: interestRepositoryPkg,
	}
}