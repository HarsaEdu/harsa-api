package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (categoryService *CategoryServiceImpl) UploadImage(ctx echo.Context, request *web.CategoryUploadImageRequest, id int) error {
	var err error
	existingCategoryId, _ := categoryService.CategoryRepository.FindById(id)
	if existingCategoryId == nil {
		return fmt.Errorf("category not found")
	}

	response := conversionRequest.CategoryUploadImageRequestToCategoriesModel(*request)
	response.Image_url, err = categoryService.cloudinaryUploader.Uploader(ctx, "file", "categories", true)
	if err != nil {
		return fmt.Errorf("error uploading image : %s", err.Error())
	}

	err = categoryService.CategoryRepository.UpdateImage(response, id)
	if err != nil {
		return fmt.Errorf("error when updating : %s", err.Error())
	}

	err = categoryService.Validator.Struct(response)
	if err != nil {
		return err
	}

	return nil
}
