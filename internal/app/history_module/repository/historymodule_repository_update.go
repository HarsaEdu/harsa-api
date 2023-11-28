package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (historyModuleRepository *HistoryModuleRepositoryImpl) Update(id int, historyModule *domain.HistoryModule) error {
	result := historyModuleRepository.DB.Where("id = ?", id).Updates(&domain.HistoryModule{UserID: historyModule.UserID})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
