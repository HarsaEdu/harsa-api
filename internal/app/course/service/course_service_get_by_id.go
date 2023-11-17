package service

import (
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseService *CourseServiceImpl) GetById(id uint) (*web.GetCourseResponse, error) {
	result, err := courseService.CourseRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := conversion.CourseDomainToCourseGetByIdResponse(result)

	return response, nil
}