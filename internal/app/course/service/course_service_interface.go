package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/course/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CourseService interface {
	Create(request *web.CourseCreateRequest, instructorId uint) error
	GetAll() ([]web.GetCourseResponse, error)
	GetById(id uint) (*web.GetCourseResponse, error)
	Update(id uint, request *web.CourseUpdateRequest) error
	UpdateImage(ctx echo.Context, id uint, request *web.CourseUpdateImageRequest) error
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