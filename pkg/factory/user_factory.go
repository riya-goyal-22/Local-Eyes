package factory

import (
	"local-eyes/internal/models"
)

func CreateUser(id, username, password, userType string) *models.User {
	return &models.User{
		ID:           id,
		Username:     username,
		Password:     password,
		Type:         userType,
		Notification: make([]models.Notification, 0),
		//Notification: make(chan models.Notification, 10),
	}
}
