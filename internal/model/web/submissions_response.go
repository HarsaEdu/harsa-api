package web

type SubmissionsResponseModule struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type SubmissionsResponseModuleMobile struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
}