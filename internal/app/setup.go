package app

import (
	auth "github.com/HarsaEdu/harsa-api/internal/app/auth"
	category "github.com/HarsaEdu/harsa-api/internal/app/categories"
	course "github.com/HarsaEdu/harsa-api/internal/app/course"
	faqs "github.com/HarsaEdu/harsa-api/internal/app/faqs"
	feedback "github.com/HarsaEdu/harsa-api/internal/app/feedback"
	module "github.com/HarsaEdu/harsa-api/internal/app/module"
	options "github.com/HarsaEdu/harsa-api/internal/app/options"
	questions "github.com/HarsaEdu/harsa-api/internal/app/questions"
	quizzes "github.com/HarsaEdu/harsa-api/internal/app/quizzes"
	user "github.com/HarsaEdu/harsa-api/internal/app/user"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader, e *echo.Echo) {

	userRoutes, userRepo := user.UserSetup(db, validate)
	authRoutes := auth.AuthSetup(db, validate, userRepo)
	moduleRoutes := module.ModuleSetup(db, validate)
	categoryRoutes := category.CategorySetup(db, validate, cloudinary)
	faqsRoutes := faqs.FaqsSetup(db, validate)
	courseRoutes := course.CourseSetup(db, validate, cloudinary)
	quizzesRoutes := quizzes.QuizzesSetup(db, validate)
	questionsRoutes := questions.QuestionsSetup(db, validate)
	optionsRoutes := options.OptionsSetup(db, validate)
	feedbackRoutes := feedback.FeedbackSetup(db, validate)

	apiGroupWeb := e.Group("web")
	authRoutes.AuthWeb(apiGroupWeb)
	userRoutes.UserWeb(apiGroupWeb)
	categoryRoutes.CategoryWeb(apiGroupWeb)
	courseRoutes.CourseWeb(apiGroupWeb)
	faqsRoutes.FaqsWeb(apiGroupWeb)
	coursesGroup := courseRoutes.CourseWeb(apiGroupWeb)
	moduleRoutes.ModuleWeb(coursesGroup)
	quizzesRoutes.QuizzesWeb(coursesGroup)
	questionsRoutes.QuestionsWeb(coursesGroup)
	optionsRoutes.OptionsWeb(coursesGroup)
	feedbackRoutes.FeedbackWeb(apiGroupWeb)

	apiGroupMobile := e.Group("mobile")
	authRoutes.AuthMobile(apiGroupMobile)
	userRoutes.UserMobile(apiGroupMobile)
	categoryRoutes.CategoryMobile(apiGroupMobile)
	courseRoutes.CourseMobile(apiGroupMobile)
	faqsRoutes.FaqsMobile(apiGroupMobile)
	coursesGroup = courseRoutes.CourseMobile(apiGroupMobile)
	moduleRoutes.ModuleMobile(apiGroupMobile)
	quizzesRoutes.QuizzesMobile(coursesGroup)
	feedbackRoutes.FeedbackMobile(apiGroupMobile)
}
