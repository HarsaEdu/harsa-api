package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func FaqsCreateRequestToFaqsDomain(faqs *web.FaqsRequest) *domain.Faqs {
	return &domain.Faqs{Question: faqs.Question, Answer: faqs.Answer}
}

func FaqsUpdateRequestToFaqsDomain(faqs *web.FaqsUpdateRequest) *domain.Faqs {
	return &domain.Faqs{Question: faqs.Question, Answer: faqs.Answer}
}
