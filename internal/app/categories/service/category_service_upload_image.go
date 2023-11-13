package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (categoryService *CategoryServiceImpl) UploadImage(ctx echo.Context, request *web.CategoryUploadImageRequest, id int) error {
	err := categoryService.Validator.Struct(request)
	if err != nil {
		return validation.ValidationError(ctx, err)
	}

	existingCategoryId, _ := categoryService.CategoryRepository.FindById(id)
	if existingCategoryId == nil {
		return fmt.Errorf("category not found")
	}

	response := conversionRequest.CategoryUploadImageRequestToCategoriesModel(*request)
	response.Image_url = categoryService.cloudinaryUploader.Uploader(ctx, "file", "categories")

	err = categoryService.CategoryRepository.UpdateImage(response, id)
	if err != nil {
		return fmt.Errorf("error when updating : %s", err.Error())
	}

	return nil
}
