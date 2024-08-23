package controllers

import (
	"fmt"
	"local-eyes/constants"
	"local-eyes/internal/models"
	"local-eyes/internal/repositories"
	"local-eyes/utils"
	"strconv"
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
		fmt.Println(constants.Cyan + "\n---------------------------------")
		fmt.Println("Admin Account")
		fmt.Println("----------------------------------" + constants.Reset)
		fmt.Println(constants.Blue + "Admin actions:")
		fmt.Println("1. Delete User")
		fmt.Println("2. Delete Post")
		fmt.Println("3. List Users")
		fmt.Println("4. List Posts")
		fmt.Println("5. Exit" + constants.Reset)

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ac.ListUsers()
			ac.DeleteUser()
		case 2:
			ac.ListPosts()
			ac.DeletePost()
		case 3:
			ac.ListUsers()
		case 4:
			ac.ListPosts()
		case 5:
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func (ac *AdminController) DeleteUser() {
	userID, _ := strconv.Atoi(utils.PromptInput("Enter user ID to delete: "))

	if err := ac.UserRepo.Delete(userID); err != nil {
		fmt.Println("Error deleting user:", err)
		return
	} else {
		fmt.Println("User deleted successfully.")
	}
}

func (ac *AdminController) DeletePost() {
	postID, _ := strconv.Atoi(utils.PromptInput("Enter post ID to delete: "))

	if err := ac.PostRepo.Delete(postID); err != nil {
		fmt.Println("Error deleting post:", err)
		return
	} else {
		fmt.Println("Post deleted successfully.")
	}
}

func (ac *AdminController) ListUsers() {
	err := ac.UserRepo.UserDisplayTable()
	if err != nil {
		fmt.Println("Error displaying users table:", err)
	}
}

func (ac *AdminController) ListPosts() {
	err := ac.PostRepo.PostDisplayTable()
	if err != nil {
		fmt.Println("Error displaying posts table:", err)
	}
}
