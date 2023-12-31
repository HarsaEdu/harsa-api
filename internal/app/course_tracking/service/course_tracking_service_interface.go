package service

import (
	repositoryCourse "github.com/HarsaEdu/harsa-api/internal/app/course/repository"
	repositoryCourseTracking "github.com/HarsaEdu/harsa-api/internal/app/course_tracking/repository"
	quizzes "github.com/HarsaEdu/harsa-api/internal/app/quizzes/service"
	subscriptionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/HarsaEdu/harsa-api/internal/pkg/firebase"
)

type CourseTrackingService interface {
	Create(ctx echo.Context,request web.CourseTrackingRequest) error
	// FindByIdMobile(crourseTrackingId uint) (*web.CourseTrackingResponseMobile, error)
	FindSubByIdMobile(moduleID uint, userID uint) (*web.CourseTrackingSub, error)
	FindSubModuleByID(ctx echo.Context,  moduleID uint, subModuleID uint, userID uint) (*web.SubModuleTracking, error)
	FindSubmissionByID(ctx echo.Context, moduleID uint, userID uint, submissionID uint) (*web.SubmissionAnswerTrackingByIDResponse, error)
	FindQuizzByID(ctx echo.Context, moduleID uint, userID uint, quizzID uint) (*web.HistoryQuizIDTracking, error)
	FindModuleHistory(ctx echo.Context, moduleID uint, userID uint) (*web.ModuleTrackingByID, error)
	GetAllCourseByUserIdMobile(offset, limit int, search string, userID uint, status string) ([]web.GetAllCourseForTraking, *web.Pagination, error)
	GetAllCourseByUserIdWeb(offset, limit int, userID uint) ([]web.CourseTrackingResponseWeb, *web.Pagination, error)
	GetAllUserCourseWeb(offset, limit int, courseID uint, search string) ([]web.CourseTrackingUserWeb, *web.Pagination, error)
	Delete(courseTrackingId uint, courseId uint,userId uint, role string) error
	FindByIdMobileByUserIdAndCourseId(ctx echo.Context, userID uint, courseID uint) (*web.CourseTrackingResponseMobile, error)
	CreateWeb(request web.CourseTrackingRequest) error
}

type CourseTrackingServiceImpl struct {
	Subscription             subscriptionServicePkg.SubscriptionService
	CourseRepository         repositoryCourse.CourseRepository
	CourseTrackingRepository repositoryCourseTracking.CourseTrackingRepository
	QuizzService             quizzes.QuizzesService
	Validator                *validator.Validate
	Firebase                 firebase.Firebase
}

func NewCourseTrackingService(repositoryTracking repositoryCourseTracking.CourseTrackingRepository, validator *validator.Validate, courseRepository repositoryCourse.CourseRepository, quizzService quizzes.QuizzesService, subscription subscriptionServicePkg.SubscriptionService, firebaseImpl firebase.Firebase) CourseTrackingService {
	return &CourseTrackingServiceImpl{
		Subscription:             subscription,
		CourseRepository:         courseRepository,
		CourseTrackingRepository: repositoryTracking,
		QuizzService:             quizzService,
		Validator:                validator,
		Firebase:                 firebaseImpl,
	}
}
