package repository

import (
	"fmt"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (userRepository *UserRepositoryImpl) UserDelete(id uint) error {
	// insert user into database
	result := userRepository.DB.Model(&domain.User{}).Where("id = ?", id).Update("deleted_at", time.Now())

	// check if error when user inserted
	fmt.Println(result.Error)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
