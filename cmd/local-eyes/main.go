package main

import (
	"local-eyes/cli"
	"local-eyes/internal/controllers"
	"local-eyes/internal/repositories"
)

func main() {
	userRepo := repositories.NewUserRepository()
	postRepo := repositories.NewPostRepository()
	notificationCtrl := controllers.NewNotificationController(userRepo)

	cli.StartCLI(userRepo, postRepo, notificationCtrl)
}
