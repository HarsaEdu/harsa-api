package web

type InterestRequest struct {
	CategoryID []uint `json:"category_id" validate:"required"`
}
