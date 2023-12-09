package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (courseTrackingService *CourseTrackingServiceImpl) Create(ctx echo.Context,request web.CourseTrackingRequest) error {
	
	isSubscription, err:= courseTrackingService.Subscription.IsSubscription(ctx, request.UserID)
	if err != nil {
		return fmt.Errorf("eror when cek subscription  :%s", err.Error())
	}

	if !isSubscription {
		return fmt.Errorf("unauthorized")
	}
	
	exist, err := courseTrackingService.CourseTrackingRepository.Cek(request.UserID, request.CourseID)
	if exist != nil {
		return fmt.Errorf("course tracking already exist")
	}
	
	err = courseTrackingService.Validator.Struct(request)
	if err != nil {
		return err
	}

	courseTraking := conversionRequest.CourseTrackingReqToDomain(&request)

	err = courseTrackingService.CourseTrackingRepository.Create(courseTraking)
	if err != nil {
		return fmt.Errorf("error when creating Course Traking %s", err.Error())
	}

	return nil

}

