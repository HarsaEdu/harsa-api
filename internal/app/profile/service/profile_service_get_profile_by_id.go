package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (profileService *ProfileServiceImpl) GetProfileByID(id uint) (*web.GetProfileResponse, error) {
	result, err := profileService.ProfileRepository.FindByUserID(id)
	if err != nil {
		return nil, fmt.Errorf("profile not found")
	}

	response := conversion.ProfileModelToResponse(result)
	return response, nil
}
