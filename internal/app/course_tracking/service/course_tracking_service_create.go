package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversionRequest "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (courseTrackingService *CourseTrackingServiceImpl) Create(request web.CourseTrackingRequest) error {
	err := courseTrackingService.Validator.Struct(request)
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
