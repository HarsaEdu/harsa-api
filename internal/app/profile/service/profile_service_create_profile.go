package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (profileService *ProfileServiceImpl) CreateProfile(ctx echo.Context, request *web.CreateProfileRequest, userID uint) error {
	err := profileService.Validator.Struct(request)
	if err != nil {
		return err
	}

	profileExists, _ := profileService.ProfileRepository.FindByUserID(userID)
	if profileExists != nil {
		return fmt.Errorf("profile already exists")
	}

	profile := conversion.ProfileCreateRequestToModel(userID, request)

	profile.ImageUrl, err = profileService.cloudinaryUploader.Uploader(ctx, "image", "profiles", false)
	if err != nil {
		return fmt.Errorf("error uploading image : %s", err.Error())
	}

	err = profileService.ProfileRepository.CreateProfile(profile)
	if err != nil {
		return fmt.Errorf("something wrong, cannot create profile")
	}

	return nil
}