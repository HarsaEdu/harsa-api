package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/feedback/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type FeedbackService interface {
	Create(feedback web.FeedbackCreateRequest, userId uint) error
	Update(feedback web.FeedbackUpdateRequest, id int) error
	FindById(id int) (*domain.Feedback, error)
	GetAll(offset, limit int, search string) ([]domain.Feedback, *web.Pagination, error)
	Delete(id int) error
}

type FeedbackServiceImpl struct {
	FeedbackRepository repository.FeedbackRepository
	Validator          *validator.Validate
}

func NewFeedbackService(cr repository.FeedbackRepository, validate *validator.Validate) FeedbackService {
	return &FeedbackServiceImpl{FeedbackRepository: cr, Validator: validate}
}
