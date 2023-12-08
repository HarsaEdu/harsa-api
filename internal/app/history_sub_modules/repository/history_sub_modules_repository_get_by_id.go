package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *HistorySubModuleRepositoryImpl) GetHistorySubModuleByID(id uint) (*domain.HistorySubModule, error) {
	historySubModule := domain.HistorySubModule{}
	result := repository.DB.Where("id = ?", id).First(&historySubModule)
	if result.Error != nil {
		return nil, result.Error
	}
	return &historySubModule, nil
}
