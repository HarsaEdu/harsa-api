package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (subsPlanService *SubsPlanServiceImpl) GetAll(search string) ([]domain.SubsPlan, *web.Pagination, error) {
	result, total, err := subsPlanService.SubsPlanRepository.GetAll(search)

	if total == 0 {
		return nil, nil, fmt.Errorf("subs plan not found")
	}

	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}

	pagination := conversion.RecordToPaginationResponse(0, 0, total)

	return result, pagination, nil
}
