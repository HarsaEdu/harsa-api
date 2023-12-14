package web

type CategoryCreateRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description"`
	Image_Url   string `json:"image" form:"image" `
	Icon        string `json:"icon" form:"icon" `
}

type CategoryUpdateRequest struct {
	Name          string `json:"name" form:"name"`
	Description   string `json:"description" form:"description" `
	CategoryImage string `json:"image" form:"image"`
}

type CategoryUploadImageRequest struct {
	CategoryImage string `json:"image" form:"image" validate:"required"`
	CategoryIcon  string `json:"icon" form:"icon" validate:"required"`
}
