package service

import (
	"fmt"
	"regexp"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (categoryService *CategoryServiceImpl) UploadImage(ctx echo.Context, request *web.CategoryUploadImageRequest, id int, isIconExist, isImageExist bool) error {
	var err error
	var iconUrl string
	var imageUrl string

	existingCategoryId, _ := categoryService.CategoryRepository.FindById(id)
	if existingCategoryId == nil {
		return fmt.Errorf("category not found")
	}

	response := conversionRequest.CategoryUploadImageRequestToCategoriesModel(*request)

	if isImageExist {
		testHeader, _:= ctx.FormFile("image")
		fmt.Println(testHeader.Filename)
		imageUrl, err = categoryService.cloudinaryUploader.Uploader(ctx, "image", "categories", true)
		if imageUrl != "" && !regexp.MustCompile(`\.png$|\.jpg$`).MatchString(imageUrl) {
			return fmt.Errorf("invalid file format")
		}
		if err != nil {
			return fmt.Errorf("error uploading image : %s", err.Error())
		}
	}

	if isIconExist {
		iconUrl, err = categoryService.cloudinaryUploader.Uploader(ctx, "icon", "categories", true)
		if iconUrl != "" && !regexp.MustCompile(`\.png$|\.jpg$`).MatchString(iconUrl) {
			return fmt.Errorf("invalid file format")
		}
		if err != nil {
			return fmt.Errorf("error uploading icon : %s", err.Error())
		}
	}

	err = categoryService.Validator.Struct(response)
	if err != nil {
		return err
	}

	err = categoryService.CategoryRepository.UpdateImage(imageUrl, iconUrl, id)
	if err != nil {
		return fmt.Errorf("error when update image or icon %s", err.Error())
	}
	return nil
}
