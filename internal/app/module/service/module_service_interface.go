package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/module/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type ModuleService interface {
	CreateSection(request *web.SectionRequest, courseId uint, userId uint, role string) error
	CreateModule(equest *web.ModuleRequest, courseId uint, userId uint, role string) error
	GetAllSectionByCourseId(offset, limit int, search string, courseId uint) ([]web.SectionResponse, *web.Pagination, error)
	GetAllModuleBySectionId(sectionId uint) (*web.SectionResponse, error)
	UpdateModule(request *web.ModuleRequest, moduleId uint, userId uint, role string) error
	UpdateModuleOrder(request *web.ModuleOrderRequest, moduleId uint, userId uint, role string) error
	UpdateSectionOrder(request *web.SectionOrderRequest, sectionId uint, userId uint, role string) error
	UpdateSection(request *web.SectionRequest, sectionId uint, userId uint, role string) error
	DeleteModule(moduleId uint, userId uint, role string) error
	DeleteSection(sectionId uint, userId uint, role string) error
	GetModuleById(moduleId uint) (*web.ModuleResponse, error)
	DeleteSubModule(subModuleId uint, userId uint, role string) error
}

type ModuleServiceImpl struct {
	ModuleRepository    repository.ModuleRepository
	Validate            *validator.Validate
}

func NewModuleService(ModuleRepository repository.ModuleRepository, Validate *validator.Validate) ModuleService {
	return &ModuleServiceImpl{
		ModuleRepository:    ModuleRepository,
		Validate:            Validate,
	}
}