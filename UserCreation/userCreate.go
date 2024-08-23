package userCreate

import (
	"local-eyes/internal/models"
)

func CreateUser(id int, username, password, userType string) *models.User {
	return &models.User{
		ID:           id,
		Username:     username,
		Password:     password,
		Type:         userType,
		Notification: make([]models.Notification, 0),
	}
}
