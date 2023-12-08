package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
	
)

func (courseTrackingService *CourseTrackingServiceImpl) GetAllCourseByUserIdMobile(offset, limit int, search string, userID uint, status string) ([]web.GetAllCourseForTraking, *web.Pagination, error) {
	
	result, total ,err := courseTrackingService.CourseTrackingRepository.GetAllCourseTrackingMobile(offset, limit, userID, search, status)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return result, pagination,nil
}

func (courseTrackingService *CourseTrackingServiceImpl) GetAllCourseByUserIdWeb(offset, limit int, userID uint) ([]web.CourseTrackingResponseWeb, *web.Pagination, error) {
	
	result, total ,err := courseTrackingService.CourseTrackingRepository.GetAllCourseTrackingWeb(offset, limit, userID)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	response := conversion.ConvertAllCourseTrackingResponeWeb(result)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return response, pagination,nil
}

func (courseTrackingService *CourseTrackingServiceImpl) GetAllUserCourseWeb(offset, limit int, courseID uint, search string) ([]web.CourseTrackingUserWeb, *web.Pagination, error) {
	
	result, total ,err := courseTrackingService.CourseTrackingRepository.GetAllCourseTrackingUserWeb(offset, limit,courseID, search)
	if err != nil {
		return nil, nil, err
	}

	if result == nil {
		return nil, nil, fmt.Errorf("user not found")
	}

	response := conversion.ConvertAllCourseTackingUserWeb(result)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)
	
	return response, pagination,nil
}