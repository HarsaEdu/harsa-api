package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (userService *UserServiceImpl) UserGetAll(offset int, limit int, search string) ([]domain.UserEntity, int64, error) {
	users, total, err := userService.UserRepository.UserGetAll(offset, limit, search)
	if len(search) > 0 && total <= 0 {
		return nil, total, fmt.Errorf("users not found")
	}
	if err != nil {
		return nil, total, fmt.Errorf("internal Server Error")
	}
	return users, total, nil
}

func (userService *UserServiceImpl) GetUserDetail(userRequest web.UserGetByIDRequest) (*domain.UserDetail, error) {
	err := userService.Validate.Struct(userRequest)

	if err != nil {
		return nil, err
	}

	users, err := userService.UserRepository.GetUserByID(userRequest.ID)

	if err != nil {
		return nil, err
	}

	return users, nil
}
