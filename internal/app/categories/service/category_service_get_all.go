package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/labstack/echo/v4"
)

func (categoryService *CategoryServiceImpl) GetAll(ctx echo.Context) ([]domain.Category, error) {
	result, err := categoryService.CategoryRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("category not found")
	}

	return result, nil
}
