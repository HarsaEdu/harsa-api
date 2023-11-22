package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/interest/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type InterestService interface {
	CreateInterest(profileID uint, request *web.InterestRequest) error
	GetInterest(profileID uint) ([]domain.Course, error)
}

type InterestServiceImpl struct {
	InterestRepository repository.InterestRepository
	Validator          *validator.Validate
}

func NewInterestService(repository repository.InterestRepository, validator *validator.Validate) InterestService {
	return &InterestServiceImpl{
		InterestRepository: repository,
		Validator:          validator,
	}
}
