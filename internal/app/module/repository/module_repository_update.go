package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (moduleRepository *ModuleRepositoryImpl) Update(updateModul *domain.Module, moduleExist *domain.Module) error {

	tx := moduleRepository.DB.Begin()

	moduleExist.Title = updateModul.Title
	moduleExist.Description = updateModul.Description
    moduleExist.Type = updateModul.Type
	moduleExist.Order = updateModul.Order
	moduleExist.SubModules = updateModul.SubModules

	if err := tx.Save(&moduleExist).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, updatedSubModules := range updateModul.SubModules {
        if err := tx.Save(&updatedSubModules).Error; err != nil {
            tx.Rollback()
            return err
        }
    }

    tx.Commit()

	return nil
}