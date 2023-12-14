package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
	"github.com/labstack/echo/v4"
)


func (courseTrackingService *CourseTrackingServiceImpl) FindByIdMobileByUserIdAndCourseId(ctx echo.Context, userID uint, courseID uint) (*web.CourseTrackingResponseMobile, error) {
	isSubscription, err:= courseTrackingService.Subscription.IsSubscription(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("eror when cek subscription  :%s", err.Error())
	}
	
	courseTraking, err := courseTrackingService.CourseTrackingRepository.FindByUserIdAndCourseID(courseID,userID)
	if err != nil { 
		return nil, fmt.Errorf("eror when find course tracking by id  :%s", err.Error())
	}

	course, countModule,countEnroled, err := courseTrackingService.CourseRepository.GetByIdMobile(courseID)
	if err != nil { 
		return nil, fmt.Errorf("eror counting find course tracking  :%s", err.Error())
	}
	if course.ID == 0{
		return nil, fmt.Errorf("course not found")
	}

	listModule, progress, err := courseTrackingService.CourseTrackingRepository.FindAllModuleTrackingWithProgress(course.Section ,userID, courseID)
	if err != nil { 
		return nil, fmt.Errorf("eror when find module tracking :%s", err.Error())
	}

	res := conversion.ConvertCourseTrackingRespose(courseTraking,course,listModule,countEnroled,countModule,progress,isSubscription)
	return res, nil

}

