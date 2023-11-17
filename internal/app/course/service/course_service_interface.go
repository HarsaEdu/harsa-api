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
	GetAll(offset, limit int, search string) ([]web.GetCourseResponse, *web.Pagination, error)
	GetById(id uint) (*web.GetCourseResponse, error)
	Update(id uint, request *web.CourseUpdateRequest) error
	UpdateImage(ctx echo.Context, id uint, request *web.CourseUpdateImageRequest) error
	Delete(id uint) error
}

type CourseServiceImpl struct {
	CourseRepository repository.CourseRepository
	Validate *validator.Validate
	CloudinaryUpdloader cloudinary.CloudinaryUpdloader
}

func NewCourseService(CourseRepository repository.CourseRepository, Validate *validator.Validate, CloudinaryUpdloader cloudinary.CloudinaryUpdloader) CourseService {
	return &CourseServiceImpl {
		CourseRepository: CourseRepository,
		Validate: Validate,
		CloudinaryUpdloader: CloudinaryUpdloader,
	}
}