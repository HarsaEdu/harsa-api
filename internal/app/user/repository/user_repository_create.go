package repository

import (

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (userRepository *UserRepositoryImpl) UserCreate(user *domain.User) (*domain.User, error) {
	// insert user into database
	result := userRepository.DB.Create(&user)

	// check if error when user inserted
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

