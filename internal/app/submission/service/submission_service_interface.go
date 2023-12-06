package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/submission/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type SubmissionService interface {
	Create(ctx echo.Context, request web.SubmissionRequest) error
	Update(ctx echo.Context, request *web.SubmissionUpdateRequest, submissionId int) error
	Delete(id int) error
	GetAll(moduleId int) ([]web.SubmissionsResponseModule, error)
	FindById(id int) (*web.SubmissionsResponseModule, error)
	GetAllMobile(moduleId int) ([]web.SubmissionsResponseModuleMobile, error)
}

type SubmissionServiceImpl struct {
	SubmissionRepository repository.SubmissionRepository
	Validator            *validator.Validate
}

func NewSubmissionService(repository repository.SubmissionRepository, validate validator.Validate) SubmissionService {
	return &SubmissionServiceImpl{SubmissionRepository: repository, Validator: &validate}
}
