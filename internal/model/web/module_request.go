package web

type ModuleCreateRequest struct {
	Title       string                   `json:"title" form:"title" validate:"required"`
	Description string                   `json:"description" form:"description"`
	Type        string                   `json:"type" form:"type" validate:"required"`
	Order       int                      `json:"order" form:"order" validate:"required,min=1"`
	SubModules  []SubModuleCreateRequest `json:"sub_module" form:"sub_module"`
}

type SubModuleCreateRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Type        string `json:"type" form:"type" validate:"required"`
	ContentUrl  string `json:"content_url" form:"content_url"`
	ContentBody string `json:"content_body" form:"content_body"`
}