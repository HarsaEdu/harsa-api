package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
	"github.com/labstack/echo/v4"
)

func (courseService *CourseServiceImpl) Update(ctx echo.Context, id uint, userId uint, role string, request *web.CourseUpdateRequest) error {
	err := courseService.Validate.Struct(request)
	if err != nil {
		return err
	}

	existingCourse, err := courseService.CourseRepository.CekIdFromCourse(userId, id, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user in udapte course  : %s", err.Error())
	}
	if existingCourse == nil {
		return fmt.Errorf("course not found")
	}


	updatedData := conversion.CourseUpdateRequestToCourseDomain(request, id)

	imageUrl, err := courseService.CloudinaryUploader.Uploader(ctx, "file", "courses", false)
	if err != nil {
		return fmt.Errorf("error when uploading image : %s", err.Error())
	}

	if imageUrl != "" {
		updatedData.ImageUrl = imageUrl
	}

	err = courseService.CourseRepository.Update(id, updatedData)
	if err != nil {
		return fmt.Errorf("error when update course %s", err.Error())
	}

	return nil
}