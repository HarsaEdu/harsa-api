package web

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)


type SubModuleResponseModule struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	ContentUrl  string                `json:"content_url"`
	ContentBody string                `json:"content_body"`
	Type        domain.SubModuleType  `json:"type"`
}

type ModuleResponse struct {
	ID          uint           `json:"id"`
	CourseID    uint           `json:"course_id"`
	Section     string         `json:"section"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Order       int            `json:"order"`
	Type        string         `json:"type"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	SubModules  []SubModuleResponseModule `json:"sub_modules"`
	Submissions []SubmissionsResponseModule  `json:"submissions"`
	Quizzes 	[]QuizResponseModule 	   `json:"quizzes"`
}