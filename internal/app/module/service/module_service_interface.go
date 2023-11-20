package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/module/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ModuleService interface {
	Create(ctx echo.Context, request *web.ModuleCreateRequest, courseId uint) error
	GetAllByCourseId(offset, limit int, search string, courseId uint) ([]domain.Module, *web.Pagination, error)
	// GetById(id uint) (*web.GetModuleResponse, error)
	// Update(id uint, request *web.ModuleUpdateRequest) error
	// UpdateImage(ctx echo.Context, id uint, request *web.ModuleUpdateImageRequest) error
	// Delete(id uint) error
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