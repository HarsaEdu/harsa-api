package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (moduleRepository *ModuleRepositoryImpl) DeleteModule(module *domain.Module) error {

	if err := moduleRepository.DB.Delete(&module).Error; err != nil {
		return  err
	}

	return nil
}

func (moduleRepository *ModuleRepositoryImpl) DeleteSection(section *domain.Section) error {

	if err := moduleRepository.DB.Delete(&section).Error; err != nil {
		return  err
	}

	return nil
}