package service

import (

	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (categoryService *DashboardServiceImpl) CountAll() (*web.AllCountDashboard, error) {

	countCourse, err := categoryService.DashboardRepository.CountCourse()
	if err != nil {
		return nil, err
	}

	counInstructor, err := categoryService.DashboardRepository.CountIntructure()
	if err != nil {
		return nil, err
	}

	countInterest, countStudent, err := categoryService.DashboardRepository.CountInterest()
	if err != nil {
		return nil, err
	}

	allCount := &web.AllCountDashboard{
		CountInterest: countInterest,
		CountCourse: countCourse,
		CountStudent: countStudent,
		CountIntructure: counInstructor,
	}

	return allCount, nil
}