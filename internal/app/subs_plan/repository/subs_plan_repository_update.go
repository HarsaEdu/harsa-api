package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (SubsPlanRepository *SubsPlanRepositoryImpl) Update(subsPlan *domain.SubsPlan, id int) error {
	result := SubsPlanRepository.DB.Where("id=?", id).Updates(&domain.SubsPlan{Title: subsPlan.Title, Description: subsPlan.Description, Price: subsPlan.Price, Image_url: subsPlan.Image_url, Duration_days: subsPlan.Duration_days})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
