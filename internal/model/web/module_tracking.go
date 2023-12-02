package web

type SubModuleTrackingResponse struct {
	ID               uint                           `json:"id"`
	Title            string                         `json:"title"`
	Description      string                         `json:"description"`
	Progress         float32                        `json:"progress"`
	Order            int                            `json:"order"`
	HistorySubModule HistorySubModuleResponseMobile `json:"history_sub_module"`
	SubModule        SubModuleResponseModule        `json:"sub_module"`
}
