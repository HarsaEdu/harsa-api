package web

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)


type SubModuleResponseModule struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	ContentUrl  string                `json:"content_url"`
	Type        domain.SubModuleType  `json:"type"`
}

type SubModuleResponseModuleMobile struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Type        domain.SubModuleType  `json:"type"`
}

type ModuleResponse struct {
	ID          uint           `json:"id"`
	SectionID    uint           `json:"section_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Order       int            `json:"order"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	SubModules  []SubModuleResponseModule `json:"sub_modules"`
	Submissions []SubmissionsResponseModule  `json:"submissions"`
	Quizzes 	[]QuizResponseModule 	   `json:"quizzes"`
}

type ModuleResponseMobile struct {
	ID          uint           `json:"id"`
	SectionID    uint           `json:"section_id"`
	Section     string         `json:"section"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Order       int            `json:"order"`
	Type        string         `json:"type"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	SubModules  []SubModuleResponseModuleMobile `json:"sub_modules"`
	Submissions []SubmissionsResponseModuleMobile  `json:"submissions"`
	Quizzes 	[]QuizResponseModuleMobile 	   `json:"quizzes"`
}

type SectionResponse struct {
	ID          uint           `json:"id"`
	CourseID    uint           `json:"course_id"`
	Title       string         `json:"title"`
	Order       int            `json:"order"`
	Modules     []ModuleResponseWithTitle `json:"modules"`
}

type SectionResponseMobile struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	Order       int            `json:"order"`
	Modules     []ModuleResponseForTracking `json:"modules"`
}
type ModuleResponseWithTitle struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Order       int            `json:"order"`
}