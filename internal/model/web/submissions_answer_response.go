package web

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type SubmissionAnswerResponseMobile struct {
	ID         uint                            `json:"id"`
	Status     domain.StatusSubmissionAnswer   `json:"status"`
	Submission SubmissionsResponseModuleMobile `json:"submission"`
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
