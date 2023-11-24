package app

import (
	auth "github.com/HarsaEdu/harsa-api/internal/app/auth"
	category "github.com/HarsaEdu/harsa-api/internal/app/categories"
	course "github.com/HarsaEdu/harsa-api/internal/app/course"
	faqs "github.com/HarsaEdu/harsa-api/internal/app/faqs"
	feedback "github.com/HarsaEdu/harsa-api/internal/app/feedback"
	subsPlan "github.com/HarsaEdu/harsa-api/internal/app/subs_plan"
	interest "github.com/HarsaEdu/harsa-api/internal/app/interest"
	module "github.com/HarsaEdu/harsa-api/internal/app/module"
	profile "github.com/HarsaEdu/harsa-api/internal/app/profile"
	quizzes "github.com/HarsaEdu/harsa-api/internal/app/quizzes"
	questions "github.com/HarsaEdu/harsa-api/internal/app/questions"
	options "github.com/HarsaEdu/harsa-api/internal/app/options"
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
	subsPlanRoutes := subsPlan.SubsPlanSetup(db, validate, cloudinary)
	profileRoutes := profile.ProfileSetup(db, validate, e, cloudinary)
	quizzesRoutes := quizzes.QuizzesSetup(db, validate)
	interestRoutes := interest.InterestSetup(db, validate)
	questionsRoutes := questions.QuestionsSetup(db, validate)
	optionsRoutes := options.OptionsSetup(db, validate)
	feedbackRoutes := feedback.FeedbackSetup(db, validate)

	apiGroupWeb := e.Group("web")
	authRoutes.AuthWeb(apiGroupWeb)
	userRoutes.UserWeb(apiGroupWeb)
	categoryRoutes.CategoryWeb(apiGroupWeb)
	courseRoutes.CourseWeb(apiGroupWeb)
	faqsRoutes.FaqsWeb(apiGroupWeb)
	subsPlanRoutes.SubsPlanWeb(apiGroupWeb)
	coursesGroup := courseRoutes.CourseWeb(apiGroupWeb)
	moduleRoutes.ModuleWeb(coursesGroup)
	quizzesRoutes.QuizzesWeb(coursesGroup)
	interestRoutes.WebInterest(apiGroupWeb)
	questionsRoutes.QuestionsWeb(coursesGroup)
	optionsRoutes.OptionsWeb(coursesGroup)
	feedbackRoutes.FeedbackWeb(apiGroupWeb)

	apiGroupMobile := e.Group("mobile")
	authRoutes.AuthMobile(apiGroupMobile)
	userRoutes.UserMobile(apiGroupMobile)
	categoryRoutes.CategoryMobile(apiGroupMobile)
	courseRoutes.CourseMobile(apiGroupMobile)
	faqsRoutes.FaqsMobile(apiGroupMobile)
	subsPlanRoutes.SubsPlanMobile(apiGroupMobile)
	coursesGroup = courseRoutes.CourseMobile(apiGroupMobile)
	moduleRoutes.ModuleMobile(apiGroupMobile)
	profileRoutes.ProfileMobile(apiGroupMobile)
	quizzesRoutes.QuizzesMobile(coursesGroup)
	feedbackRoutes.FeedbackMobile(apiGroupMobile)
	interestRoutes.MobileInterest(apiGroupMobile)
}
