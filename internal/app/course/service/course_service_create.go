package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (courseService *CourseServiceImpl) Create(request web.CourseCreateRequest, instructorId uint) error {
	err := courseService.Validate.Struct(request)
	if err != nil {
		return err
	}

	course := conversion.CourseCreateRequestToCourseDomain(request, instructorId)

	err = courseService.CourseRepository.Create(course)
	if err != nil {
		return fmt.Errorf("error when creating course %s", err.Error())
	}

	return nil
}
