package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/course/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CourseService interface {
	Create(ctx echo.Context, request *web.CourseCreateRequest, instructorId uint) error
	GetAll(offset, limit int, search string, category string) ([]web.GetCourseResponse, *web.Pagination, error)
	GetById(id uint) (*web.GetCourseResponseById, error)
	Update(id uint, userId uint, role string, request *web.CourseUpdateRequest) error
	UpdateImage(ctx echo.Context, id uint, userId uint, role string, request *web.CourseUpdateImageRequest) error
	Delete(id uint, userId uint, role string) error
	GetAllMobile(offset, limit int, search string, category uint) ([]web.GetCourseResponseMobile, *web.Pagination, error)
}

type CourseServiceImpl struct {
	CourseRepository repository.CourseRepository
	Validate *validator.Validate
	CloudinaryUploader cloudinary.CloudinaryUploader
}

func NewCourseService(CourseRepository repository.CourseRepository, Validate *validator.Validate, CloudinaryUploader cloudinary.CloudinaryUploader) CourseService {
	return &CourseServiceImpl {
		CourseRepository: CourseRepository,
		Validate: Validate,
		CloudinaryUploader: CloudinaryUploader,
	}
}