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
	ID      uint   `gorm:"type:int;primarykey" json:"id"`
	Title   string `json:"title"`
	Content string `gorm:"type:text" json:"content"`
	Peserta *UserForCourseResponse
	Answer  *SubmissionAnswerResponseMobile
}
