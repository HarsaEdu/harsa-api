package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (profileService *ProfileServiceImpl) GetProfileByID(request *web.UserGetByIDRequest) (*domain.ProfileDetail, error) {
	err := profileService.Validator.Struct(request)
	if err != nil {
		return nil, err
	}

	result, err := profileService.ProfileRepository.GetProfileByID(request.ID)
	if err != nil {
		return nil, fmt.Errorf("profile not found")
	}

	// response := conversion.ProfileModelToResponse(result)
	return result, nil
}
