package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (service *InterestServiceImpl) GetInterestRecommendation(profileID uint) (*[]web.InterestResponse, error) {
	interest, total, err := service.InterestRepository.GetInterestRecommendation(profileID)
	if err != nil {
		return nil, fmt.Errorf("failed to get interests, something happen")
	}

	if total == 0 {
		return nil, fmt.Errorf("no interest")
	}

	response := conversion.InterestResultToResponse(interest)

	return response, nil
}
