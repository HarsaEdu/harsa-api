package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseTrackingService *CourseTrackingServiceImpl) FindByIdMobile(crourseTrackingId uint) (*web.CourseTrackingResponseMobile, error) {
	
	courseTraking ,countEnroled, countModule ,err := courseTrackingService.CourseTrackingRepository.FindById(crourseTrackingId)
	if err != nil { 
		return nil, fmt.Errorf(" :%s", err.Error())
	}

	res := conversion.ConvertCourseTrackingRespose(courseTraking,countEnroled,countModule )

	return res, nil

}
