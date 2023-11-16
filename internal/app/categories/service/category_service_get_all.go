package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (categoryService *CategoryServiceImpl) GetAll() ([]domain.Category, error) {
	result, err := categoryService.CategoryRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("category not found")
	}

	return result, nil
}
