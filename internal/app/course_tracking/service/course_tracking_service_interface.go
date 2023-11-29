package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/course_tracking/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type CourseTrackingService interface {
	Create(request web.CourseTrackingRequest) error
	FindByIdMobile(crourseTrackingId uint) (*web.CourseTrackingResponseMobile, error)
}

type CourseTrackingServiceImpl struct {
	CourseTrackingRepository repository.CourseTrackingRepository
	Validator          *validator.Validate
}

func NewCourseTrackingService(repository repository.CourseTrackingRepository, validator *validator.Validate) CourseTrackingService {
	return &CourseTrackingServiceImpl{
		CourseTrackingRepository: repository,
		Validator:          validator,
	}
}
