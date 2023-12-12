package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (subsPlanService *SubsPlanServiceImpl) FindById(id int) (*web.SubsPlanResposne, error) {
	result, err := subsPlanService.SubsPlanRepository.FindById(id)

	if result == nil {
		return nil, fmt.Errorf("subs plan not found")
	}

	if err != nil {
		return nil, fmt.Errorf("internal Server Error")
	}

	response := conversion.ConvertSubPlanResponse(result)

	return response , nil
}
