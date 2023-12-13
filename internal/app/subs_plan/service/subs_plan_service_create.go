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

	file, _ := ctx.FormFile("image")

	
	if file != nil {
		imageUrl, err := subsPlanService.CloudinaryUploader.Uploader(ctx, "image", "subs-plan", false)
		if err != nil {
			return fmt.Errorf("error when uploading image : %s", err.Error())
		}
		if !regexp.MustCompile(`\.png$|\.jpg$`).MatchString(imageUrl) {
			return fmt.Errorf("invalid file format")
		}
		result.Image_url = imageUrl
	}


	err = subsPlanService.SubsPlanRepository.Create(result)
	if err != nil {
		return fmt.Errorf("error when creating subs plan %s", err.Error())
	}
	return nil
}

func (subsPlanService *SubsPlanServiceImpl) CreateFromExisting(ctx echo.Context, request *web.SubsPlanUpdateRequest, id uint) error {
	err := subsPlanService.Validator.Struct(request)
	if err != nil {
		return err
	}

	existingPlan, _ := subsPlanService.SubsPlanRepository.FindById(int(id))
	if existingPlan == nil {
		return fmt.Errorf("subs plan not found")
	}

	newSubsPlan := conversion.ExistingSubsPlanToSubsPlanDomain(request, existingPlan)

	file, _ := ctx.FormFile("image")

	
	if file != nil {
		imageUrl, err := subsPlanService.CloudinaryUploader.Uploader(ctx, "image", "subs-plan", false)
		if err != nil {
			return fmt.Errorf("error when uploading image : %s", err.Error())
		}
		if !regexp.MustCompile(`\.png$|\.jpg$`).MatchString(imageUrl) {
			return fmt.Errorf("invalid file format")
		}
		newSubsPlan.Image_url = imageUrl
	}

	err = subsPlanService.SubsPlanRepository.Create(newSubsPlan)
	if err != nil {
		return fmt.Errorf("error when creating subs plan %s", err.Error())
	}

	err = subsPlanService.SubsPlanRepository.UpdateStatus(false, id)
	if err != nil {
		return fmt.Errorf("error when updating subs plan status %s", err.Error())
	}

	return nil
}