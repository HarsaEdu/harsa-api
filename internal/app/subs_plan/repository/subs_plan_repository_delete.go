package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (SubsPlanRepository *SubsPlanRepositoryImpl) Delete(id int) error {
	result := SubsPlanRepository.DB.Where("id = ?", id).Delete(&domain.SubsPlan{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
