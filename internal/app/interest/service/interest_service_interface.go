package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/interest/repository"
	profile "github.com/HarsaEdu/harsa-api/internal/app/profile/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type InterestService interface {
	CreateInterest(profileID uint, request *web.InterestRequest) error
	GetInterestRecommendation(profileID uint) (*[]web.InterestResponse, error)
}

type InterestServiceImpl struct {
	InterestRepository repository.InterestRepository
	ProfileRepository  profile.ProfileRepository
	Validator          *validator.Validate
}

func NewInterestService(repository repository.InterestRepository, validator *validator.Validate, profileRepository profile.ProfileRepository) InterestService {
	return &InterestServiceImpl{
		InterestRepository: repository,
		ProfileRepository:  profileRepository,
		Validator:          validator,
	}
}
