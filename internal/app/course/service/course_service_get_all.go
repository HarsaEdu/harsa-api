package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseService *CourseServiceImpl) GetAll(offset, limit int, search string, category string) ([]web.GetCourseResponse, *web.Pagination, error) {
	var courses []web.GetCourseResponse
	result, total, err := courseService.CourseRepository.GetAll(offset, limit, search, category)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	courses = conversion.CourseDomainToCourseGetAllResponse(result)

	if courses == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return courses, pagination,nil
}

func (courseService *CourseServiceImpl) GetAllMobile(offset, limit int, search string, category uint) ([]web.GetCourseResponseMobile, *web.Pagination, error) {
	var courses []web.GetCourseResponse
	result, total, err := courseService.CourseRepository.GetAllMobile(offset, limit, search, category)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	course := conversion.CourseDomainToCourseGetAllResponseMobile(result)

	if courses == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return course, pagination,nil
}