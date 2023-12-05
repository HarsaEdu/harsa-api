package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"gorm.io/gorm"
)

type HistoryQuizRepository interface {
	Answering(userId uint,answers []web.AnswersRequest, quizId uint) error
	GetAllHistoryQuizByQuizId(offset, limit int, quizID uint,search string) ([]domain.HistoryQuiz, int64, error)
	GetMyHistoryQuizByQuizId(userID uint, quizID uint) (*domain.HistoryQuiz, error)
	GetById(id uint) (*domain.HistoryQuiz, error)
	Cek(userId uint, quizID uint) (*domain.HistoryQuiz, error)

}

type HistoryQuizRepositoryImpl struct {
	DB *gorm.DB
}

func NewHistoryQuizRepository(db *gorm.DB) HistoryQuizRepository {
	return &HistoryQuizRepositoryImpl{
		DB: db,
	}
}