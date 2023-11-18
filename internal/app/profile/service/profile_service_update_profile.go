package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/labstack/echo/v4"
)

func (profileService *ProfileServiceImpl) UpdateProfile(ctx echo.Context, profile *domain.UserProfile, profileID uint) error {
	err := profileService.Validator.Struct(profile)
	if err != nil {
		return err
	}

	profileExists, _ := profileService.ProfileRepository.FindByUserID(profile.UserID)
	if profileExists == nil {
		return fmt.Errorf("profile not found")
	}

	if image, _ := ctx.FormFile("image"); image == nil {
		profile.ImageUrl = profileExists.ImageUrl
	} else {
		profile.ImageUrl, err = profileService.cloudinaryUploader.Uploader(ctx, "image", "profiles")
		if err != nil {
			return fmt.Errorf("error uploading image : %s", err.Error())
		}
	}

	err = profileService.ProfileRepository.UpdateProfile(profile, profileID)
	if err != nil {
		return fmt.Errorf("something wrong, cannot update profile")
	}
	return nil
}
