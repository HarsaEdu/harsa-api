package web

type CourseCreateRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description"`
	CategoryID  int    `json:"category_id" form:"category_id" validate:"required"`
}

type CourseUpdateRequest struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	CategoryID  int    `json:"category_id" form:"category_id"`
}

type CourseUpdateImageRequest struct {
	ImageUrl string `form:"file"`
}