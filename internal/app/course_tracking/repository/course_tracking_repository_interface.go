package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"gorm.io/gorm"
)

type CourseTrackingRepository interface {
	Create(enrolled *domain.CourseTracking) error
	FindById(courseTrackingId uint) (*domain.CourseTracking, error)
	CountProgressModule(moduleID uint, userID uint) (float32, error)
	CountProgressCourse(courseID uint, userID uint) (float32, error)
	FindAllModuleTracking(module []domain.Module, userID uint) ([]web.ModuleResponseForTracking, error)
	FindAllSubModule(moduleId uint, userID uint) ([]web.SubModuleResponseForTracking, error)
	FindAllSubmission(moduleId uint, userID uint) ([]web.SubmissionsResponseModuleMobile, error)
	FindAllQuiz(moduleId uint, userID uint) ([]web.QuizResponseForTracking, error)
	FindAllSub(moduleId uint, userID uint) (*web.CourseTrackingSub, error)
	FindModuleTracking(moduleID uint, userID uint) (*web.ModuleResponseForTracking, error)
	FindSubModuleByID(moduleID uint, userID uint, subModuleID uint) (*domain.HistorySubModule, *domain.SubModule, error)
}

type CourseTrackingRepositoryImpl struct {
	DB *gorm.DB
}

func NewCourseTrackingRepository(db *gorm.DB) CourseTrackingRepository {
	return &CourseTrackingRepositoryImpl{DB: db}
}
