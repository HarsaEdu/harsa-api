package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (subsPlanService *SubsPlanServiceImpl) FindById(id int) (*domain.SubsPlan, error) {
	result, err := subsPlanService.SubsPlanRepository.FindById(id)

	if result == nil {
		return nil, fmt.Errorf("subs plan not found")
	}

	if err != nil {
		return nil, fmt.Errorf("internal Server Error")
	}

	return result, nil
}
