package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseTrackingRepository *CourseTrackingRepositoryImpl) GetAllCourseTracking(offset, limit int, userID uint, search string, status string) ([]domain.CourseTracking, int64, error) {
	if offset < 0 || limit <= 0 {
		return nil, 0, fmt.Errorf("invalid offset or limit")
	}

	courseTracking := []domain.CourseTracking{}
	var total int64

	query := courseTrackingRepository.DB.Model(&courseTracking).Where("user_id = ?", userID).Order("created_at DESC")

	if status != "" {
		s := "%" + status + "%"
		query = query.Where("status LIKE ?", s)
	}

	query = query.Preload("Course.User.UserProfile").Preload("User.UserProfile")

	if search != "" {
		query = query.Joins("JOIN courses ON courses.id = course_trackings.course_id").
			Where("courses.title LIKE ? ", "%"+search+"%")
	}

	query.Count(&total)

	result := query.Limit(limit).Offset(offset).Find(&courseTracking)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 || offset >= int(total) {
		return nil, 0, nil
	}

	return courseTracking, total, nil
}

func (courseTrackingRepository *CourseTrackingRepositoryImpl) GetAllCourseTrackingMobile(offset, limit int, userID uint,search string, status string) ([]web.GetAllCourseForTraking, int64, error) {
	
	coursesTracking , total ,err := courseTrackingRepository.GetAllCourseTracking(offset, limit, userID,search, status)
	if err != nil {
		return nil,0,  fmt.Errorf("get all course tracking eror : %s", err.Error())
	}

	if coursesTracking == nil {
		return nil, 0, fmt.Errorf("course not found")
	}

	allCourseTracking := []web.GetAllCourseForTraking{}

	for _, tracking := range coursesTracking {
		progress ,err := courseTrackingRepository.CountProgressCourse(tracking.Course.ID, userID)
		if err != nil { 
			return nil, 0,fmt.Errorf("progress course eror :%s", err.Error())
		}

		allCourseTracking = append(allCourseTracking, *conversion.ConvertCourseTrakingMobile(&tracking, progress))

	}

	return allCourseTracking, total, nil
}


func (courseTrackingRepository *CourseTrackingRepositoryImpl) GetAllCourseTrackingWeb(offset, limit int, userID uint) ([]domain.CourseTracking, int64, error) {
	if offset < 0 || limit <= 0 {
		return nil, 0, fmt.Errorf("invalid offset or limit")
	}

	courseTracking := []domain.CourseTracking{}
	var total int64

	query := courseTrackingRepository.DB.Model(&courseTracking).Where("user_id = ?", userID).Order("created_at DESC")

	query = query.Preload("Course")

	query.Count(&total)

	result := query.Limit(limit).Offset(offset).Find(&courseTracking)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 || offset >= int(total) {
		return nil, 0, nil
	}

	return courseTracking, total, nil
}


func (courseTrackingRepository *CourseTrackingRepositoryImpl) GetAllCourseTrackingUserWeb(offset, limit int, courseID uint, search string) ([]domain.CourseTracking, int64, error) {
	if offset < 0 || limit <= 0 {
		return nil, 0, fmt.Errorf("invalid offset or limit")
	}

	courseTracking := []domain.CourseTracking{}
	var total int64

	query := courseTrackingRepository.DB.Model(&courseTracking).Where("course_id = ?",courseID).Order("created_at DESC")


	query = query.Preload("User.UserProfile")

	if search != "" {
		query = query.Joins("JOIN user_profiles ON user_profiles.user_id = course_trackings.user_id").
		Where("user_profiles.first_name LIKE ? OR user_profiles.last_name LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Count(&total)

	result := query.Limit(limit).Offset(offset).Find(&courseTracking)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 || offset >= int(total) {
		return nil, 0, nil
	}

	return courseTracking, total, nil
}

