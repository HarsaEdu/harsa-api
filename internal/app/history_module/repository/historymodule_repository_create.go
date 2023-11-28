package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (historyModuleRepository *HistoryModuleRepositoryImpl) Create(historyModule *domain.HistoryModule) error {
	result := historyModuleRepository.DB.Create(&historyModule)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
