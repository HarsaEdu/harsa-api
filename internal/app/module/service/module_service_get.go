package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (moduleService *ModuleServiceImpl) GetAll(offset, limit int, search string) ([]domain.Module, *web.Pagination, error) {
	result, total, err := moduleService.ModuleRepository.GetAll(offset, limit, search)
	if err != nil {
		return nil, nil, err
	
	}

	if result == nil {
		return nil, nil, fmt.Errorf("module not found")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return result, pagination, nil
}