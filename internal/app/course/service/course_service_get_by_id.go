package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseService *CourseServiceImpl) GetById(id uint) (*web.GetCourseResponseByIdWeb, error) {
	result,  err := courseService.CourseRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	response := conversion.ConvertCourseGetByIdResponseWeb(result)

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

	if result.ID == 0{
		return nil, fmt.Errorf("not found")
	}

	response := conversion.ConvertCourse(result,countEnroled,countModule)

	return response, nil
}