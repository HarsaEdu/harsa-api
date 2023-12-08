package web

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type SubmissionAnswerRequest struct {
	SubmittedUrl string `form:"file" validate:"required"`
	Status       domain.StatusSubmissionAnswer
}

type SubmissionAnswerUpdateRequest struct {
	SubmittedUrl string `form:"file" validate:"required"`
}
