package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (moduleService *ModuleServiceImpl) GetAllSectionByCourseId(offset, limit int, search string, courseId uint) ([]web.SectionResponse, *web.Pagination, error) {
	result, total, err := moduleService.ModuleRepository.GetAllSectionByCourseId(offset, limit, search, courseId)
	if err != nil {
		return nil, nil, err
	
	}

	if result == nil {
		return nil, nil, fmt.Errorf("Section not found")
	}

	res := conversion.ConvertAllSectionResponse(result)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return res, pagination, nil
}

func (moduleService *ModuleServiceImpl) GetAllModuleBySectionId(sectionId uint) (*web.SectionResponse, error) {
	result, err := moduleService.ModuleRepository.GetAllModuleBySecsionId(sectionId)
	if err != nil {
		return nil,  fmt.Errorf("error when get module by section id:%s", err.Error())
	}

	res := conversion.ConvertSectionResponse(result)

	return res,  nil
}

func (moduleService *ModuleServiceImpl) GetModuleById(moduleId uint) (*web.ModuleResponse, error) {
	result, err := moduleService.ModuleRepository.GetModuleById(moduleId)
	if err != nil {
		return nil,  fmt.Errorf("error when get module by section id:%s", err.Error())
	}

	res := conversion.ConvertModuleResponse(result)

	return res,  nil
}