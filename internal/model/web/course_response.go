package web

import "time"

type RoleForUserForCourseResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserForCourseResponse struct {
	ID    uint                         `json:"id"`
	Email string                       `json:"email"`
	Role  RoleForUserForCourseResponse `json:"role"`
}

type CategoryForCourseResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetCourseResponse struct {
	ID          uint                      `json:"id"`
	Category    CategoryForCourseResponse `json:"category"`
	User        UserForCourseResponse     `json:"user"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	ImageUrl    string                    `json:"image_url"`
	Enrolled    int                       `json:"enrolled"`
	Rating      int                       `json:"rating"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
}