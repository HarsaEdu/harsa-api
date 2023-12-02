package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/feedback/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type FeedbackService interface {
	CreateByUserAndCourseId(feedback web.FeedbackCreateRequest, userId uint, courseId uint) error
	UpdateByUserAndCourseId(feedback web.FeedbackUpdateRequest, userId, courseId uint) error
	FindById(id int) (*domain.Feedback, error)
	GetAll(offset, limit int, search string) ([]domain.Feedback, *web.Pagination, error)
	GetByIdUserAndCourseId(userId, courseId uint) (*domain.Feedback, error)
	DeleteByUserAndCourseId(userId, courseId uint) error
}

type FeedbackServiceImpl struct {
	FeedbackRepository repository.FeedbackRepository
	Validator          *validator.Validate
}

func NewFeedbackService(cr repository.FeedbackRepository, validate *validator.Validate) FeedbackService {
	return &FeedbackServiceImpl{FeedbackRepository: cr, Validator: validate}
}
