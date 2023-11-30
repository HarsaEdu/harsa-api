package web

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type SubModuleResponseForTracking struct {
	ID          uint                 `json:"id"`
	Title       string               `json:"title"`
	Type        domain.SubModuleType `json:"type"`
	Is_complete bool                 `json:"is_complete"`
}

type HistorySubModuleResponseMobile struct {
	ID         uint                         `json:"id"`
	IsComplete bool                         `json:"is_complete"`
	SubModule  SubModuleResponseForTracking `json:"sub_module"`
}