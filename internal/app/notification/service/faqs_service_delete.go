package service

import "fmt"

func (notificationService *NotificationServiceImpl) Delete(id int) error {
	IfExist, _ := notificationService.NotificationRepository.FindById(id)
	if IfExist == nil {
		return fmt.Errorf("notification not found")
	}
	err := notificationService.NotificationRepository.Delete(id)

	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}
