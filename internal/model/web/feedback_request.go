package web

type FeedbackCreateRequest struct {
	CourseID uint    `json:"course_id" validate:"required"`
	UserID   uint    `json:"user_id" validate:"required"`
	Rating  int       `json:"rating" validate:"required"`
	Content string    `json:"content"`
}

type FeedbackUpdateRequest struct {
	Rating  int    `json:"rating"`
	Content string `json:"content"`
}
