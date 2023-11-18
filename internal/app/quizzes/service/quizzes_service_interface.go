package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/quizzes/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type QuizzesService interface {
	Create(ctx echo.Context, request web.QuizRequest) error
	Update(ctx echo.Context, request web.QuizRequest, quizId uint, role string) error
	FindById(ctx echo.Context, quizId uint) (*web.QuizResponse, error)
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