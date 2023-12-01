package web

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type SubmissionsResponseModule struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type SubmissionsResponseModuleMobile struct {
	Id               uint                `json:"id"`
	Title            string              `json:"title"`
	SubmissionAnswer SubmissionAnswerRes `json:"submission_answer"`
	Is_complete      bool                `json:"is_complete"`
}

type SubmissionAnswerRes struct {
	Id     uint                          `json:"id"`
	Status domain.StatusSubmissionAnswer `json:"status"`
}
