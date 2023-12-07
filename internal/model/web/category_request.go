package web

type CategoryCreateRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description"`
	Image_Url   string `json:"image" form:"image" `
	Icon        string `json:"icon" form:"icon" `
}

type CategoryUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryUploadImageRequest struct {
	CategoryImage string `json:"image" form:"image" validate:"required"`
	CategoryIcon  string `json:"icon" form:"icon" validate:"required"`
}
