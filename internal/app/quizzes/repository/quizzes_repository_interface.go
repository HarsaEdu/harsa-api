package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type QuizzesRepository interface {
	Create(newQuiz *domain.Quizzes) error
	Update(UpdateQuiz *domain.Quizzes, quizExist *domain.Quizzes) error
	CekId(userId uint, quizId uint, role string) (*domain.Quizzes, error)
	FindById(quizId uint) (*domain.Quizzes, error)
}

type QuizzesRepositoryImpl struct {
	DB *gorm.DB
}

func NewQuizzesRepository(db *gorm.DB) QuizzesRepository {
	return &QuizzesRepositoryImpl{
		DB: db,
	}
}