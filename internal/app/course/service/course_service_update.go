package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (courseService *CourseServiceImpl) Update(id uint, request *web.CourseUpdateRequest) error {
	err := courseService.Validate.Struct(request)
	if err != nil {
		return err
	}

	existingCourse, _ := courseService.CourseRepository.GetById(id)
	if existingCourse == nil {
		return fmt.Errorf("course not found")
	}


	updatedData := conversion.CourseUpdateRequestToCourseDomain(request, id)

	err = courseService.CourseRepository.Update(id, updatedData)
	if err != nil {
		return fmt.Errorf("error when update course %s", err.Error())
	}

	return nil
}