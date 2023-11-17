package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/labstack/echo/v4"
)

func (profileService *ProfileServiceImpl) UpdateProfile(ctx echo.Context, profile *domain.Profile, profileID uint) error {
	err := profileService.Validator.Struct(profile)
	if err != nil {
		return err
	}

	profileExists, _ := profileService.ProfileRepository.FindByUserID(profileID)
	if profileExists == nil {
		return fmt.Errorf("profile not found")
	}
	//fix code below getting image from header
	if file, _ := ctx.FormFile("image"); file != nil {
		profile.ImageURL, err = profileService.cloudinaryUploader.Uploader(ctx, "image", "profiles")
		if err != nil {
			return fmt.Errorf("error uploading image : %s", err.Error())
		}
	} else {
		profile.ImageURL = profileExists.ImageURL
	}

	err = profileService.ProfileRepository.UpdateProfile(profile, profileID)
	if err != nil {
		return fmt.Errorf("something wrong, cannot update profile")
	}
	return nil
}
