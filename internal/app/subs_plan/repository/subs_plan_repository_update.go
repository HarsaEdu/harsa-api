package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (SubsPlanRepository *SubsPlanRepositoryImpl) UpdateStatus(isActive bool, id uint) error {
	result := SubsPlanRepository.DB.Model(domain.SubsPlan{}).Where("id=?", id).Update("is_active", isActive)
	if result.Error != nil {
		return result.Error
	}
	return nil

}