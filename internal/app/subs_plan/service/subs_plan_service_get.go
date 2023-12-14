package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (subsPlanService *SubsPlanServiceImpl) GetAll(search string) ([]web.SubsPlanResposne, *web.Pagination, error) {
	result, total, err := subsPlanService.SubsPlanRepository.GetAllActive(search)

	if total == 0 {
		return nil, nil, fmt.Errorf("subs plan not found")
	}

	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}

	response := conversion.ConvertAllSubPlanResponse(result)

	pagination := conversion.RecordToPaginationResponse(0, 0, total)

	return response, pagination, nil
}
