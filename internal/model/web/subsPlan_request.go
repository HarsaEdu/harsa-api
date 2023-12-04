package web

type SubsPlanCreateRequest struct {
	Title         string  `json:"title" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	Image_url     string  `form:"image"`
	Duration_days int     `json:"duration" validate:"required"`
}

type SubsPlanUpdateRequest struct {
	Title         string  `json:"title" `
	Description   string  `json:"description" `
	Price         float64 `json:"price" `
	Duration_days int     `json:"duration" `
}

type SubsPlanUpdateImage struct {
	Image_url string `form:"image" validate:"required"`
}
