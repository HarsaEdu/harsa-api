package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (moduleRepository *ModuleRepositoryImpl) CreateModule(module *domain.Module) error {
	result := moduleRepository.DB.Create(&module)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (moduleRepository *ModuleRepositoryImpl) CreateSection(section *domain.Section) error {
	result := moduleRepository.DB.Create(&section)
	if result.Error != nil {
		return result.Error
	}

	return nil
}