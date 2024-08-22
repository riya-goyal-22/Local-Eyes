package controllers

import (
	"fmt"
	"local-eyes/internal/models"
	"local-eyes/internal/repositories"
)

type AdminController struct {
	User     *models.User
	UserRepo *repositories.UserRepository
	PostRepo *repositories.PostRepository
}

func NewAdminController(user *models.User, userRepo *repositories.UserRepository, postRepo *repositories.PostRepository) *AdminController {
	return &AdminController{User: user, UserRepo: userRepo, PostRepo: postRepo}
}

func (ac *AdminController) HandleAdminActions() {
	for {
		fmt.Println("\nAdmin actions:")
		fmt.Println("1. Delete User")
		fmt.Println("2. Delete Post")
		fmt.Println("3. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ac.DeleteUser()
		case 2:
			ac.DeletePost()
		case 3:
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func (ac *AdminController) DeleteUser() {
	var userID string
	fmt.Print("Enter user ID to delete: ")
	fmt.Scan(&userID)

	if err := ac.UserRepo.Delete(userID); err != nil {
		fmt.Println("Error deleting user:", err)
	} else {
		fmt.Println("User deleted successfully.")
	}
}

func (ac *AdminController) DeletePost() {
	var postID string
	fmt.Print("Enter post ID to delete: ")
	fmt.Scan(&postID)

	if err := ac.PostRepo.Delete(postID); err != nil {
		fmt.Println("Error deleting post:", err)
	} else {
		fmt.Println("Post deleted successfully.")
	}
}
