package web

type SubsPlanResposne struct {
	ID            uint    `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Price         float64 `json:"price" `
	Image_url     string  `json:"image"`
	Duration_days int     `json:"duration"`
}