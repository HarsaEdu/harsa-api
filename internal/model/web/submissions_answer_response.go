package web

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type SubmissionAnswerResponseMobile struct {
	ID         uint                            `json:"id"`
	Status     domain.StatusSubmissionAnswer   `json:"status"`
	Submission SubmissionsResponseModuleMobile `json:"submission"`
}
