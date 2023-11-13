package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (categoryService *CategoryServiceImpl) Update(ctx echo.Context, request web.CategoryUpdateRequest, id int) error {
	// Check if the request is valid
	err := categoryService.Validator.Struct(request)
	if err != nil {
		return validation.ValidationError(ctx, err)
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

	category := conversionRequest.CategoryUpdateRequestToCategoriesModel(request)

	err = categoryService.CategoryRepository.Update(category, id)
	if err != nil {
		return fmt.Errorf("error when updating : %s", err.Error())
	}

	return nil
}
