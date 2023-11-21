package setup

import (
	courseHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/course/handler"
	courseRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/course/repository"
	courseRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/course/routes"
	courseServicePkg "github.com/HarsaEdu/harsa-api/internal/app/course/service"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func CourseSetup(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader) courseRoutesPkg.CourseRoutes{
	courseRepository := courseRepositoryPkg.NewCourseRepository(db)
	courseService := courseServicePkg.NewCourseService(courseRepository, validate, cloudinary)
	courseHandler := courseHandlerPkg.NewCourseHandler(courseService)
	courseRoute := courseRoutesPkg.NewCourseRoutes(courseHandler)

	return courseRoute
}
