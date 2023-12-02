package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (courseTrackingService *CourseTrackingServiceImpl) FindSubByIdMobile(moduleID uint, userID uint) (*web.CourseTrackingSub, error) {
	
	res ,err := courseTrackingService.CourseTrackingRepository.FindAllSub(moduleID, userID)
	if err != nil { 
		return nil, fmt.Errorf("eror when find course tracking by id  :%s", err.Error())
	}

	return res, nil

}


