package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/labstack/echo/v4"
)

func (profileService *ProfileServiceImpl) CreateProfile(ctx echo.Context, profile *domain.Profile) error {
	err := profileService.Validator.Struct(profile)
	if err != nil {
		return err
	}

	profileExists, _ := profileService.ProfileRepository.FindByUserID(profile.UserID)
	if profileExists != nil {
		return fmt.Errorf("profile already exists")
	}

	if image, _ := ctx.FormFile("image"); image == nil {
		profile.ImageURL = ""
	} else {
		profile.ImageURL, err = profileService.cloudinaryUploader.Uploader(ctx, "image", "profiles")
		if err != nil {
			return fmt.Errorf("error uploading image : %s", err.Error())
		}
	}

	err = profileService.ProfileRepository.CreateProfile(profile)
	if err != nil {
		return err
	}

	return nil
}
