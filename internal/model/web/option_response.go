package web

type OptionsResForQuestion struct {
	Id      uint   `json:"id"`
	Value   string `json:"value"`
	IsRight bool   `json:"is_right"`
}

type OptionsResForQuestionMobile struct {
	Id      uint   `json:"id"`
	Value   string `json:"value"`
}
