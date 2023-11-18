package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (categoryService *CategoryServiceImpl) GetAll(offset, limit int, search string) ([]domain.Category, *web.Pagination, error) {
	result, total, err := categoryService.CategoryRepository.GetAll(offset, limit, search)

	if total == 0 {
		return nil, nil, fmt.Errorf("category not found")
	}

	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return result, pagination, nil
}
