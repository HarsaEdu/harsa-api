package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *HistorySubModuleRepositoryImpl) UpdateHistorySubModule(model *domain.HistorySubModule, id uint) error {
	result := repository.DB.Where("id = ?", id).Updates(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
