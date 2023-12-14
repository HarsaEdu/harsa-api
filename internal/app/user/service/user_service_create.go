package service

import (
	"fmt"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (userService *UserServiceImpl) UserCreate(ctx echo.Context, userRequest web.UserCreateRequest) error {
	// validate the request
	err := userService.Validate.Struct(userRequest)

	// check errors when validate the request
	if err != nil {
		return err
	}

	// check available username and email
	existingUser, _ := userService.UserRepository.UserAvailable(userRequest.Username, userRequest.Email)
	if existingUser != nil {
		return fmt.Errorf("already exists")
	}

	// convert request to user model
	userAccount := conversionRequest.UserCreateRequestToUserModel(userRequest)

	// hash password
	userAccount.Password, err = userService.Password.HashPassword(userAccount.Password)
	if err != nil {
		return err
	}

	userService.UserRepository.HandleTrx(ctx, func(repo repository.UserRepository) error {
		// insert data and get back user data with id and role name
		res, err := userService.UserRepository.UserCreate(userAccount)

		// check if error when insert data
		if err != nil {
			return fmt.Errorf("error when creating user %s:", err.Error())
		}
		// convert birth date
		birthDate, err := time.Parse("2006-01-02", userRequest.DateBirth)

		// convert request to user model
		userProfile := conversionRequest.UserCreateRequestToUserProfileModel(userRequest, res.ID, birthDate)

		err = userService.UserRepository.UserProfileCreate(userProfile)

		if err != nil {
			return fmt.Errorf("error when creating user profile %s:", err.Error())
		}
		return nil
	})

	return nil
}
