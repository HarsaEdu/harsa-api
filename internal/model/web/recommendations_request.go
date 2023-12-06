package web

type GetRecommendationsRequest struct {
	UserId     uint `json:"user_id" validate:"required"`
	MaxResults int  `json:"max_results" validate:"required"`
}