package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (profileRepository *ProfileRepositoryImpl) FindByUserID(userID uint) (*domain.Profile, error) {
	profile := domain.Profile{}

	result := profileRepository.DB.Where("user_id = ?", userID).First(&profile)
	if result.Error != nil {
		return nil, result.Error
	}

	return &profile, nil
}
