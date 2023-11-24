package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (profileRepository *ProfileRepositoryImpl) FindByUserID(userID uint) (*domain.UserProfile, error) {
	profile := domain.UserProfile{}

	result := profileRepository.DB.Where("user_id = ?", userID).First(&profile)
	if result.Error != nil {
		return nil, result.Error
	}

	return &profile, nil
}
