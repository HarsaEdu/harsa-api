package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (categoryRepository *CategoryRepositoryImpl) GetAll() ([]domain.Category, error) {
	category := []domain.Category{}

	result := categoryRepository.DB.Find(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}
