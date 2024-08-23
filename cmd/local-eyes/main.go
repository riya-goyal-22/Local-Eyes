package main

import (
	"local-eyes/cli"
	"local-eyes/constants"
	"local-eyes/internal/controllers"
	"local-eyes/internal/repositories"
)

func main() {
	userRepo := repositories.NewUserRepository(constants.UserFile)
	postRepo := repositories.NewPostRepository(constants.PostFile)
	notificationCtrl := controllers.NewNotificationController(userRepo)

	cli.StartCLI(userRepo, postRepo, notificationCtrl)
}
