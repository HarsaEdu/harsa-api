package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/quizzes/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type QuizzesService interface {
	Create(request web.QuizRequest, role string) error
	Update(request web.QuizRequest, quizId uint, role string) error
	FindById(quizId uint) (*web.QuizResponse, error)
	Delete(userId uint, quizId uint, role string) error
	GetAll(moduleId uint, offset int, limit int, search string) ([]web.GetAllQuizResponse, *web.Pagination, error)
}

type QuizzesServiceImpl struct {
	QuizzesRepository  repository.QuizzesRepository
	Validator          *validator.Validate
}

func NewQuizzesService(qr repository.QuizzesRepository, validate *validator.Validate) QuizzesService {
	return &QuizzesServiceImpl{
		QuizzesRepository: qr, 
		Validator: validate, 
	}
}