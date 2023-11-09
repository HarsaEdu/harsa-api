package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (r *UserRepositoryImpl) UserAvailable(username, email string) (*domain.User, error) {
	// define domain user
	user := domain.User{}

	// get data from database
	result := r.DB.Where("username = ? OR email = ?", username, email).First(&user)

	// check if error when get data
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
