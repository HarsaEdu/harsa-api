package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (userRepository *UserRepositoryImpl) UserGetAll(offset, limit int, search string) ([]domain.UserEntity, int64, error) {
	var users []domain.UserEntity
	var total int64

	query := userRepository.DB.Model(&domain.User{}).Select("users.id as id, email, username, phone_number, roles.name as role_name, first_name, last_name").
		Joins("left join user_profiles on user_profiles.user_id = users.id").
		Joins("left join roles on roles.id = users.role_id")

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("users.username LIKE ? OR users.email LIKE ? OR user_profiles.first_name LIKE ? OR user_profiles.last_name LIKE ?", s, s, s, s)
	}

	query.Find(&users).Count(&total)

	query = query.Offset(offset).Limit(limit)

	result := query.Find(&users)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, total, nil
}

func (userRepository *UserRepositoryImpl) GetUserByID(userID uint) (*domain.UserDetail, error) {
	user := &domain.UserDetail{}

	result := userRepository.DB.Model(&domain.User{}).Select("users.id as user_id, user_profiles.id as user_profile_id, roles.id as role_id, email, username, phone_number, roles.name as role_name, first_name, last_name, bio, address, city, gender, job, date_birth").
		Joins("left join user_profiles on user_profiles.user_id = users.id").
		Joins("left join roles on roles.id = users.role_id").
		Where("users.id = ?", userID).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
