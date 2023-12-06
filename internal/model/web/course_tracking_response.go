package web

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

type UserForTracking struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type UserForCourse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Job string `json:"job"`
	ImageUrl string `json:"image_url"`
}

type CourseForTraking struct {
	ID           uint          `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	User         UserForCourse `json:"user"`
	ImageUrl     string        `json:"image_url"`
	Enrolled     int64         `json:"enrolled"`
	Rating       float32       `json:"rating"`
	TotalModules int64         `json:"total_modules"`
	Feedback     []FeedBackResponseForTracking `json:"feedback"`   
}

type GetAllCourseForTraking struct {
	TrakingID    uint          `json:"traking_id"`
	ID           uint          `json:"course_id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	UserIntructur   UserForCourse `json:"user_intructur"`
	UserStudent   UserForTracking `json:"user_student"`
	ImageUrl     string        `json:"image_url"`
	Status       domain.StatusCourseTraking `json:"status"`
	Progress     float32       `json:"progress"`
	CreatedAt    time.Time        `json:"created_at"`
}

type CourseTrackingResponse struct {
	ID       uint             `json:"id"`
	User     UserForTracking  `json:"user"`
	Progress float32          `json:"progress"`
	Status   domain.StatusCourseTraking `json:"status"`
}

type CourseTrackingResponseMobile struct {
	CourseTracking CourseTrackingResponse `json:"course_tracking"`
	Course         CourseForTraking `json:"course"`
	Sections		   []SectionResponseMobile `json:"sections"`
}

type CourseTrackingSub struct {
	SubModules []SubModuleResponseForTracking `json:"sub_modules"`
	Submissions []SubmissionsResponseModuleMobile `json:"submissions"`
	Quizzes     []QuizResponseForTracking `json:"quizzes"`
}

type CourseTrackingResponseWeb struct{
	ID       uint             `json:"id"`
	Title    string           `json:"title"`
}

type CourseTrackingUserWeb struct{
	CourseTrakingID uint             `json:"course_tracking_id"`
	UserID       uint             `json:"user_id"`
	Name     string           `json:"name"`
	UserName string           `json:"user_name"`
	Email    string           `json:"email"`
	PhoneNumber string        `json:"phone_number"`
	Address  string           `json:"address"`
}