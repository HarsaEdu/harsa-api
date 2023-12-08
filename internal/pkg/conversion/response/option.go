package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertOptionsResForQuestion(option *domain.Options) *web.OptionsResForQuestion {
	optionRes := web.OptionsResForQuestion{
		Id:      option.ID,
		Value:   option.Value,
		IsRight: option.IsRight,
	}
	return &optionRes
}
