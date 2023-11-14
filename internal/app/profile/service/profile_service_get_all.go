package service

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (profileService *ProfileServiceImpl) GetAllProfiles() ([]domain.Profile, error) {
	result, err := profileService.ProfileRepository.GetAllProfiles()
	if err != nil {
		return nil, err
	}
	return result, nil
}
