package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (profileRepository *ProfileRepositoryImpl) GetAllProfiles() ([]domain.Profile, error) {
	profiles := []domain.Profile{}

	result := profileRepository.DB.Find(&profiles)
	if result.Error != nil {
		return nil, result.Error
	}

	return profiles, nil
}
