package service

import (
	"fmt"
	"regexp"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (subsPlanService *SubsPlanServiceImpl) Create(ctx echo.Context, subsPlan *web.SubsPlanCreateRequest) error {
	err := subsPlanService.Validator.Struct(subsPlan)
	if err != nil {
		return err
	}

	result := conversion.SubsPlanRequestToSubsPlanDomain(subsPlan)

	imageUrl, err := subsPlanService.CloudinaryUploader.Uploader(ctx, "image", "subs-plan", false)
	if !regexp.MustCompile(`\.png$|\.jpg$`).MatchString(imageUrl) {
		return fmt.Errorf("invalid file format")
	}

	result.Image_url = imageUrl

	if err != nil {
		return fmt.Errorf("error when uploading image : %s", err.Error())
	}

	err = subsPlanService.SubsPlanRepository.Create(result)
	if err != nil {
		return fmt.Errorf("error when creating subs plan %s", err.Error())
	}
	return nil
}
