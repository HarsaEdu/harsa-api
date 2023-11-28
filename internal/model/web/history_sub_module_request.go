package web

type CreateHistorySubModuleRequest struct {
	SubModuleID uint `json:"sub_module_id" validate:"required"`
}

type GetHistorySubModuleRequest struct {
	UserID uint `json:"user_id" validate:"required"`
}

type UpdateHistorySubModuleRequest struct {
	ID uint `json:"id" validate:"required"`
}
