package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/labstack/echo/v4"
)

func (profileService *ProfileServiceImpl) CreateProfile(ctx echo.Context, profile *domain.Profile) error {
	fmt.Println(profile)
	err := profileService.Validator.Struct(profile)
	if err != nil {
		return err
	}
	fmt.Println("image__url", profile.ImageURL)

	_, errProfileExists := profileService.ProfileRepository.FindByUserID(profile.UserID)
	if errProfileExists == nil {
		return fmt.Errorf("profile already exists")
	}

	profile.ImageURL, err = profileService.cloudinaryUploader.Uploader(ctx, "image", "profiles")
	if err != nil {
		return fmt.Errorf("error uploading image : %s", err.Error())
	}

	fmt.Println("image__url", profile.ImageURL)

	err = profileService.ProfileRepository.CreateProfile(profile)
	if err != nil {
		return err
	}

	return nil
}
