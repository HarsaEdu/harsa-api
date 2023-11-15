package service

import (
	"fmt"
)

func (categoryService *CategoryServiceImpl) Delete(id int) error {

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
