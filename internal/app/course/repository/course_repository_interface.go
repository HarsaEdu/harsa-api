package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *domain.Course) error
	GetAll(offset, limit int, search string, category string) ([]domain.Course, int64, error)
	GetById(id uint) (*domain.Course, error)
	Update(id uint, course *domain.Course) error
	UpdateImage(course *domain.Course) error
	Delete(course *domain.Course) error
	CekIdFromCourse(userId uint, courseId uint, role string) (*domain.Course, error)
}

type CourseRepositoryImpl struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{
		DB: db,
	}
}