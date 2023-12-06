package web

type GetRecommendationsRequest struct {
	UserId uint `json:"user_id" validate:"required"`
	Max    int  `json:"max" validate:"required"`
}