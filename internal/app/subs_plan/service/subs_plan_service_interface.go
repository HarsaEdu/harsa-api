package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/subs_plan/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type SubsPlanService interface {
	GetAll(search string) ([]domain.SubsPlan, *web.Pagination, error)
	Create(ctx echo.Context, subsPlan *web.SubsPlanCreateRequest) error
	Update(subsPlan *web.SubsPlanUpdateRequest, id int) error
	Delete(id int) error
	UpdateImage(ctx echo.Context, subsPlan *web.SubsPlanUpdateImage, id int) error
}

type SubsPlanServiceImpl struct {
	SubsPlanRepository repository.SubsPlanRepository
	Validator          *validator.Validate
	CloudinaryUploader cloudinary.CloudinaryUploader
}

func NewsubsPlanService(subsPlanRepository repository.SubsPlanRepository, validator *validator.Validate, cloudinary cloudinary.CloudinaryUploader) SubsPlanService {
	return &SubsPlanServiceImpl{SubsPlanRepository: subsPlanRepository, Validator: validator, CloudinaryUploader: cloudinary}

}
