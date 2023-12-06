package web

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type ModuleCreateRequest struct {
	Title       string                   `json:"title" form:"title" validate:"required"`
	Description string                   `json:"description" form:"description"`
	Type        string                   `json:"type" form:"type" validate:"required"`
	Order       int                      `json:"order" form:"order" validate:"required,min=1"`
	SubModules  []SubModuleCreateRequest `json:"sub_module" form:"sub_module"`
}

type SubModuleCreateRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Type        string `json:"type" form:"type" validate:"required"`
	ContentUrl  string `json:"content_url" form:"content_url"`
	ContentBody string `json:"content_body" form:"content_body"`
}

type ModuleRequest struct {
	SectionID    uint   `json:"section_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Order       int    `json:"order"`
	SubModules  []domain.SubModule `json:"sub_modules"`
}

type ModuleOrderRequest struct {
	Order       int    `json:"order"`
}

type SectionRequest struct {
	CourseID  uint `json:"course_id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Order    int    `json:"order" `
	Modules  domain.Module `json:"modules"`
}

type SectionOrderRequest struct {
	Order    int    `json:"order"`
}