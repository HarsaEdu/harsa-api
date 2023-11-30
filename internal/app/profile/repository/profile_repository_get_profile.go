package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (profileRepository *ProfileRepositoryImpl) GetProfileByID(profileID uint) (*domain.ProfileDetail, error) {
	user := &domain.ProfileDetail{}

	result := profileRepository.DB.Model(&domain.User{}).Select("users.id as user_id, user_profiles.id as user_profile_id, roles.id as role_id, email, username, phone_number, roles.name as role_name, first_name, last_name, bio, address, city, gender, job, date_birth, image_url").
		Joins("left join user_profiles on user_profiles.user_id = users.id").
		Joins("left join roles on roles.id = users.role_id").
		Where("user_profiles.id = ?", profileID).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
