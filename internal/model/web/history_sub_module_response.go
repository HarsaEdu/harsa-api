package web

type SubModuleResponseForTracking struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type HistorySubModuleResponseMobile struct {
	ID         uint                         `json:"id"`
	IsComplete bool                         `json:"is_complete"`
	SubModule  SubModuleResponseForTracking `json:"sub_module"`
}
