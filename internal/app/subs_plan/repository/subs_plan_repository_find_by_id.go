package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (SubsPlanRepository *SubsPlanRepositoryImpl) FindById(id int) (*domain.SubsPlan, error) {

	subsPlan := &domain.SubsPlan{}

	result := SubsPlanRepository.DB.Where("id = ?", id).First(subsPlan)
	if result.Error != nil {
		return nil, result.Error
	}

	return subsPlan, nil
}
