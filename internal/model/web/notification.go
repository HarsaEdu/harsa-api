package web

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
