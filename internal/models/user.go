package models

type User struct {
	ID             int
	Username       string
	Password       string
	Type           string
	Notification   []Notification
	NotificationCh chan Notification `json:"-"`
}
