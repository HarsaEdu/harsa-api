package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (profileRepository *ProfileRepositoryImpl) FindByUserID(userID uint) (*domain.ProfileDetailMobile, error) {
	user := &domain.ProfileDetailMobile{}

	result := profileRepository.DB.Model(&domain.User{}).Select("users.id as user_id, user_profiles.id as user_profile_id, roles.id as role_id, email, username, phone_number, roles.name as role_name, first_name, last_name, bio, address, city, gender, job, date_birth, image_url, subscriptions.id as subscription_id, subscriptions.end_date as end_date, subscriptions.start_date as start_date").
		Joins("left join user_profiles on user_profiles.user_id = users.id").
		Joins("left join subscriptions on subscriptions.user_id = users.id").
		Joins("left join roles on roles.id = users.role_id").
		Where("users.id = ?", userID).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (profileRepository *ProfileRepositoryImpl) IsExists(userID uint) bool {
	var user = &domain.UserProfile{}
	if err := profileRepository.DB.Model(&domain.UserProfile{}).Where("user_id = ? ", userID).
    Find(&user ).
    Error; err != nil || user.ID != 0 {
    	return true
    }
	return false
}
