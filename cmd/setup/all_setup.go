package setup

import (
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AllSetup(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader,e *echo.Echo) {

	userRoutes, userRepo := UserSetup(db, validate)
	authRoutes := AuthSetup(db, validate,userRepo)
	moduleRoutes := ModuleSetup(db, validate)
	categoryRoutes := CategorySetup(db, validate, cloudinary)
	faqsRoutes := FaqsSetup(db, validate)
	courseRoutes := CourseSetup(db, validate, cloudinary)

	apiGroupWeb := e.Group("web")
	authRoutes.AuthWeb(apiGroupWeb)
	userRoutes.UserWeb(apiGroupWeb)
	categoryRoutes.CategoryWeb(apiGroupWeb)
	courseRoutes.CourseWeb(apiGroupWeb)
	faqsRoutes.FaqsWeb(apiGroupWeb)
	coursesGroup := courseRoutes.CourseWeb(apiGroupWeb)
	moduleRoutes.ModuleWeb(coursesGroup)

	apiGroupMobile := e.Group("mobile")
	authRoutes.AuthWeb(apiGroupMobile)
	userRoutes.UserWeb(apiGroupMobile)
	categoryRoutes.CategoryWeb(apiGroupMobile)
	courseRoutes.CourseWeb(apiGroupMobile)
	faqsRoutes.FaqsWeb(apiGroupMobile)
	coursesGroup = courseRoutes.CourseWeb(apiGroupMobile)
	moduleRoutes.ModuleWeb(apiGroupMobile)
}