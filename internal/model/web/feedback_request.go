package web

type FeedbackCreateRequest struct {
	Rating  int    `json:"rating" validate:"required"`
	Content string `json:"content"`
}

type FeedbackUpdateRequest struct {
	Rating  int    `json:"rating"`
	Content string `json:"content"`
}
