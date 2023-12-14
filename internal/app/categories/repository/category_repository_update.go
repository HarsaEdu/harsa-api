package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (categoryRepository *CategoryRepositoryImpl) Update(category *domain.Category, id int) error {

	result := categoryRepository.DB.Where("id=?", id).Updates(&domain.Category{ID: category.ID, Name: category.Name, Description: category.Description, Image_url: category.Image_url})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
