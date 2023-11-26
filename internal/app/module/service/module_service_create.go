package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (moduleService *ModuleServiceImpl) Create(ctx echo.Context, request *web.ModuleCreateRequest, courseId uint, userId uint, role string) error {
	
	err := moduleService.ModuleRepository.CekIdFromCourse(userId, courseId, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user from course :%s", err.Error())
	}

	err = moduleService.Validate.Struct(request)
	if err != nil {
		return err
	}

	existingModule, _ := moduleService.ModuleRepository.GetByTitleAndCourseId(request.Title, courseId)
	if existingModule != nil {
		return fmt.Errorf("module name already exists")
	}

	existingOrder, _ := moduleService.ModuleRepository.GetByOrderAndCourseId(request.Order, courseId)
	if existingOrder != nil {
		return fmt.Errorf("module order already exists")
	}

	module := conversion.ModuleCreateRequestToModuleDomain(request, courseId)

	err = moduleService.ModuleRepository.Create(module)
	if err != nil {
		return fmt.Errorf("error when create module : %s", err.Error())
	}

	return nil
}