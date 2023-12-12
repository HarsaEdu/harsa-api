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

type UserForCourseResponseMobile struct {
	ID    uint                         `json:"id"`
	Name string                        `json:"name"`
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

type CourseResponseForIntructur struct {
	ID          uint                       `json:"id"`
	Title       string                     `json:"title"`
	ImageUrl    string                     `json:"image_url"`
	FeedBack    []FeedBackResponseForTracking `json:"feedback"`
	TotalUser   int64                        `json:"total_user"`
	TotalActiveUser int64                        `json:"total_active_user"`
}

type AllCourseResponseForIntructur struct {
	ID          uint                       `json:"id"`
	Title       string                     `json:"title"`
	ImageUrl    string                     `json:"image_url"`
}

type DashboardAllCourseIntructur struct {
	FirstName string `json:"first_name"`
	Course    []AllCourseResponseForIntructur `json:"course"`
}

type DashboardIntructur struct {
	FirstName string `json:"first_name"`
	Course    []CourseResponseForIntructur `json:"course"`
}

type GetCourseResponseMobile struct {
	ID          uint                       `json:"id"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	ImageUrl    string                     `json:"image_url"`
	Rating      float32                        `json:"rating"`
	User        *UserForCourseResponseMobile      `json:"user"`
	Category    *CategoryForCourseResponse  `json:"category"`
}

type CourseResponse struct{
	ID          uint                       `json:"id"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	UserID      uint                       `json:"user_id"`
}

type GetCourseResponseById struct {
	ID          uint                       `json:"id"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	ImageUrl    string                     `json:"image_url"`
	Enrolled    int                        `json:"enrolled"`
	Rating      float32                        `json:"rating"`
	TotalModules int64                        `json:"total_modules"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
	User        *UserForCourseResponse      `json:"user"`
	Category    *CategoryForCourseResponse  `json:"category"`
}

type GetCourseResponseByIdWeb struct{
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImageUrl string `json:"image_url"`
	User        *UserForCourseResponseMobile      `json:"user"`
	Section []SectionResponseGetByIdWeb `json:"section"`
}

type SectionResponseGetByIdWeb struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	OrderBy int `json:"order"`
	Module []ModuleResponseGetByIdWeb `json:"module"`
}

type ModuleResponseGetByIdWeb struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	OrderBy int `json:"order"`
}