package service

import (
	"fmt"
	"regexp"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (categoryService *CategoryServiceImpl) Create(ctx echo.Context, request web.CategoryCreateRequest) error {
	
	var err error

	err = categoryService.Validator.Struct(request)
	if err != nil {
		return err
	}

	existingName, _ := categoryService.CategoryRepository.FindByName(request.Name)
	if existingName != nil {
		return fmt.Errorf("category name already exist")
	}

	category := conversionRequest.CategoryCreateRequestToCategoriesModel(request)
	imageUrl, err := categoryService.cloudinaryUploader.Uploader(ctx, "image", "categories", false)

	if imageUrl != "" && !regexp.MustCompile(`\.png$|\.jpg$`).MatchString(imageUrl) {
		return fmt.Errorf("invalid file format")
	}

	category.Image_url = imageUrl

	if err != nil {
		return fmt.Errorf("error uploading image : %s", err.Error())
	}

	iconUrl, err := categoryService.cloudinaryUploader.Uploader(ctx, "icon", "categories", false)
	if iconUrl != "" && !regexp.MustCompile(`\.png$|\.jpg$`).MatchString(iconUrl) {
		return fmt.Errorf("invalid file format")
	}

	category.Icon = iconUrl

	if err != nil {
		return fmt.Errorf("error uploading icon : %s", err.Error())
	}


	err = categoryService.CategoryRepository.Create(category)
	if err != nil {
		return fmt.Errorf("error when creating category %s", err.Error())
	}

	return nil

}
