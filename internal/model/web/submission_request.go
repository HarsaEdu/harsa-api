package web

type SubmissionRequest struct {
	Title       string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	ModuleID    uint `json:"module_id" validate:"required"`
}

type SubmissionUpdateRequest struct {
	Title       string `json:"title" `
	Content string `json:"content" `
}
