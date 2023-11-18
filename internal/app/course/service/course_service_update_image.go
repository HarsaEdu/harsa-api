package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (courseService *CourseServiceImpl) UpdateImage(ctx echo.Context, id uint, request *web.CourseUpdateImageRequest) error {
	err := courseService.Validate.Struct(request)
	if err != nil {
		return err
	}

	existingCourse, err := courseService.CourseRepository.GetById(id)
	if existingCourse == nil {
		return fmt.Errorf("course not found")
	}

	updatedData := conversion.CourseUpdateImageRequestToCourseDomain(request, id)
	updatedData.ImageUrl, err = courseService.CloudinaryUploader.Uploader(ctx, "file", "courses")
	if err != nil {
		return fmt.Errorf("error when uploading image : %s", err.Error())
	}

	err = courseService.CourseRepository.UpdateImage(updatedData)
	if err != nil {
		return fmt.Errorf("error when updating image course : %s", err.Error())
	}

	return nil
}