package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (authRepository *AuthRepositoryImpl) RegisterUser(user *domain.User) (*domain.Auth, error) {
	// define domain auth
	auth := &domain.Auth{}

	// insert user into database
	result := authRepository.DB.Create(&user)

	// check if error when user inserted
	if result.Error != nil {
		return nil, result.Error
	}

	// get user data from database
	authRepository.DB.Model(&domain.User{}).
		Select("users.id as id, username, email, role.name as role_name").
		Joins("left join roles on roles.id as role_id").
		First(&auth)

	// check if error when get user data from database
	if result.Error != nil {
		return nil, result.Error
	}

	return auth, nil
}
