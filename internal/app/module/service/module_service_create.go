package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (moduleService *ModuleServiceImpl) CreateSection(request *web.SectionRequest, courseId uint, userId uint, role string) error {
	
	err := moduleService.ModuleRepository.CekIdFromCourse(userId, courseId, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user from course :%s", err.Error())
	}

	request.CourseID = courseId

	err = moduleService.Validate.Struct(request)
	if err != nil {
		return err
	}

	existingModule, _ := moduleService.ModuleRepository.GetByTitleAndCourseId(request.Modules.Title, courseId)
	if existingModule != nil {
		return fmt.Errorf("module name already exists")
	}

	existingSection, _ := moduleService.ModuleRepository.GetByTitleSectionAndCourseId(request.Title, courseId)
	if existingSection != nil {
		return fmt.Errorf("section name already exists")
	}

	module := conversion.SectionRequestToSectionDomain(request)

	err = moduleService.ModuleRepository.CreateSection(module)
	if err != nil {
		return fmt.Errorf("error when create module : %s", err.Error())
	}

	return nil
}

func (moduleService *ModuleServiceImpl) CreateModule(request *web.ModuleRequest, courseId uint, userId uint, role string) error {
	
	err := moduleService.ModuleRepository.CekIdFromCourse(userId, courseId, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user from course :%s", err.Error())
	}

	request.CourseID = courseId

	err = moduleService.Validate.Struct(request)
	if err != nil {
		return err
	}

	existingModule, _ := moduleService.ModuleRepository.GetByTitleAndCourseId(request.Title, courseId)
	if existingModule != nil {
		return fmt.Errorf("module name already exists")
	}


	module := conversion.ModuleRequestToModuleDomain(request)

	err = moduleService.ModuleRepository.CreateModule(module)
	if err != nil {
		return fmt.Errorf("error when create module : %s", err.Error())
	}

	return nil
}