package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repository *HistorySubModuleRepositoryImpl) CreateHistorySubModule(request *domain.HistorySubModule) error {
	result := repository.DB.Create(&request)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
