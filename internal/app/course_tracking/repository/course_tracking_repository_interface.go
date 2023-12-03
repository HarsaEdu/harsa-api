package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"gorm.io/gorm"
)

type CourseTrackingRepository interface {
	Create(enrolled *domain.CourseTracking) error
	FindById(courseTrackingId uint) (*domain.CourseTracking ,error)
	CountProgressModule(moduleID uint, userID uint) (float32 ,error)
	CountProgressCourse(courseID uint, userID uint) (float32 ,error)
	FindAllModuleTracking(sections []domain.Section, userID uint) ([]web.SectionResponseMobile, error)
	FindAllSubModule(moduleId uint, userID uint) ([]web.SubModuleResponseForTracking, error)
	FindAllSubmission(moduleId uint, userID uint) ([]web.SubmissionsResponseModuleMobile, error)
	FindAllQuiz(moduleId uint, userID uint) ([]web.QuizResponseForTracking, error)
	FindAllSub(moduleId uint, userID uint) (*web.CourseTrackingSub, error)
	FindModuleTracking(moduleID uint, userID uint) (*web.ModuleResponseForTracking, error)
	FindSubModuleByID(moduleID uint, userID uint, subModuleID uint) (*domain.HistorySubModule, *domain.SubModule, error)
	FindSubmissionByID(moduleID uint, userID uint, subModuleID uint) (*domain.SubmissionAnswer, *domain.Submissions, error)
	FindQuizzByID(moduleID uint, userID uint, quizID uint) (*domain.HistoryQuiz, error)
	GetAllCourseTracking(offset, limit int, userID uint,search string, status string) ([]domain.CourseTracking, int64, error)
	GetAllCourseTrackingMobile(offset, limit int, userID uint,search string, status string) ([]web.GetAllCourseForTraking, int64, error)
	FindAllModuleTrackingWithProgress(sections []domain.Section, userID uint, courseID uint) ([]web.SectionResponseMobile, float32 ,error)
	GetAllCourseTrackingWeb(offset, limit int, userID uint) ([]domain.CourseTracking, int64, error)
	GetAllCourseTrackingUserWeb(offset, limit int, courseID uint, search string) ([]domain.CourseTracking, int64, error)
	Delete(id uint) error
	CekIdFromCourse(userId uint, courseId uint, role string) error
}

type CourseTrackingRepositoryImpl struct {
	DB *gorm.DB
}

func NewCourseTrackingRepository(db *gorm.DB) CourseTrackingRepository {
	return &CourseTrackingRepositoryImpl{DB: db}
}
