package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func (firebaseImpl *FirebaseImpl) SendNotificationPersonal(notif *web.NotificationPersonal) {

	credential, err := GetDecodedFireBaseKey(*firebaseImpl.Config)
	if err != nil {
		logrus.Error("Failed to get credential: %v", err)
	}

	// Konfigurasi Firebase Admin SDK
	opt := option.WithCredentialsJSON(credential)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logrus.Error("Failed to create Firebase app: %v", err)
	}

	// Inisialisasi FCM client
	client, err := app.Messaging(context.Background())
	if err != nil {
		logrus.Error("Failed to create FCM client: %v", err)
	}

	// Data pesan yang akan dikirim
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: notif.Title,
			Body:  notif.Message,
		},
		Token: notif.RegistrationToken, // Token perangkat klien yang akan menerima notifikasi
	}

	// Kirim notifikasi
	_, err = client.Send(context.Background(), message)
	if err != nil {
		logrus.Error("Failed to send notification: %v", err)
	}

}

func (firebaseImpl *FirebaseImpl) SendNotificationMulticast(notif *web.NotificationMultiCast) {

	credential, err := GetDecodedFireBaseKey(*firebaseImpl.Config)
	if err != nil {
		logrus.Error("Failed to get credential: %v", err)
	}

	// Konfigurasi Firebase Admin SDK
	opt := option.WithCredentialsJSON(credential)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logrus.Error("Failed to create Firebase app: %v", err)
	}

	// Inisialisasi FCM client
	client, err := app.Messaging(context.Background())
	if err != nil {
		logrus.Error("Failed to create FCM client: %v", err)
	}

	// Data pesan yang akan dikirim
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: notif.Title,
			Body:  notif.Message,
		},
		Tokens: notif.RegistrationToken, // Token perangkat klien yang akan menerima notifikasi
	}

	// Kirim notifikasi
	_, err = client.SendMulticast(context.Background(), message)
	if err != nil {
		logrus.Error("Failed to send notification: %v", err)
	}

}
