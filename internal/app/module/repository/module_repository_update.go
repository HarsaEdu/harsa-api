package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (moduleRepository *ModuleRepositoryImpl) UpdateModule(updateModul *domain.Module, moduleExist *domain.Module) error {

	tx := moduleRepository.DB.Begin()

	moduleExist.Title = updateModul.Title
	moduleExist.Description = updateModul.Description
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

func (moduleRepository *ModuleRepositoryImpl) UpdateOrderModule(order int, moduleExist *domain.Module) error {

	result := moduleRepository.DB.Where("id = ?", moduleExist.ID).Updates(&domain.Module{OrderBy: order})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (moduleRepository *ModuleRepositoryImpl) UpdateOrderSection(order int, sectionExist *domain.Section) error {

	result := moduleRepository.DB.Where("id = ?", sectionExist.ID).Updates(&domain.Module{OrderBy: order})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (moduleRepository *ModuleRepositoryImpl) UpdateSection(UpdateSection *domain.Section, sectionExist *domain.Section) error {

	sectionExist.Title = UpdateSection.Title

	if err := moduleRepository.DB.Save(&sectionExist).Error; err != nil {
		return err
	}

	return nil
}