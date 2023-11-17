package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseService *CourseServiceImpl) GetAll(offset, limit int, search string) ([]web.GetCourseResponse, *web.Pagination, error) {
	var courses []web.GetCourseResponse
	result, total, err := courseService.CourseRepository.GetAll(offset, limit, search)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	courses = conversion.CourseDomainToCourseGetAllResponse(result)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return courses, pagination,nil
}