package web

type SubmissionRequest struct {
	Title       string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	ModuleID    uint
}

type SubmissionUpdateRequest struct {
	Title       string `json:"title" `
	Content string `json:"content" `
}
