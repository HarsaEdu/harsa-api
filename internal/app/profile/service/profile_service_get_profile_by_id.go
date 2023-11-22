package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (profileService *ProfileServiceImpl) GetProfileByID(profileID uint) (*web.GetProfileResponse, error) {
	result, err := profileService.ProfileRepository.GetProfileByID(profileID)
	if err != nil {
		return nil, fmt.Errorf("profile not found")
	}

	response := conversion.ProfileModelToResponse(result)
	return response, nil
}
