package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (subsPlanRepository *SubsPlanRepositoryImpl) UpdateImage(imagUrl *domain.SubsPlan, id int) error {
	result := subsPlanRepository.DB.Model(&domain.SubsPlan{}).Where("id=?", id).Update("image_url", imagUrl.Image_url)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
