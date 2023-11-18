package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (userRepository *UserRepositoryImpl) UserProfileCreate(userProfile *domain.UserProfile) error {
	// insert user into database
	result := userRepository.DB.Create(&userProfile)

	// check if error when user inserted
	if result.Error != nil {
		return result.Error
	}

	return nil
}
