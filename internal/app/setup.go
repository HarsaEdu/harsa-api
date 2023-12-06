package app

import (
	auth "github.com/HarsaEdu/harsa-api/internal/app/auth"
	category "github.com/HarsaEdu/harsa-api/internal/app/categories"
	"github.com/HarsaEdu/harsa-api/internal/app/chatbot"
	course "github.com/HarsaEdu/harsa-api/internal/app/course"
	courseTraking "github.com/HarsaEdu/harsa-api/internal/app/course_tracking"
	faqs "github.com/HarsaEdu/harsa-api/internal/app/faqs"
	feedback "github.com/HarsaEdu/harsa-api/internal/app/feedback"
	historySubModule "github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules"
	interest "github.com/HarsaEdu/harsa-api/internal/app/interest"
	module "github.com/HarsaEdu/harsa-api/internal/app/module"
	options "github.com/HarsaEdu/harsa-api/internal/app/options"
	"github.com/HarsaEdu/harsa-api/internal/app/payment"
	profile "github.com/HarsaEdu/harsa-api/internal/app/profile"
	questions "github.com/HarsaEdu/harsa-api/internal/app/questions"
	quizzes "github.com/HarsaEdu/harsa-api/internal/app/quizzes"
	"github.com/HarsaEdu/harsa-api/internal/app/recommendations"
	submission "github.com/HarsaEdu/harsa-api/internal/app/submission"
	subsPlan "github.com/HarsaEdu/harsa-api/internal/app/subs_plan"

	user "github.com/HarsaEdu/harsa-api/internal/app/user"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/HarsaEdu/harsa-api/internal/pkg/midtrans"
	"github.com/HarsaEdu/harsa-api/internal/pkg/openai"
	recommendationsApi "github.com/HarsaEdu/harsa-api/internal/pkg/recommendations"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB, validate *validator.Validate, cloudinary cloudinary.CloudinaryUploader, e *echo.Echo, openai openai.OpenAi, midtransCoreApi midtrans.MidtransCoreApi, recommendationsApi recommendationsApi.RecommendationsApi) {

	userRoutes, userRepo := user.UserSetup(db, validate)
	authRoutes := auth.AuthSetup(db, validate, userRepo)
	moduleRoutes := module.ModuleSetup(db, validate)
	categoryRoutes := category.CategorySetup(db, validate, cloudinary)
	faqsRoutes := faqs.FaqsSetup(db, validate)
	courseRepsoitory, courseRoutes := course.CourseSetup(db, validate, cloudinary)
	subsPlanRoutes, subsPlanRepo := subsPlan.SubsPlanSetup(db, validate, cloudinary)
	quizzesRoutes, quizzService := quizzes.QuizzesSetup(db, validate)
	profileRoutes, profileRepo := profile.ProfileSetup(db, validate, e, cloudinary)
	interestRoutes := interest.InterestSetup(db, validate, profileRepo)
	questionsRoutes := questions.QuestionsSetup(db, validate)
	optionsRoutes := options.OptionsSetup(db, validate)
	feedbackRoutes := feedback.FeedbackSetup(db, validate)
	chatbotRoutes := chatbot.ChatbotSetup(db, validate, userRepo, openai)
	submissionRoutes := submission.SubmissionSetup(db, validate)
	paymentRoutes := payment.PaymentSetup(db, validate, midtransCoreApi, userRepo, subsPlanRepo)
	courseTrakingRoutes := courseTraking.CourseTrackingSetup(db, validate, courseRepsoitory, quizzService)
	historySubModuleRoutes := historySubModule.HistorySubModuleSetup(db, validate)
	recommendationsRoutes := recommendations.RecommendationsSetup(validate, recommendationsApi, userRepo)

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
	submissionRoutes.SubmissionWeb(coursesGroup)
	paymentRoutes.PaymentWeb(apiGroupWeb)
	courseTrakingRoutes.CourseTrackingWeb(apiGroupWeb)
	recommendationsRoutes.RecommendationsWeb(apiGroupWeb)

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
	feedbackRoutes.FeedbackMobile(coursesGroup)
	interestRoutes.MobileInterest(apiGroupMobile)
	chatbotRoutes.ChatbotMobile(apiGroupMobile)
	submissionRoutes.SubmissionMobile(coursesGroup)

	paymentRoutes.PaymentMobile(apiGroupMobile)
	paymentRoutes.PaymentSubscriptionsMobile(apiGroupMobile)
	courseTrakingRoutes.CourseTrackingMobile(apiGroupMobile)
	historySubModuleRoutes.MobileHistorySubModule(apiGroupMobile)
	recommendationsRoutes.RecommendationsMobile(apiGroupMobile)
}
