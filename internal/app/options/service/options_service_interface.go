package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/options/repository"
	"github.com/go-playground/validator"
)

type OptionsService interface {
	Delete(userId uint, optionId uint, role string) error 
}

type OptionsServiceImpl struct {
	OptionsRepository  repository.OptionsRepository
	Validator            *validator.Validate
}

func NewOptionsService(qr repository.OptionsRepository, validate *validator.Validate) OptionsService {
	return &OptionsServiceImpl{
		OptionsRepository: qr, 
		Validator: validate, 
	}
}