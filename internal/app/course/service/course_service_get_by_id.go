package service

import (
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseService *CourseServiceImpl) GetById(id uint) (*web.GetCourseResponseById, error) {
	result, module , err := courseService.CourseRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := conversion.CourseDomainToCourseGetByIdResponse(result, module)

	return response, nil
}

func (courseService *CourseServiceImpl) GetByIdMobile(id uint) (*web.CourseForTraking, error) {
	result, countModule, countEnroled , err := courseService.CourseRepository.GetByIdMobile(id)
	if err != nil {
		return nil, err
	}

	response := conversion.ConvertCourse(result,countEnroled,countModule)

	return response, nil
}