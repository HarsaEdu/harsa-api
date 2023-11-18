package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (profileService *ProfileServiceImpl) CreateProfile(ctx echo.Context, request *web.ProfileRequest, userID uint) error {
	err := profileService.Validator.Struct(request)
	if err != nil {
		return err
	}

	profileExists, _ := profileService.ProfileRepository.FindByUserID(userID)
	if profileExists != nil {
		return fmt.Errorf("profile already exists")
	}

	profile := conversion.ProfileRequestToProfileModel(userID, request)

	if image, _ := ctx.FormFile("image"); image == nil {
		profile.ImageUrl = ""
	} else {
		profile.ImageUrl, err = profileService.cloudinaryUploader.Uploader(ctx, "image", "profiles")
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
