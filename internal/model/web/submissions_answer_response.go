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
	ID   uint `json:"id"`
	Name string `json:"name"`
	Status  string `json:"status"`
}

type SubmissionAnswerResponseWeb struct {
	Submission SubmissionsResponseModule `json:"submission"`
	SubmissionAnswer []SubmissionAnswerList `json:"submission_answer"`
}
