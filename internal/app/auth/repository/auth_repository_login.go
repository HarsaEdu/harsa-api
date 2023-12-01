package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (authRepository *AuthRepositoryImpl) LoginUser(id uint) (*domain.Auth, error) {
	// define domain auth
	auth := &domain.Auth{}

	// get user data from database
	result := authRepository.DB.Model(&domain.User{}).
		Select("users.id as id, username, email, roles.name as role_name, users.created_at as created_at").
		Joins("left join roles on roles.id = users.role_id").
		Where("users.id = ?", id).
		First(&auth)

	// check if error when get user data from database
	if result.Error != nil {
		return nil, result.Error
	}

	return auth, nil
}
