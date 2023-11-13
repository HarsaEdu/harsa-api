package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/categories/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CategoryService interface {
	Create(ctx echo.Context, category web.CategoryCreateRequest) error
	Update(ctx echo.Context, category web.CategoryUpdateRequest, id int) error
	UploadImage(ctx echo.Context, category *web.CategoryUploadImageRequest, id int) error
	FindById(ctx echo.Context, id int) (*domain.Category, error)
	GetAll(ctx echo.Context) ([]domain.Category, error)
	Delete(ctx echo.Context, id int) error
}

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validator          *validator.Validate
	cloudinaryUploader cloudinary.CloudinaryUpdloader
}

func NewCategoryService(cr repository.CategoryRepository, validate *validator.Validate, cloudinaryUploader cloudinary.CloudinaryUpdloader) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: cr, Validator: validate, cloudinaryUploader: cloudinaryUploader}
}
