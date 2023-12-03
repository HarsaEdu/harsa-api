package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *domain.Course) error
	GetAllByUserId(offset, limit int, search string, userID uint) ([]domain.Course, int64, error)
	GetById(id uint) (*domain.Course, error)
	Update(id uint, course *domain.Course) error
	GetNameUser(userId uint) (string, error)
	UpdateImage(course *domain.Course) error
	Delete(course *domain.Course) error
	CekIdFromCourse(userId uint, courseId uint, role string) (*domain.Course, error)
	GetByIdMobile(id uint) (*domain.Course, int64, int64,error)
	GetAll(offset, limit int, search string, categoryId uint) ([]domain.Course, int64, error)
	GetDashBoardIntructur(offset, limit int, search string, userID uint) (*web.DashboardIntructur, int64,error)
	GetAllCourseByUserId(offset, limit int, search string, userID uint) ([]domain.Course, int64, error)
	GetAllCourseDashBoardIntructur(offset, limit int, search string, userID uint) (*web.DashboardAllCourseIntructur, int64,error)
	GetDetailDashBoardIntructur(courseID uint) (*web.CourseResponseForIntructur, error)
}

type CourseRepositoryImpl struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{
		DB: db,
	}
}