package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HarsaEdu/harsa-api/configs"
	authHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/handler"
	authRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/repository"
	authRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/auth/routes"
	authServicePkg "github.com/HarsaEdu/harsa-api/internal/app/auth/service"
	categoryHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/categories/handler"
	categoryRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/categories/repository"
	categoryRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/categories/routes"
	categoryServicePkg "github.com/HarsaEdu/harsa-api/internal/app/categories/service"
	courseHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/course/handler"
	courseRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/course/repository"
	courseRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/course/routes"
	courseServicePkg "github.com/HarsaEdu/harsa-api/internal/app/course/service"
	faqsHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/faqs/handler"
	faqsRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/faqs/repository"
	faqsRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/faqs/routes"
	faqsServicePkg "github.com/HarsaEdu/harsa-api/internal/app/faqs/service"
	feedbackHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/feedback/handler"
	feedbackRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/feedback/repository"
	feedbackRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/feedback/routes"
	feedbackServicePkg "github.com/HarsaEdu/harsa-api/internal/app/feedback/service"
	moduleHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/module/handler"
	moduleRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/module/repository"
	moduleRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/module/routes"
	moduleServicePkg "github.com/HarsaEdu/harsa-api/internal/app/module/service"
	userHandlerPkg "github.com/HarsaEdu/harsa-api/internal/app/user/handler"
	userRepositoryPkg "github.com/HarsaEdu/harsa-api/internal/app/user/repository"
	userRoutesPkg "github.com/HarsaEdu/harsa-api/internal/app/user/routes"
	userServicePkg "github.com/HarsaEdu/harsa-api/internal/app/user/service"
	"github.com/HarsaEdu/harsa-api/internal/infrastructure/database"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/HarsaEdu/harsa-api/web"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load configuration
	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatal("Error loading config:", err.Error())
	}

	// Initialize database connection
	db, err := database.NewMySQLConnection(&config.MySQL)
	if err != nil {
		logrus.Fatal("Error connecting to MySQL:", err.Error())
	}

	cloudinaryUploader := cloudinary.NewClodinaryUploader(&config.Cloudinary)

	// Create an validator instance
	validate := validator.New()

	// Create an Echo instance
	e := echo.New()

	// Setup App
	// Repository
	authRepository := authRepositoryPkg.NewAuthRepository(db)
	userRepository := userRepositoryPkg.NewUserRepository(db)
	categoryRepository := categoryRepositoryPkg.NewCategoryRepository(db)
	courseRepository := courseRepositoryPkg.NewCourseRepository(db)
	feedbackRepository := feedbackRepositoryPkg.NewFeedbackRepository(db)
	faqsRepository := faqsRepositoryPkg.NewFaqRepository(db)
	moduleRepository := moduleRepositoryPkg.NewModuleRepository(db)

	// Service
	authService := authServicePkg.NewAuthService(authRepository, userRepository, validate)
	userService := userServicePkg.NewUserService(userRepository, validate)
	categoryService := categoryServicePkg.NewCategoryService(categoryRepository, validate, cloudinaryUploader)
	courseService := courseServicePkg.NewCourseService(courseRepository, validate, cloudinaryUploader)
	feedbackService := feedbackServicePkg.NewFeedbackService(feedbackRepository, validate)
	faqsService := faqsServicePkg.NewFaqsService(faqsRepository, validate)
	moduleService := moduleServicePkg.NewModuleService(moduleRepository, validate)

	// Handler
	authHandler := authHandlerPkg.NewAuthHandler(authService)
	userHandler := userHandlerPkg.NewUserHandler(userService)
	categoryHandler := categoryHandlerPkg.NewCategoryHandler(categoryService)
	courseHandler := courseHandlerPkg.NewCourseHandler(courseService)
	feedbackHandler := feedbackHandlerPkg.NewFeedbackHandler(feedbackService)
	faqsHandler := faqsHandlerPkg.NewFaqsHandler(faqsService)
	moduleHandler := moduleHandlerPkg.NewModuleHandler(moduleService)

	// Routes
	authRoutes := authRoutesPkg.NewAuthRoutes(e, authHandler)
	userRoutes := userRoutesPkg.NewUserRoutes(userHandler)
	categoryRoutes := categoryRoutesPkg.NewCategoryRoutes(e, categoryHandler)
	courseRoutes := courseRoutesPkg.NewCourseRoutes(courseHandler)
	feedbackRoutes := feedbackRoutesPkg.NewFeedbackRoutes(e, feedbackHandler)
	faqsRoutes := faqsRoutesPkg.NewFaqsRoutes(e, faqsHandler)
	moduleRoutes := moduleRoutesPkg.NewModuleRoutes(moduleHandler)

	// Setup Routes
	apiGroup := e.Group("api")
	authRoutes.Auth(apiGroup)
	userRoutes.User(apiGroup)
	categoryRoutes.Category(apiGroup)
	courseRoutes.Course(apiGroup)
	feedbackRoutes.Feedback(apiGroup)
	faqsRoutes.Faqs(apiGroup)
	coursesGroup := courseRoutes.Course(apiGroup)
	moduleRoutes.Module(coursesGroup)

	// Serve static HTML file for the root path
	e.GET("/", func(c echo.Context) error {
		file, err := web.Content.ReadFile("index.html")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error reading HTML file")
		}
		return c.HTMLBlob(http.StatusOK, file)
	})

	// Middleware and server configuration
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	// Start the Echo server in a goroutine
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Shutdown Echo gracefully
	if err := e.Shutdown(context.Background()); err != nil {
		logrus.Fatal("Error shutting down server:", err)
	}

	logrus.Info("Server shut down gracefully")
}
