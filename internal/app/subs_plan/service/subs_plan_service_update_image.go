package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (subsPlanService *SubsPlanServiceImpl) UpdateImage(ctx echo.Context, subsPlan *web.SubsPlanUpdateImage, id int) error {

	ifExist, _ := subsPlanService.SubsPlanRepository.FindById(id)
	if ifExist == nil {
		return fmt.Errorf("subs plan not found")
	}

	result := conversion.SubsPlanUpdateImageToSubsPlanDomain(subsPlan)
	Image_url, err := subsPlanService.CloudinaryUploader.Uploader(ctx, "image", "subs_plan", true)
	if err != nil {
		return fmt.Errorf("error when uploading image : %s", err.Error())
	}

	if Image_url != "" {
		result.Image_url = Image_url
	}

	err = subsPlanService.Validator.Struct(result)
	if err != nil {
		return err
	}

	err = subsPlanService.SubsPlanRepository.Update(result, id)
	if err != nil {
		return fmt.Errorf("error when updating image subs plan %s", err.Error())
	}

	return nil
}
