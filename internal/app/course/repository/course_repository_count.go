package repository

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (courseRepository *CourseRepositoryImpl) CountUserInCourse(courseID uint) (int64,error) {
	
	var countUser int64

	if err := courseRepository.DB.Model(&domain.CourseTracking{}).Where("course_id = ?", courseID).Count(&countUser).Error; err != nil {
		return 0, err
	}

	return countUser, nil
}

func (courseRepository *CourseRepositoryImpl) CountActiveUsersInCourse(courseID uint) (int64, error) {
	var countActiveUsers int64

	if err := courseRepository.DB.Model(&domain.CourseTracking{}).
		Joins("JOIN subscriptions ON course_trackings.user_id = subscriptions.user_id AND subscriptions.end_date >= ?", time.Now()).
		Where("course_id = ?", courseID).
		Count(&countActiveUsers).
		Error; err != nil {
		return 0, err
	}

	return countActiveUsers, nil
}