package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (profileRepository *ProfileRepositoryImpl) UpdateProfile(profile *domain.UserProfile) error {
	result := profileRepository.DB.Where("user_id=?", profile.UserID).Updates(&profile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
