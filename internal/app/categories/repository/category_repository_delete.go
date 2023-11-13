package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (categoryRepository *CategoryRepositoryImpl) Delete(id int) error {
	result := categoryRepository.DB.Where("id = ?", id).Delete(&domain.Category{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
