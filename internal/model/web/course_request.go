package web

type CourseCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id" validate:"required"`
}