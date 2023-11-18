package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (profileRepository *ProfileRepositoryImpl) UpdateProfile(profile *domain.UserProfile, id uint) error {
	result := profileRepository.DB.Where("id=?", id).Updates(&profile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
