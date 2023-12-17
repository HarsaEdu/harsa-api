package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/feedback/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type FeedbackService interface {
	CreateByUserAndCourseId(feedback web.FeedbackCreateRequest, userId uint, courseId uint) error
	UpdateByUserAndCourseId(feedback web.FeedbackUpdateRequest, userId, courseId uint) error
	GetByIdAndCourseId(courseId, id uint) (*web.FeedBackResponseForTracking, error)
	GetAllByCourseId(courseId uint, offset, limit int, search string) ([]web.FeedBackResponseForTracking, *web.Pagination, error)
	GetByIdUserAndCourseId(userId, courseId uint) (*web.FeedBackResponseForTracking, error)
	DeleteByUserAndCourseId(userId, courseId uint) error
	DeleteById(id uint) error
}

type FeedbackServiceImpl struct {
	FeedbackRepository repository.FeedbackRepository
	Validator          *validator.Validate
}

func NewFeedbackService(cr repository.FeedbackRepository, validate *validator.Validate) FeedbackService {
	return &FeedbackServiceImpl{FeedbackRepository: cr, Validator: validate}
}
