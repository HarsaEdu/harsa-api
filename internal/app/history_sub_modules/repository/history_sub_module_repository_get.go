package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *HistorySubModuleRepositoryImpl) GetHistorySubModuleByUserID(userID uint) ([]domain.HistorySubModule, int64, error) {
	models := []domain.HistorySubModule{}
	result := repository.DB.Preload("SubModule").Where("user_id = ?", userID).
		Find(&models)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64
	result.Count(&total)
	return models, total, nil
}
