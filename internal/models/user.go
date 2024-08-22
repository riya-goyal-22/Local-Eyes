package models

type User struct {
	ID             string
	Username       string
	Password       string
	Type           string
	Notification   []Notification
	NotificationCh chan Notification `json:"-"`
}
