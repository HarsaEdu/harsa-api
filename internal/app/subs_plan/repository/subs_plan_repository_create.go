package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (SubsPlanRepository *SubsPlanRepositoryImpl) Create(subsPlan *domain.SubsPlan) error {
	result := SubsPlanRepository.DB.Create(&subsPlan)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
