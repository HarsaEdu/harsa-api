package web

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

type UserForCourseResponse struct {
	ID    uint                         `json:"id"`
	Name string                        `json:"name"`
	Role  string `json:"role"`
}

type CategoryForCourseResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type SubModulesForModuleForCourseResponse struct {
	ID    uint                 `json:"id"`
	Title string               `json:"title"`
	Type  domain.SubModuleType `json:"type"`
}

type ModulesForCourseResponse struct {
	ID         uint                                   `json:"id"`
	Title      string                                 `json:"title"`
	Order      int                                    `json:"order"`
	Type       string                                 `json:"type"`
	SubModules []*SubModulesForModuleForCourseResponse `json:"sub_modules"`
}

type GetCourseResponse struct {
	ID          uint                       `json:"id"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	ImageUrl    string                     `json:"image_url"`
	Enrolled    int                        `json:"enrolled"`
	Rating      float32                        `json:"rating"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
	User        *UserForCourseResponse      `json:"user"`
	Category    *CategoryForCourseResponse  `json:"category"`
}

type GetCourseResponseById struct {
	ID          uint                       `json:"id"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	ImageUrl    string                     `json:"image_url"`
	Enrolled    int                        `json:"enrolled"`
	Rating      int                        `json:"rating"`
	TotalModules int64                        `json:"total_modules"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
	User        *UserForCourseResponse      `json:"user"`
	Category    *CategoryForCourseResponse  `json:"category"`
}