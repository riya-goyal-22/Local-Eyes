package controllers

import (
	"fmt"
	"local-eyes/internal/models"
	"local-eyes/internal/repositories"
)

type NotificationController struct {
	userRepo *repositories.UserRepository
}

func NewNotificationController(userRepo *repositories.UserRepository) *NotificationController {
	return &NotificationController{
		userRepo: userRepo,
	}
}

func (nc *NotificationController) NotifyNewPost(post *models.Post) {
	newbies, err := nc.userRepo.FindAllNewbies()
	if err != nil {
		return
	}
	for i, newbie := range newbies {
		notification := models.Notification{
			Message: fmt.Sprintf("Notifying Newbie %s about new post: %s\n", newbie.Username, post.Title),
		}
		newbies[i].Notification = append(newbies[i].Notification, notification)
	}
}
