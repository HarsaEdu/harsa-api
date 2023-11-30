package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (categoryService *CategoryServiceImpl) FindById(id int) (*domain.Category, error) {
	result, _ := categoryService.CategoryRepository.FindById(id)

	if result == nil {
		return nil, fmt.Errorf("category not found")
	}

	return result, nil
}
