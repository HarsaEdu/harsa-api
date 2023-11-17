package service

import (
	"fmt"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (userService *UserServiceImpl) UserProfileUpdate(userRequest web.UserProfileUpdateRequest) error {
	// validate the request
	err := userService.Validate.Struct(userRequest)

	// check errors when validate the request
	if err != nil {
		return err
	}

	// check available username and email
	existingUser, _ := userService.UserRepository.UserProfileAvailableByID(userRequest.ID)
	if existingUser == nil {
		return fmt.Errorf("user profile not found")
	}

	// convert birth date
	birthDate, err := time.Parse("2006-01-02", userRequest.DateBirth)

	if err != nil {
		return fmt.Errorf("failed to parse birth date")
	}

	// convert request to user model
	userAccount := conversionRequest.UserProfileUpdateRequestToUserModel(userRequest, birthDate)
	fmt.Println(userAccount)

	// update data
	err = userService.UserRepository.UserProfileUpdate(userAccount)

	// check if error when update data
	if err != nil {
		return fmt.Errorf("error when update user profile %s:", err.Error())
	}

	return nil
}
