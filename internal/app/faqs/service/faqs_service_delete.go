package service

import "fmt"

func (faqsService *FaqsServiceImpl) Delete(id int) error {
	IfExist, _ := faqsService.FaqRepository.FindById(id)
	if IfExist == nil {
		return fmt.Errorf("faqs not found")
	}
	err := faqsService.FaqRepository.Delete(id)

	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}
