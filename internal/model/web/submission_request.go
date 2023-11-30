package web

type SubmissionRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	ModuleID    uint
}

type SubmissionUpdateRequest struct {
	Title       string `json:"title" `
	Description string `json:"description" `
}
