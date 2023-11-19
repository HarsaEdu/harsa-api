package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (faqsService *FaqsServiceImpl) Create(faqs web.FaqsRequest) error {
	err := faqsService.Validator.Struct(faqs)
	if err != nil {
		return err
	}

	result := conversion.FaqsCreateRequestToFaqsDomain(&faqs)

	err = faqsService.FaqRepository.Create(result)
	if err != nil {
		return fmt.Errorf("error when creating category %s", err.Error())
	}

	return nil

}
