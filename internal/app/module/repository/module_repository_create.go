package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (moduleRepository ModuleRepositoryImpl) Create(module *domain.Module) error {
	result := moduleRepository.DB.Create(&module)
	if result.Error != nil {
		return result.Error
	}

	return nil
}