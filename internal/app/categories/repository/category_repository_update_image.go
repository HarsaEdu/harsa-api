package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (categoryRepository *CategoryRepositoryImpl) UpdateImage(imageUrl string, icon string, id int) error {

	result := categoryRepository.DB.Table("categories").
		Where("id=?", id).
		Updates(domain.Category{
			Image_url: imageUrl,
			Icon: icon,
		})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
