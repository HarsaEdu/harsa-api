package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (profileService *ProfileServiceImpl) GetProfileByID(id uint) (*domain.Profile, error) {
	result, err := profileService.ProfileRepository.FindByUserID(id)
	if err != nil {
		return nil, fmt.Errorf("profile not found")
	}
	return result, nil
}
