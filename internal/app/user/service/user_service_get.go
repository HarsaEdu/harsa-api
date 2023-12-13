package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (userService *UserServiceImpl) UserGetAll(offset int, limit int, search string, roleId int) ([]domain.UserEntity, *web.Pagination, error) {
	users, total, err := userService.UserRepository.UserGetAll(offset, limit, search, roleId)
	if len(search) > 0 && total <= 0 {
		return nil, nil, fmt.Errorf("users not found")
	}
	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}
	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	return users, pagination, nil
}

func (userService *UserServiceImpl) UserGetAllStudentSubscribe(offset int, limit int, search string, courseId uint) ([]domain.UserEntity, *web.Pagination, error) {
	users, total, err := userService.UserRepository.UserGetAllStudentSubscribe(offset, limit, search, courseId)
	if len(search) > 0 && total <= 0 {
		return nil, nil, fmt.Errorf("users not found")
	}
	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}
	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	return users, pagination, nil
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


func (userService *UserServiceImpl) GetUserAccount(userID uint) (*web.UserAccountResponse, error) {

	users, err := userService.UserRepository.GetUserAccountByID(userID)

	if err != nil {
		return nil, err
	}
	result := conversion.ConvertUserAccountResponse(users)
	return result, nil
}
