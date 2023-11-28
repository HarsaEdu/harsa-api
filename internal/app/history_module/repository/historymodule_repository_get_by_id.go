package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (historyModuleRepository *HistoryModuleRepositoryImpl) GetById(id int) (*domain.HistoryModule, error) {
	historyModule := &domain.HistoryModule{}

	result := historyModuleRepository.DB.First(&historyModule, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return historyModule, nil
}
