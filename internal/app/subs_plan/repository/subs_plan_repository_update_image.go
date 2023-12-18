package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (subsPlanRepository *SubsPlanRepositoryImpl) UpdateImage(imageUrl string, id int) error {
	result := subsPlanRepository.DB.Model(&domain.SubsPlan{}).Where("id=?", id).Update("image_url", imageUrl)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
