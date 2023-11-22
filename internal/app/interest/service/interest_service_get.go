package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (service *InterestServiceImpl) GetInterestRecommendation(profileID uint) ([]domain.Course, error) {
	result, total, err := service.InterestRepository.GetInterestRecommendation(profileID)
	if err != nil {
		return nil, fmt.Errorf("failed to get interests, something happen")
	}

	if total == 0 {
		return nil, fmt.Errorf("no interest")
	}
	return result, nil
}
