package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type QuizzesRepository interface {
	Create(newQuiz *domain.Quizzes) error
	Update(UpdateQuiz *domain.Quizzes, quizExist *domain.Quizzes) error
	CekIdFromQuiz(userId uint, quizId uint, role string) (*domain.Quizzes, error)
	CekIdFromModule(userId uint, moduleId uint, role string) error
	FindById(quizId uint) (*domain.Quizzes, error)
	Delete(quiz *domain.Quizzes) error
	GetAll(moduleId uint, offset, limit int, search string) ([]domain.Quizzes, int64, error)
}

type QuizzesRepositoryImpl struct {
	DB *gorm.DB
}

func NewQuizzesRepository(db *gorm.DB) QuizzesRepository {
	return &QuizzesRepositoryImpl{
		DB: db,
	}
}