package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseService *CourseServiceImpl) GetAll() ([]web.GetCourseResponse, error) {
	var courses []web.GetCourseResponse
	result, err := courseService.CourseRepository.GetAll()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf("course not found")
	}

	courses = conversion.CourseDomainToCourseGetAllResponse(result)

	return courses, nil
}