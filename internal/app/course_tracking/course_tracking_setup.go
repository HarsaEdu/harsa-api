package coursetracking

import (
	"github.com/HarsaEdu/harsa-api/internal/app/course/repository"
	courseTrackingHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/course_tracking/handler"
	courseTrackingRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/course_tracking/repository"
	courseTrackingRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/course_tracking/routes"
	courseTrackingServicePkg "github.com/HarsaEdu/harsa-api/internal/app/course_tracking/service"
	quizzesServicePkg "github.com/HarsaEdu/harsa-api/internal/app/quizzes/service"
	subscriptionServicePkg "github.com/HarsaEdu/harsa-api/internal/app/subscription/service"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"github.com/HarsaEdu/harsa-api/internal/pkg/firebase"
)

func CourseTrackingSetup(db *gorm.DB, validate *validator.Validate, courseRepository repository.CourseRepository, quizzService quizzesServicePkg.QuizzesService,subscriptionService subscriptionServicePkg.SubscriptionService,firebaseImpl firebase.Firebase) (courseTrackingRoutesPkg.CourseTrackingRoutes, courseTrackingRepositoryPkg.CourseTrackingRepository) {
	courseTrackingRepository := courseTrackingRepositoryPkg.NewCourseTrackingRepository(db)
	courseTrackingService := courseTrackingServicePkg.NewCourseTrackingService(courseTrackingRepository, validate, courseRepository, quizzService, subscriptionService, firebaseImpl)
	courseTrackingHandler := courseTrackingHandlerPkg.NewCourseTrackingHandler(courseTrackingService)
	courseTrackingRoute := courseTrackingRoutesPkg.NewCourseTrackingRoutes(courseTrackingHandler)

	return courseTrackingRoute, courseTrackingRepository
}
