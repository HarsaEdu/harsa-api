package web

type GetHistorySubModuleResponse struct {
	ID          uint `json:"id"`
	SubModuleID uint `json:"module_id"`
	IsCompleted bool `json:"is_completed"`
}
