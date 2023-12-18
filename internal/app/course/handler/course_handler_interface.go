package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/course/service"
	"github.com/labstack/echo/v4"
)

type CourseHandler interface {
	Create(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetById(ctx echo.Context) error
	Update(ctx echo.Context) error
	UpdateImage(ctx echo.Context) error
	Delete(ctx echo.Context) error
	GetAllByUserId(ctx echo.Context) error
	GetDetailCourseById(ctx echo.Context) error
	GetAllCourseByUserId(ctx echo.Context) error
	GetAllByCategory(ctx echo.Context) error
	GetByIdMobile(ctx echo.Context) error
	GetAllMyCourse(ctx echo.Context) error
	GetAllByRating(ctx echo.Context) error
	GetAllMobile(ctx echo.Context) error 
}

type CourseHandlerImpl struct {
	CourseService service.CourseService
}

func NewCourseHandler(CourseService service.CourseService) CourseHandler {
	return &CourseHandlerImpl{
		CourseService: CourseService,
	}
}