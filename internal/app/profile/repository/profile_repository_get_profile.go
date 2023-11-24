package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *ProfileRepositoryImpl) GetProfileByID(profileID uint) (*domain.UserProfile, error) {
	profile := domain.UserProfile{}
	result := repository.DB.Where("id = ?", profileID).First(&profile)
	if result.Error != nil {
		return nil, result.Error
	}
	return &profile, nil
}
