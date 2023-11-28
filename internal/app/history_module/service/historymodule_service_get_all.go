package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (historyModuleService *HistoryModuleServiceImpl) GetAll(offset, limit int, search string) ([]domain.HistoryModule, *web.Pagination, error) {
	result, total, err := historyModuleService.HistoryModuleRepository.GetAll(offset, limit, search)

	if total == 0 {
		return nil, nil, fmt.Errorf("history module not found")
	}

	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return result, pagination, nil
}
