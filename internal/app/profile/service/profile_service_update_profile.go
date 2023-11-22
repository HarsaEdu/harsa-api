package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (profileService *ProfileServiceImpl) UpdateProfile(ctx echo.Context, request *web.UpdateProfileRequest, profileID uint) error {
	err := profileService.Validator.Struct(request)
	if err != nil {
		return err
	}

	profileExists, _ := profileService.ProfileRepository.GetProfileByID(profileID)
	if profileExists == nil {
		return fmt.Errorf("profile not found")
	}

	profile := conversion.ProfileRequestToProfileModel(profileID, request)

	profile.ImageUrl, err = profileService.cloudinaryUploader.Uploader(ctx, "image", "profiles", false)
	if err != nil {
		return fmt.Errorf("error uploading image : %s", err.Error())
	}

	if profile.ImageUrl == "" {
		profile.ImageUrl = profileExists.ImageUrl
	}

	err = profileService.ProfileRepository.UpdateProfile(profile)
	if err != nil {
		return fmt.Errorf("something wrong, cannot update profile")
	}
	return nil
}
