package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_quiz/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type HistoryQuizService interface {
	GetById(id uint) (*web.HistoryQuizResponseWeb, error)
	GetAllByQuizWeb(quizId uint, offset, limit int, search string) ([]web.HistoryQuizResponseWeb, *web.Pagination, error)
	Create(request []web.AnswersRequest, userID uint, quizID uint) error
}

type HistoryQuizServiceImpl struct {
	HistoryQuizRepository  repository.HistoryQuizRepository
	Validator          *validator.Validate
}

func NewHistoryQuizService(qr repository.HistoryQuizRepository, validate *validator.Validate) HistoryQuizService {
	return &HistoryQuizServiceImpl{
		HistoryQuizRepository: qr, 
		Validator: validate, 
	}
}