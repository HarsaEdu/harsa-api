package web

type SuccessResponse struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Pagination any    `json:"pagination,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
}

type Pagination struct {
	Offset int   `json:"offset"`
	Limit  int   `json:"limit"`
	Total  int64 `json:"total"`
}
