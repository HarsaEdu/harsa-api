package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (profileRepository ProfileRepositoryImpl) CreateProfile(profile *domain.UserProfile) error {
	//insert user profile to database
	result := profileRepository.DB.Create(&profile)

	// check if error when user inserted
	if result.Error != nil {
		return result.Error
	}

	return nil
}
