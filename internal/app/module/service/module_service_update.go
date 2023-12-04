package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (moduleService *ModuleServiceImpl) UpdateModule(request *web.ModuleRequest, moduleId uint, userId uint, role string) error {

	existingModule, err := moduleService.ModuleRepository.CekIdFromModule(userId, moduleId, role)
	if err != nil {
		return fmt.Errorf("error when cek id user from course :%s", err.Error())
	}

	request.SectionID = existingModule.SectionID

	err = moduleService.Validate.Struct(request)
	if err != nil {
		return err
	}

	existingModuleTitle, _ := moduleService.ModuleRepository.GetByTitleAndSectionId(request.Title, existingModule.SectionID)
	if existingModuleTitle != nil {
		return fmt.Errorf("module name already exists")
	}

	module := conversion.ModuleRequestToModuleDomain(request)

	err = moduleService.ModuleRepository.UpdateModule(module, existingModule)
	if err != nil {
		return fmt.Errorf("error when update module : %s", err.Error())
	}

	return nil
}

func (moduleService *ModuleServiceImpl) UpdateModuleOrder(request *web.ModuleOrderRequest, moduleId uint, userId uint, role string) error {

	existingModule, err := moduleService.ModuleRepository.CekIdFromModule(userId, moduleId, role)
	if err != nil {
		return fmt.Errorf("error when cek id user from course :%s", err.Error())
	}

	err = moduleService.Validate.Struct(request)
	if err != nil {
		return err
	}

	err = moduleService.ModuleRepository.UpdateOrderModule(request.Order, existingModule)
	if err != nil {
		return fmt.Errorf("error when create module : %s", err.Error())
	}

	return nil
}

func (moduleService *ModuleServiceImpl) UpdateSectionOrder(request *web.SectionOrderRequest, sectionId uint, userId uint, role string) error {

	existingSection, err := moduleService.ModuleRepository.CekIdFromSection(userId, sectionId, role)
	if err != nil {
		return fmt.Errorf("error when cek id user from course :%s", err.Error())
	}

	err = moduleService.Validate.Struct(request)
	if err != nil {
		return err
	}

	err = moduleService.ModuleRepository.UpdateOrderSection(request.Order, existingSection)
	if err != nil {
		return fmt.Errorf("error when create module : %s", err.Error())
	}

	return nil
}

func (moduleService *ModuleServiceImpl) UpdateSection(request *web.SectionRequest, sectionId uint, userId uint, role string) error {
	
	existingSection, err := moduleService.ModuleRepository.CekIdFromSection(userId, sectionId, role)
	if err != nil {
		return fmt.Errorf("error when cek id user from course :%s", err.Error())
	}

	request.CourseID = existingSection.CourseID

	err = moduleService.Validate.Struct(request)
	if err != nil {
		return err
	}

	existingSectionTitle, _ := moduleService.ModuleRepository.GetByTitleSectionAndCourseId(request.Title, existingSection.CourseID)
	if existingSectionTitle != nil {
		return fmt.Errorf("section name already exists")
	}

	module := conversion.SectionRequestToSectionDomain(request)

	err = moduleService.ModuleRepository.UpdateSection(module, existingSection)
	if err != nil {
		return fmt.Errorf("error when update section : %s", err.Error())
	}

	return nil
}