package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (categoryService *CategoryServiceImpl) Delete(ctx echo.Context, id int) error {

	IfExist, _ := categoryService.CategoryRepository.FindById(id)
	if IfExist == nil {
		return fmt.Errorf("category not found")
	}
	err := categoryService.CategoryRepository.Delete(id)

	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}
