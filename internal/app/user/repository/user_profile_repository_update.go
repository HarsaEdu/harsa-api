package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (userRepository *UserRepositoryImpl) UserProfileUpdate(userProfile *domain.UserProfile) error {

	// update user into database
	result := userRepository.DB.Where("id = ?", userProfile.ID).Updates(&userProfile)

	// check if error when user updated
	if result.Error != nil {
		return result.Error
	}

	return nil
}
