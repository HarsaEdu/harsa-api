package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseTrackingService *CourseTrackingServiceImpl) FindByIdMobile(crourseTrackingId uint) (*web.CourseTrackingResponseMobile, error) {
	
	courseTraking ,err := courseTrackingService.CourseTrackingRepository.FindById(crourseTrackingId)
	if err != nil { 
		return nil, fmt.Errorf("eror when find course tracking by id  :%s", err.Error())
	}

	course, countModule ,countEnroled ,err := courseTrackingService.CourseRepository.GetByIdMobile(courseTraking.CourseID)
	if err != nil { 
		return nil, fmt.Errorf(" :%s", err.Error())
	}

	progress ,err := courseTrackingService.CourseTrackingRepository.CountProgressCourse(courseTraking.CourseID,courseTraking.UserID)
	if err != nil { 
		return nil, fmt.Errorf(" :%s", err.Error())
	}

	listModule ,err := courseTrackingService.CourseTrackingRepository.FindAllModuleTracking(course.Modules ,courseTraking.UserID)
	if err != nil { 
		return nil, fmt.Errorf("eror when find module tracking :%s", err.Error())
	}

	res := conversion.ConvertCourseTrackingRespose(courseTraking,course,listModule,countEnroled,countModule,progress)

	return res, nil

}
