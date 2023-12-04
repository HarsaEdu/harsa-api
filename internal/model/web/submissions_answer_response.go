package web

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

type SubmissionAnswerResponseMobile struct {
	ID         uint                          `json:"id"`
	Status     domain.StatusSubmissionAnswer `json:"status"`
	Submission string                        `json:"submission"`
	Feedback   string                        `json:"feedback"`
}

type SubmissionAnswerList struct {
	Peserta string `json:"peserta"`
	Status  string `json:"status"`
}

type SubmissionAnswerResponseWeb struct {
	Submission SubmissionsResponseModule `json:"submission"`
	SubmissionAnswer []domain.SubmissionsAnswerDetail `json:"submission_answer"`
}
