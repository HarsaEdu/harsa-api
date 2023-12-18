package service

import (
	"fmt"
	"regexp"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (categoryService *CategoryServiceImpl) Update(ctx echo.Context, request web.CategoryUpdateRequest, id int, isImageExist bool) error {
	var imageUrl string

	// Check if the request is valid
	err := categoryService.Validator.Struct(request)
	if err != nil {
		return err
	}

	// Check if the category exists
	existingCategoryId, _ := categoryService.CategoryRepository.FindById(id)
	if existingCategoryId == nil {
		return fmt.Errorf("category not found")
	}

	// Check if the category name already exists
	existingCategoryName, _ := categoryService.CategoryRepository.FindByName(request.Name)
	if existingCategoryName != nil && int(existingCategoryName.ID) != id {
		return fmt.Errorf("category name already exist")
	}

	if isImageExist {
		imageUrl, err = categoryService.cloudinaryUploader.Uploader(ctx, "image", "categories", true)
		if imageUrl != "" && !regexp.MustCompile(`\.png$|\.jpg$`).MatchString(imageUrl) {
			return fmt.Errorf("invalid file format")
		}
		if err != nil {
			return fmt.Errorf("error uploading image : %s", err.Error())
		}
	}

	category := conversionRequest.CategoryUpdateRequestToCategoriesModel(request)
	category.Image_url = imageUrl

	err = categoryService.CategoryRepository.Update(category, id)
	if err != nil {
		return fmt.Errorf("error when updating : %s", err.Error())
	}

	return nil
}
