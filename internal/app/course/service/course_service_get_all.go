package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseService *CourseServiceImpl) GetAllByUserId(offset, limit int, search string, userID uint) (*web.DashboardIntructur, *web.Pagination, error) {
	
	result, total ,err := courseService.CourseRepository.GetDashBoardIntructur(offset, limit, search, userID)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return result, pagination,nil
}

func (courseService *CourseServiceImpl) GetAllCourseByUserId(offset, limit int, search string, userID uint) (*web.DashboardAllCourseIntructur, *web.Pagination, error) {
	
	result, total ,err := courseService.CourseRepository.GetAllCourseDashBoardIntructur(offset, limit, search, userID)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return result, pagination,nil
}

func (courseService *CourseServiceImpl) GetAll(offset, limit int, search string, category string) ([]web.GetCourseResponseMobile, *web.Pagination, error) {
	
	result, total, err := courseService.CourseRepository.GetAll(offset, limit, search, category)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	course := conversion.CourseDomainToCourseGetAllResponseMobile(result)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return course, pagination,nil
}

func (courseService *CourseServiceImpl) GetAllByCategory(offset, limit int, search string, category uint) ([]web.GetCourseResponseMobile, *web.Pagination, error) {
	
	result, total, err := courseService.CourseRepository.GetAllByCategory(offset, limit, search, category)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	course := conversion.CourseDomainToCourseGetAllResponseMobile(result)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return course, pagination,nil
}