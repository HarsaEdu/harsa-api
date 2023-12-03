package service

import (
	"fmt"
)

func (moduleService *ModuleServiceImpl) DeleteModule(moduleId uint, userId uint, role string) error {

	existingModule, err := moduleService.ModuleRepository.CekIdFromModule(userId, moduleId, role)
	if err != nil {
		return fmt.Errorf("error when cek id user from course :%s", err.Error())
	}

	err = moduleService.ModuleRepository.DeleteModule(existingModule)
	if err != nil {
		return fmt.Errorf("error when update module : %s", err.Error())
	}

	return nil
}

func (moduleService *ModuleServiceImpl) DeleteSection(sectionId uint, userId uint, role string) error {

	existingSection, err := moduleService.ModuleRepository.CekIdFromSection(userId, sectionId, role)
	if err != nil {
		return fmt.Errorf("error when cek id user from section :%s", err.Error())
	}

	err = moduleService.ModuleRepository.DeleteSection(existingSection)
	if err != nil {
		return fmt.Errorf("error when update section : %s", err.Error())
	}

	return nil
}


func (moduleService *ModuleServiceImpl) DeleteSubModule(subModuleId uint, userId uint, role string) error {

	existingSubmodule, err := moduleService.ModuleRepository.CekIdFromSubModule(userId, subModuleId, role)
	if err != nil {
		return fmt.Errorf("error when cek id user from course :%s", err.Error())
	}

	err = moduleService.ModuleRepository.DeleteSubModule(existingSubmodule)
	if err != nil {
		return fmt.Errorf("error when update section : %s", err.Error())
	}

	return nil
}
