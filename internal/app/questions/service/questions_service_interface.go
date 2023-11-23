package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/questions/repository"
	"github.com/go-playground/validator"
)

type QuestionsService interface {
	Delete(userId uint, questionId uint, role string) error
}

type QuestionsServiceImpl struct {
	QuestionsRepository  repository.QuestionsRepository
	Validator            *validator.Validate
}

func NewQuestionsService(qr repository.QuestionsRepository, validate *validator.Validate) QuestionsService {
	return &QuestionsServiceImpl{
		QuestionsRepository: qr, 
		Validator: validate, 
	}
}