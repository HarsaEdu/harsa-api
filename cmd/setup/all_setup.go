package setup

import (
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AllSetup(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader,e *echo.Echo) {

	userRoutes, userRepo := UserSetup(db, validate, e)
	authRoutes := AuthSetup(db, validate,userRepo, e)
	moduleRoutes := ModuleSetup(db, validate, e)
	categoryRoutes := CategorySetup(db, validate, cloudinary ,e)
	faqsRoutes := FaqsSetup(db, validate, e)
	courseRoutes := CourseSetup(db, validate, cloudinary, e)

	apiGroup := e.Group("api")
	authRoutes.Auth(apiGroup)
	userRoutes.User(apiGroup)
	categoryRoutes.Category(apiGroup)
	courseRoutes.Course(apiGroup)
	faqsRoutes.Faqs(apiGroup)
	coursesGroup := courseRoutes.Course(apiGroup)
	moduleRoutes.Module(coursesGroup)
}