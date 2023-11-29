package web

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type UserForTracking struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserForCourse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Job string `json:"job"`
}

type CourseForTraking struct {
	ID           uint          `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	User         UserForCourse `json:"user"`
	ImageUrl     string        `json:"image_url"`
	Enrolled     int64           `json:"enrolled"`
	Rating       float32       `json:"rating"`
	TotalModules int64         `json:"total_modules"`
	Feedback     []FeedBackResponseForTracking `json:"feedback"`   
}

type CourseTrackingResponseMobile struct {
	ID       uint             `json:"id"`
	User     UserForTracking  `json:"user"`
	Progress float32          `json:"progress"`
	Status   domain.StatusCourseTraking `json:"status"`
	Course   CourseForTraking `json:"course"`
	HistoryModule []HistoryModuleResponseMobile `json:"history_module"`
}