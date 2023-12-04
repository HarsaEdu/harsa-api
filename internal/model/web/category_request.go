package web

type CategoryCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type CategoryUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryUploadImageRequest struct {
	CategoryImage string `json:"file" form:"file" validate:"required"`
}
