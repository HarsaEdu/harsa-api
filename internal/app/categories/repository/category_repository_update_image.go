package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (categoryRepository *CategoryRepositoryImpl) UpdateImage(imageUrl *domain.Category, id int) error {
	result := categoryRepository.DB.Model(&domain.Category{}).Where("id=?", id).Update("image_url", imageUrl.Image_url)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
