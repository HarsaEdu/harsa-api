package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (categoryRepository *CategoryRepositoryImpl) FindById(id int) (*domain.Category, error) {
	category := &domain.Category{}

	result := categoryRepository.DB.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}
