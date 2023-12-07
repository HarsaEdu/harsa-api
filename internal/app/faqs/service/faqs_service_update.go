package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (faqsService *FaqsServiceImpl) Update(faqs web.FaqsUpdateRequest, id int) error {
	err := faqsService.Validator.Struct(faqs)
	if err != nil {
		return err
	}

	ifExist, _ := faqsService.FaqRepository.FindById(id)
	if ifExist == nil {
		return fmt.Errorf("faqs not found")
	}

	result := conversion.FaqsUpdateRequestToFaqsDomain(&faqs)

	err = faqsService.FaqRepository.Update(result, id)
	if err != nil {
		return fmt.Errorf("error when updating category %s", err.Error())
	}

	return nil
}
