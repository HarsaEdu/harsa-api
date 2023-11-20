package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	Create(feedback *domain.Feedback) error
	GetAll(offset, limit int, search string) ([]domain.Feedback, int64, error)
	GetById(id int) (*domain.Feedback, error)
	Update(id int, feedback *domain.Feedback) error
	Delete(id int) error
}

type FeedbackRepositoryImpl struct {
	DB *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &FeedbackRepositoryImpl{
		DB: db,
	}
}
