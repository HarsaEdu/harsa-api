package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertAllQuestionsQuiz(questions []domain.Questions) []web.QuestionsResForQuiz {
	var questionRes []web.QuestionsResForQuiz

	for _, q := range questions {
		var optionRes []web.OptionsResForQuestion

		for _, opt := range q.Options {
			optionRes = append(optionRes, web.OptionsResForQuestion{
				Id:      opt.ID,
				Value:   opt.Value,
				IsRight: opt.IsRight,
			})
		}

		question := web.QuestionsResForQuiz{
			Id:       q.ID,
			Question: q.Question,
			Options:  optionRes,
		}

		questionRes = append(questionRes, question)
	}

	return questionRes
}
