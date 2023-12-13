package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/faqs/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type FaqsService interface {
	Create(faqs web.FaqsRequest) error
	GetAll(offset, limit int, search string) ([]domain.Faqs, *web.Pagination, error)
	GetById(id int) (*domain.Faqs, error)
	Delete(id int) error
	Update(faqs web.FaqsUpdateRequest, id int) error
}

type FaqsServiceImpl struct {
	FaqRepository repository.FaqRepository
	Validator     *validator.Validate
}

func NewFaqsService(faqRepository repository.FaqRepository, validator *validator.Validate) FaqsService {
	return &FaqsServiceImpl{FaqRepository: faqRepository, Validator: validator}

}
