package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (service *InterestServiceImpl) GetInterestRecommendation(userID uint) (*[]web.InterestResponse, error) {
	profile, err := service.ProfileRepository.FindByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("profile not found")
	}
	interest, total, err := service.InterestRepository.GetInterestRecommendation(profile.UserProfileID)
	if err != nil {
		return nil, fmt.Errorf("failed to get interests, something happen")
	}
	fmt.Print("profile id :", profile.UserProfileID)

	if total == 0 {
		return nil, fmt.Errorf("no interest")
	}

	response := conversion.InterestResultToResponse(interest)

	return response, nil
}
