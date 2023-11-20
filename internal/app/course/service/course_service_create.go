package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (courseService *CourseServiceImpl) Create(ctx echo.Context, request *web.CourseCreateRequest, instructorId uint) error {
	err := courseService.Validate.Struct(request)
	if err != nil {
		return err
	}

	course := conversion.CourseCreateRequestToCourseDomain(request, instructorId)
	imageUrl, err := courseService.CloudinaryUploader.Uploader(ctx, "file", "courses", false)
	if err != nil {
		return fmt.Errorf("error when uploading image : %s", err.Error())
	}

	if imageUrl != "" {
		course.ImageUrl = imageUrl
	}

	err = courseService.CourseRepository.Create(course)
	if err != nil {
		return fmt.Errorf("error when creating course %s", err.Error())
	}

	return nil
}
