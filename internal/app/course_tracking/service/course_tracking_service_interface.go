package service

import (
	repositoryCourse "github.com/HarsaEdu/harsa-api/internal/app/course/repository"
	repositoryCourseTracking "github.com/HarsaEdu/harsa-api/internal/app/course_tracking/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type CourseTrackingService interface {
	Create(request web.CourseTrackingRequest) error
	FindByIdMobile(crourseTrackingId uint) (*web.CourseTrackingResponseMobile, error)
	FindSubByIdMobile(courseID uint, userID uint) (*web.CourseTrackingSub, error)
	FindSubModuleByID(moduleID uint, subModuleID uint, userID uint) (*web.SubModuleTrackingResponse, error)
}

type CourseTrackingServiceImpl struct {
	CourseRepository         repositoryCourse.CourseRepository
	CourseTrackingRepository repositoryCourseTracking.CourseTrackingRepository
	Validator                *validator.Validate
}

func NewCourseTrackingService(repositoryTracking repositoryCourseTracking.CourseTrackingRepository, validator *validator.Validate, courseRepository repositoryCourse.CourseRepository) CourseTrackingService {
	return &CourseTrackingServiceImpl{
		CourseRepository:         courseRepository,
		CourseTrackingRepository: repositoryTracking,
		Validator:                validator,
	}
}
