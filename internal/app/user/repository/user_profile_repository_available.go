package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (userRepository *UserRepositoryImpl) UserProfileAvailableByID(id uint) (*domain.UserProfile, error) {
	// define domain user
	userProfile := domain.UserProfile{}

	// get data from database
	result := userRepository.DB.Where("id = ?", id).First(&userProfile)

	// check if error when get data
	if result.Error != nil {
		return nil, result.Error
	}

	return &userProfile, nil
}
