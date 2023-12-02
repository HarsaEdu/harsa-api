package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"google.golang.org/api/option"
)

func SendNotificationPersonal(notif web.NotificationPersonal, config configs.AppConfig) {

	credential, err := GetDecodedFireBaseKey(config)
	if err != nil {
		fmt.Errorf("Failed to get credential: %v", err)
	}

	// Konfigurasi Firebase Admin SDK
	opt := option.WithCredentialsJSON(credential)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	// Inisialisasi FCM client
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Failed to create FCM client: %v", err)
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
	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}

	fmt.Println("Successfully sent notification:", response)
}

func SendNotificationMulticast(notif web.NotificationMultiCast, config configs.AppConfig) {

	credential, err := GetDecodedFireBaseKey(config)
	if err != nil {
		fmt.Errorf("Failed to get credential: %v", err)
	}

	// Konfigurasi Firebase Admin SDK
	opt := option.WithCredentialsJSON(credential)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	// Inisialisasi FCM client
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Failed to create FCM client: %v", err)
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
	response, err := client.SendMulticast(context.Background(), message)
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}

	fmt.Println("Successfully sent notification:", response)
}
