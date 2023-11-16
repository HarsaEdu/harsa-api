package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (categoryRepository *CategoryRepositoryImpl) FindByName(name string) (*domain.Category, error) {
	category := domain.Category{}

	result := categoryRepository.DB.Where("name=?", name).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}
