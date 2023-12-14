package firebase

import (
	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

type Firebase interface {
	SendNotificationPersonal(notif *web.NotificationPersonal)
	SendNotificationMulticast(notif *web.NotificationMultiCast)
}

type FirebaseImpl struct {
	Config *configs.FirebaseConfig
}

func NewFirebase(config *configs.FirebaseConfig) Firebase {
	return &FirebaseImpl{
		Config: config,
	}
}