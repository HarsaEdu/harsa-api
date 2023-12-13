package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	Create(feedback *domain.Feedback) error
	GetAllByCourseId(courseId uint, offset, limit int, search string) ([]domain.Feedback, int64, error)
	GetByIdAndCourseId(courseId, id uint) (*domain.Feedback, error)
	GetByIdUserAndCourseId(userId, courseId uint) (*domain.Feedback, error)
	UpdateByUserAndCourseId(userId, courseId uint, feedback *domain.Feedback) error
	DeleteByUserAndCourseId(userId, courseId uint) error
	AutoUpdateRating(courseId uint) error
	Cek(userId uint, courseID uint) (*domain.Feedback, error)
}

type FeedbackRepositoryImpl struct {
	DB *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &FeedbackRepositoryImpl{
		DB: db,
	}
}
