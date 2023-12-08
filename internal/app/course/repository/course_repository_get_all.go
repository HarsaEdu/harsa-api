package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
	"gorm.io/gorm"
)

func (courseRepository *CourseRepositoryImpl) GetAllByUserId(offset, limit int, search string, userID uint) ([]domain.Course, int64, error) {
	if offset < 0 || limit <= 0 {
		return nil, 0, fmt.Errorf("invalid offset or limit")
	}

	courses := []domain.Course{}
	var total int64

	query := courseRepository.DB.Model(&courses).Order("created_at DESC")

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("courses.title LIKE ? OR courses.description LIKE ?", s, s)
	}
	query.Where("user_id = ?", userID).Count(&total).Find(&courses)
	query = query.Limit(limit).Offset(offset)

	if offset >= int(total) {
		return nil, 0,  nil
	}

	result := query.Preload("Feedback.User.UserProfile", func(db *gorm.DB) *gorm.DB {
		return db.Limit(3)
	}).Find(&courses)

	if result.Error != nil {
		return nil, 0,  result.Error
	}

	return courses, total, nil
}

func (courseRepository *CourseRepositoryImpl) GetAllCourseByUserId(offset, limit int, search string, userID uint) ([]domain.Course, int64, error) {
	if offset < 0 || limit <= 0 {
		return nil, 0, fmt.Errorf("invalid offset or limit")
	}

	courses := []domain.Course{}
	var total int64

	query := courseRepository.DB.Model(&courses).Order("created_at DESC")

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("courses.title LIKE ? OR courses.description LIKE ?", s, s)
	}
	query.Where("user_id = ?", userID).Count(&total).Find(&courses)
	query = query.Limit(limit).Offset(offset)

	if offset >= int(total) {
		return nil, 0,  nil
	}

	result := query.Find(&courses)

	if result.Error != nil {
		return nil, 0,  result.Error
	}

	return courses, total, nil
}

func (courseRepository *CourseRepositoryImpl) GetNameUser(userId uint) (string, error) {
	
	var userProfile = domain.UserProfile{}

	if err := courseRepository.DB.First(&userProfile).Where("user_id = ?", userId).Error; err != nil {
		return "", err
	}

	return userProfile.FirstName, nil
}


func (courseRepository *CourseRepositoryImpl) GetDashBoardIntructur(offset, limit int, search string, userID uint) (*web.DashboardIntructur, int64,error) {
	courses, total,  err := courseRepository.GetAllByUserId(offset, limit, search, userID)
	if err != nil {
		return nil,0, fmt.Errorf("get all eror : %s", err)
	}
	name ,err:= courseRepository.GetNameUser(userID)
	if err != nil {
		return nil,0, fmt.Errorf("get get name eror : %s", err)
	}
	res := []web.CourseResponseForIntructur{}
	for _, course := range courses {
		countUser , err:=  courseRepository.CountUserInCourse(course.ID)
		if err != nil {
			return nil,0,  fmt.Errorf("get count user eror : %s", err)
		}
		countUserActive ,err:=  courseRepository.CountActiveUsersInCourse(course.ID)
		if err != nil {
			return nil,0, fmt.Errorf("get count user active eror : %s", err)
		}

		convert := conversion.CourseDomainToCourseGetResponse(&course,countUser,countUserActive)

		res = append(res, *convert)
	}

	result := conversion.CourseResponseForIntructur(res,name)

	return result,total, err
}

func (courseRepository *CourseRepositoryImpl) GetAllCourseDashBoardIntructur(offset, limit int, search string, userID uint) (*web.DashboardAllCourseIntructur, int64,error) {
	courses, total,  err := courseRepository.GetAllCourseByUserId(offset, limit, search, userID)
	if err != nil {
		return nil,0, fmt.Errorf("get all eror : %w", err)
	}
	name ,err:= courseRepository.GetNameUser(userID)
	if err != nil {
		return nil,0, fmt.Errorf("get get name eror : %w", err)
	}
	res := []web.AllCourseResponseForIntructur{}
	for _, course := range courses {

		convert := conversion.ConvertAllCourseIntructure(&course)

		res = append(res, *convert)
	}

	result := conversion.AllCourseResponseForIntructur(res,name)

	return result,total, err
}

func (courseRepository *CourseRepositoryImpl) GetDetailDashBoardIntructur(courseID uint) (*web.CourseResponseForIntructur, error) {

	course := domain.Course{}

	result := courseRepository.DB.Preload("Feedback.User.UserProfile", func(db *gorm.DB) *gorm.DB {
		return db.Limit(3)
	}).First(&course, courseID)

	if result.Error != nil {
		return nil,   result.Error
	}

		countUser , err:=  courseRepository.CountUserInCourse(course.ID)
		if err != nil {
			return nil, fmt.Errorf("get count user eror : %w", err)
		}
		countUserActive ,err:=  courseRepository.CountActiveUsersInCourse(course.ID)
		if err != nil {
			return nil, fmt.Errorf("get count user active eror : %w", err)
		}

		convert := conversion.CourseDomainToCourseGetResponse(&course,countUser,countUserActive)

	return convert, err
}

func (courseRepository *CourseRepositoryImpl) GetAllByCategory(offset, limit int, search string, categoryId uint) ([]domain.Course, int64, error) {

	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	course := []domain.Course{}
	var total int64

	query := courseRepository.DB.Model(&course).Order("created_at DESC")

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("courses.title LIKE ? OR courses.description LIKE ?", s, s)
	}

	query.Where("category_id = ?", categoryId).Find(&course).Count(&total)

	query = query.Limit(limit).Offset(offset)

	result := query.Preload("Category").Preload("User.UserProfile").Find(&course)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if offset >= int(total) {
		return nil, 0, nil
	}

	return course, total, nil
}

func (courseRepository *CourseRepositoryImpl) GetAll(offset, limit int, search string, category string) ([]domain.Course, int64, error) {
    if offset < 0 || limit < 0 {
        return nil, 0, fmt.Errorf("Offset and limit must be non-negative")
    }

    courses := []domain.Course{}
    var total int64

    query := courseRepository.DB.Model(&courses).Preload("User.UserProfile").Order("created_at DESC")

    if search != "" {
        s := "%" + search + "%"
        query = query.Where("courses.title LIKE ? OR courses.description LIKE ?", s, s)
    }

    if category != "" {
        query = query.Joins("JOIN categories ON categories.id = courses.category_id").Where("categories.name LIKE ?", "%"+category+"%")
    }

    query.Model(&courses).Count(&total)

    result := query.Limit(limit).Offset(offset).Preload("Category").Find(&courses)

    if result.Error != nil {
        return nil, 0, result.Error
    }

    if offset >= int(total) {
        return nil, 0, nil
    }

    return courses, total, nil
}


func (courseRepository *CourseRepositoryImpl) GetAllbyUserID(offset, limit int, search string, category string, userID uint) ([]domain.Course, int64, error) {
    if offset < 0 || limit < 0 {
        return nil, 0, fmt.Errorf("Offset and limit must be non-negative")
    }

    courses := []domain.Course{}
    var total int64

    query := courseRepository.DB.Model(&courses).Where("user_id = ? ", userID).Order("created_at DESC").Preload("User.UserProfile")

    if search != "" {
        s := "%" + search + "%"
        query = query.Where("courses.title LIKE ? OR courses.description LIKE ?", s, s)
    }

    if category != "" {
        query = query.Joins("JOIN categories ON categories.id = courses.category_id").Where("categories.name LIKE ?", "%"+category+"%")
    }

    query.Model(&courses).Count(&total)

    result := query.Limit(limit).Offset(offset).Preload("Category").Find(&courses)

    if result.Error != nil {
        return nil, 0, result.Error
    }

    if offset >= int(total) {
        return nil, 0, nil
    }

    return courses, total, nil
}