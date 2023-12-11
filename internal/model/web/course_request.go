package web

type CourseCreateRequest struct {
	UserId      uint   `json:"user_id" form:"user_id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description"`
	CategoryID  int    `json:"category_id" form:"category_id" validate:"required"`
	ImageUrl    string `form:"file"`
}

type CourseUpdateRequest struct {
	UserId      uint   `json:"user_id" form:"user_id" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description"`
	CategoryID  int    `json:"category_id" form:"category_id"`
}

type CourseUpdateImageRequest struct {
	ImageUrl string `form:"file"`
}