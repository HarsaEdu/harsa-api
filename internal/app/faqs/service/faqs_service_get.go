package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (faqsService *FaqsServiceImpl) GetAll(offset, limit int, search string) ([]domain.Faqs, *web.Pagination, error) {
	result, total, err := faqsService.FaqRepository.GetAll(offset, limit, search)

	if total == 0 {
		return nil, nil, fmt.Errorf("faqs not found")
	}

	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return result, pagination, nil
}

func (faqsService *FaqsServiceImpl) GetById(id int) (*domain.Faqs, error) {
	result, err := faqsService.FaqRepository.FindById(id)
	if result == nil {
		return nil, fmt.Errorf("faqs not found")
	}

	if err != nil {
		return nil, fmt.Errorf("internal Server Error")
	}

	return result, nil
}