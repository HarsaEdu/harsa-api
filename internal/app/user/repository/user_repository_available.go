package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (registerRepository *UserRepositoryImpl) UserAvailable(username, email string) (*domain.User, error) {
	// define domain user
	user := domain.User{}

	// get data from database
	result := registerRepository.DB.Where("username = ? OR email = ?", username, email).First(&user)

	// check if error when get data
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (registerRepository *UserRepositoryImpl) UserAvailableUsername(username string) (*domain.User, error) {
	// define domain user
	user := domain.User{}

	// get data from database
	result := registerRepository.DB.Where("username = ?", username).First(&user)

	// check if error when get data
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (registerRepository *UserRepositoryImpl) UserAvailableEmail(email string) (*domain.User, error) {
	// define domain user
	user := domain.User{}

	// get data from database
	result := registerRepository.DB.Where("email = ?", email).First(&user)

	// check if error when get data
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (registerRepository *UserRepositoryImpl) UserAvailableByID(id uint) (*domain.User, error) {
	// define domain user
	user := domain.User{}

	// get data from database
	result := registerRepository.DB.Where("id = ?", id).First(&user)

	// check if error when get data
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
