package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (courseService *CourseServiceImpl) Update(id uint, userId uint, role string, request *web.CourseUpdateRequest) error {
	err := courseService.Validate.Struct(request)
	if err != nil {
		return err
	}

	existingCourse, err := courseService.CourseRepository.CekIdFromCourse(userId, id, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user in udapte course  : %s", err.Error())
	}
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

func (courseService *CourseServiceImpl) UpdateIntructure(id uint, userId uint, role string, request *web.CourseUpdateRequestIntructure) error {
	err := courseService.Validate.Struct(request)
	if err != nil {
		return err
	}

	existingCourse, err := courseService.CourseRepository.CekIdFromCourse(userId, id, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user in udapte course  : %s", err.Error())
	}
	if existingCourse == nil {
		return fmt.Errorf("course not found")
	}


	updatedData := conversion.CourseUpdateRequestToCourseDomainIntructure(request, id)

	err = courseService.CourseRepository.Update(id, updatedData)
	if err != nil {
		return fmt.Errorf("error when update course %s", err.Error())
	}

	return nil
}