package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (categoryRepository *CategoryRepositoryImpl) UpdateImage(imageUrl string, icon string, id int) error {

	updateField := map[string]any{}

	if imageUrl != "" {
		updateField["image_url"] = imageUrl
	}

	if icon != "" {
		updateField["icon"] = icon
	}

	result := categoryRepository.DB.Model(&domain.Category{}).
		Where("id=?", id).
		Updates(updateField)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
