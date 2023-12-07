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

type SubModuleTracking struct {
	ID               uint                     `json:"id"`
	Title            string                   `json:"title"`
	Description      string                   `json:"description"`
	Progress         float32                  `json:"progress"`
	Order            int                      `json:"order"`
	HistorySubModule HistorySubModuleTracking `json:"history_sub_module"`
	SubModule        SubModuleResponseModule  `json:"sub_module"`
}

type ModuleTrackingByID struct {
	ID          uint                              `json:"id"`
	Title       string                            `json:"title"`
	Description string                            `json:"description"`
	Progress    float32                           `json:"progress"`
	Order       int                               `json:"order"`
	SubModules  []SubModuleResponseForTracking    `json:"sub_modules"`
	Submissions []SubmissionsResponseModuleMobile `json:"submissions"`
	Quizzes     []QuizResponseForTracking         `json:"quizzes"`
}
