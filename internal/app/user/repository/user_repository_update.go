package repository

import (

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (userRepository *UserRepositoryImpl) UserUpdate(user *domain.User) error {
	// insert user into database
	result := userRepository.DB.Model(&domain.User{}).Where("id=?", user.ID).Updates(domain.User{Email: user.Email, Username: user.Username, Password: user.Password, RoleID: user.RoleID})
	// check if error when user inserted
	if result.Error != nil {
		return result.Error
	}

	return nil
}
