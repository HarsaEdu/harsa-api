package web

import "time"

type NotificationPersonal struct {
	Title             string `json:"title"`
	Message           string `json:"message"`
	RegistrationToken string `json:"registration_token"`
}
type NotificationMultiCast struct {
	Title             string   `json:"title"`
	Message           string   `json:"message"`
	RegistrationToken []string `json:"registration_token"`
}

type NotificationResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	IsRead    bool      `json:"is_read"`
	IsArsip   bool      `json:"is_arsip"`
}
type ReadNotification struct {
	ID uint `json:"id"`
}
type ArsipNotification struct {
	ID      uint `json:"id" params:"id"`
	IsArsip bool `json:"is_arsip"`
}
