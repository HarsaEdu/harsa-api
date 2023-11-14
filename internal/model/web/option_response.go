package web

type OptionsResForQuestion struct {
	Id              uint `json:"id"`
	Value           string `json:"value"`
	Is_right        bool `json:"is_right"`
}