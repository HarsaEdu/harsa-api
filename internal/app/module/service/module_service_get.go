package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (moduleService *ModuleServiceImpl) GetAllByCourseId(offset, limit int, search string, courseId uint) ([]web.ModuleResponse, *web.Pagination, error) {
	result, total, err := moduleService.ModuleRepository.GetAllByCourseId(offset, limit, search, courseId)
	if err != nil {
		return nil, nil, err
	
	}

	if result == nil {
		return nil, nil, fmt.Errorf("module not found")
	}

	res := conversion.ConvertAllModuleResponse(result)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return res, pagination, nil
}