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

type SubmissionAnswerTrackingResponse struct {
	ID               uint                            `json:"id"`
	Title            string                          `json:"title"`
	Progress         float32                         `json:"progress"`
	Description      string                          `json:"description"`
	Order            int                             `json:"order"`
	SubmissionAnswer SubmissionsResponseModuleMobile `json:"submission_answer"`
	Submission       SubmissionsResponseModule       `json:"submission"`
}

type SubmissionAnswerTrackingByIDResponse struct {
	ID               uint                      `json:"id"`
	Title            string                    `json:"title"`
	Progress         float32                   `json:"progress"`
	Description      string                    `json:"description"`
	Order            int                       `json:"order"`
	SubmissionAnswer SubmissionAnswerTracking  `json:"submission_answer"`
	Submission       SubmissionsResponseModule `json:"submission"`
}
