package service

import (
	submissionRepoPkg "github.com/HarsaEdu/harsa-api/internal/app/submission/repository"
	"github.com/HarsaEdu/harsa-api/internal/app/submission_answer/repository"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type SubmissionAnswerService interface {
	Create(ctx echo.Context, request *web.SubmissionAnswerRequest, idSubmission, idUser int) error
	Update(ctx echo.Context, request *web.SubmissionAnswerUpdateRequest, id, idUser int) error
	FindById(id int) (*web.SubmissionAnswerResponseMobile, error)
	Get(offset, limit, submissionID int, search string) ([]web.SubmissionAnswerList, *web.Pagination, error)
	UpdateWeb(ctx echo.Context, request *web.SubmissionAnswerUpdateWeb, id int) error
}

type SubmissionAnswerServiceImpl struct {
	Repository         repository.SubmissionAnswerRepository
	Validator          *validator.Validate
	CloudinaryUploader cloudinary.CloudinaryUploader
	SubmissionRepo     submissionRepoPkg.SubmissionRepository
}

func NewSubmissionAnswer(repository repository.SubmissionAnswerRepository, validator *validator.Validate, cloudinaryUploader cloudinary.CloudinaryUploader, submissionRepo submissionRepoPkg.SubmissionRepository) SubmissionAnswerService {
	return &SubmissionAnswerServiceImpl{Repository: repository, Validator: validator, CloudinaryUploader: cloudinaryUploader, SubmissionRepo: submissionRepo}
}
