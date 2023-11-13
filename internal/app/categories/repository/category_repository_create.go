package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (categoryRepository *CategoryRepositoryImpl) Create(category *domain.Category) error {
	result := categoryRepository.DB.Create(&category)
	if result.Error != nil {
		return result.Error
	}

	return nil
}