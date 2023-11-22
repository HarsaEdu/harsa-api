package app

import (
	auth "github.com/HarsaEdu/harsa-api/internal/app/auth"
	category "github.com/HarsaEdu/harsa-api/internal/app/categories"
	module "github.com/HarsaEdu/harsa-api/internal/app/module"
	faqs "github.com/HarsaEdu/harsa-api/internal/app/faqs"
	user "github.com/HarsaEdu/harsa-api/internal/app/user"
	course "github.com/HarsaEdu/harsa-api/internal/app/course"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader,e *echo.Echo) {

	userRoutes, userRepo := user.UserSetup(db, validate)
	authRoutes := auth.AuthSetup(db, validate,userRepo)
	moduleRoutes := module.ModuleSetup(db, validate)
	categoryRoutes := category.CategorySetup(db, validate, cloudinary)
	faqsRoutes := faqs.FaqsSetup(db, validate)
	courseRoutes := course.CourseSetup(db, validate, cloudinary)

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