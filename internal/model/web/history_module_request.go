package web

type HistoryModuleCreateRequest struct {
	UserID uint `json:"userId" validate:"required"`
}

type HistoryModuleUpdateRequest struct {
	UserID uint `json:"userId" validate:"required"`
}
