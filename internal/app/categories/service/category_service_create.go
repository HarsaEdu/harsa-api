package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (categoryService *CategoryServiceImpl) Create(ctx echo.Context, request web.CategoryCreateRequest) error {
	err := categoryService.Validator.Struct(request)
	if err != nil {
		return err
	}

	existingName, _ := categoryService.CategoryRepository.FindByName(request.Name)
	if existingName != nil {
		return fmt.Errorf("category name already exists")
	}

	category := conversionRequest.CategoryCreateRequestToCategoriesModel(request)

	err = categoryService.CategoryRepository.Create(category)
	if err != nil {
		return fmt.Errorf("error when creating category %s", err.Error())
	}

	return nil

}
