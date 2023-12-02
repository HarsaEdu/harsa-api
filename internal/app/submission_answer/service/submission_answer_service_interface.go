package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/submission_answer/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type SubmissionAnswerService interface {
	Create(ctx echo.Context, request *web.SubmissionAnswerRequest, idSubmission, idUser int) error
}

type SubmissionAnswerServiceImpl struct {
	Repository         repository.SubmissionAnswerRepository
	Validator          *validator.Validate
	CloudinaryUploader cloudinary.CloudinaryUploader
}

func NewSubmissionAnswer(repository repository.SubmissionAnswerRepository, validator *validator.Validate, cloudinaryUploader cloudinary.CloudinaryUploader) SubmissionAnswerService {
	return &SubmissionAnswerServiceImpl{Repository: repository, Validator: validator, CloudinaryUploader: cloudinaryUploader}
}
