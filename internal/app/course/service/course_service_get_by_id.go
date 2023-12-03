package service

import (
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseService *CourseServiceImpl) GetById(id uint) (*web.CourseResponse, error) {
	result,  err := courseService.CourseRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := conversion.ConvertCourseResponse(result)

	return response, nil
}

func (courseService *CourseServiceImpl) GetDeatailCourse(id uint) (*web.CourseResponseForIntructur, error) {
	result,  err := courseService.CourseRepository.GetDetailDashBoardIntructur(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (courseService *CourseServiceImpl) GetByIdMobile(id uint) (*web.CourseForTraking, error) {
	result, countModule, countEnroled , err := courseService.CourseRepository.GetByIdMobile(id)
	if err != nil {
		return nil, err
	}

	response := conversion.ConvertCourse(result,countEnroled,countModule)

	return response, nil
}