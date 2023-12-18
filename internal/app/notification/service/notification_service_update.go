package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (notificationRepository *NotificationServiceImpl) ReadNotification(id int) error {
	ifExist, _ := notificationRepository.NotificationRepository.FindById(id)
	if ifExist == nil {
		return fmt.Errorf("notification not found")
	}

	err := notificationRepository.NotificationRepository.ReadNotification(id)
	if err != nil {
		return fmt.Errorf("error when updating read %s", err.Error())
	}

	return nil
}

func (notificationRepository *NotificationServiceImpl) ArsipNotification(arsip web.ArsipNotification) error {
	ifExist, _ := notificationRepository.NotificationRepository.FindById(int(arsip.ID))
	if ifExist == nil {
		return fmt.Errorf("notification not found")
	}

	err := notificationRepository.NotificationRepository.ArsipNotification(arsip)
	if err != nil {
		return fmt.Errorf("error when updating read %s", err.Error())
	}

	return nil
}
