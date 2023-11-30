package web

type FaqsRequest struct {
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
}

type FaqsUpdateRequest struct {
	Question string `json:"question" form:"question"`
	Answer   string `json:"answer" form:"answer"`
}
