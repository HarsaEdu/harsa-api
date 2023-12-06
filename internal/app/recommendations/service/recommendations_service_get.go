package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (recommendationsService *RecommendationsServiceImpl) GetRecommendations(request *web.GetRecommendationsRequest) (*web.GetRecommendationsResponse, error) {
	err := recommendationsService.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	existingUser, err := recommendationsService.UserRepository.UserAvailableByID(request.UserId)
	if err != nil {
		return nil, fmt.Errorf("error when get user : %s", err.Error())
	}
	
	if existingUser == nil {
		return nil, fmt.Errorf("error when get user : user not found")
	}


	response, err := recommendationsService.RecommendationsApi.GetRecommendations(request)
	if err != nil {
		return nil, fmt.Errorf("error when get recommendations : %s", err.Error())
	}

	return response, nil
}