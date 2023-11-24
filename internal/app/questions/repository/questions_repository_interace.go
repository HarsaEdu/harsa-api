package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type QuestionsRepository interface {
	CekIdFromQuestion(userId uint, questionId uint, role string) (*domain.Questions, error)
	Delete(question *domain.Questions) error
}

type QuestionsRepositoryImpl struct {
	DB *gorm.DB
}

func NewQuestionsRepository(db *gorm.DB) QuestionsRepository {
	return &QuestionsRepositoryImpl{
		DB: db,
	}
}