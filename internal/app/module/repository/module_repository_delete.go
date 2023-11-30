package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (moduleRepository *ModuleRepositoryImpl) Delete(module *domain.Module) error {

	if err := moduleRepository.DB.Delete(&module).Error; err != nil {
		return  err
	}

	return nil
}