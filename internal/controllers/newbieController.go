package controllers

import (
	"fmt"
	"local-eyes/constants"
	"local-eyes/internal/models"
	"local-eyes/internal/repositories"
	"local-eyes/utils"
	"strconv"
)

type NewbieController struct {
	User     *models.User
	PostRepo *repositories.PostRepository
	Notify   *NotificationController
}

func NewNewbieController(user *models.User, postRepo *repositories.PostRepository, notify *NotificationController) *NewbieController {
	return &NewbieController{User: user, PostRepo: postRepo, Notify: notify}
}

func (nc *NewbieController) HandleNewbieActions() {
	for _, n := range nc.User.Notification {
		nc.User.NotificationCh <- n
	}
	nc.Notify.RemoveNotification(nc.User)

	select {
	case val := <-nc.User.NotificationCh:
		fmt.Println(constants.Gray+"\nYour notification:", val.Message)
		fmt.Print(constants.Reset)
	default:
		break
	}

	for {
		fmt.Println(constants.Cyan + "\n---------------------------------")
		fmt.Println("Newbie Account")
		fmt.Println("----------------------------------" + constants.Reset)
		fmt.Println(constants.Blue + "Newbie actions:")
		fmt.Println("1. View Posts")
		fmt.Println("2. View Post with filter")
		fmt.Println("3. Like Post")
		fmt.Println("4. Exit" + constants.Reset)

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			nc.ViewPosts()
		case 2:
			nc.ViewFilterPost()
		case 3:
			nc.LikePost()
		case 4:
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func (nc *NewbieController) ViewPosts() {

	err := nc.PostRepo.PostDisplayTable()
	if err != nil {
		fmt.Println("Error displaying posts:", err)
	}

}

func (nc *NewbieController) LikePost() {
	postID, _ := strconv.Atoi(utils.PromptInput("Enter post ID to like: "))

	if err := nc.PostRepo.Like(postID); err != nil {
		fmt.Println("Error liking post:", err)
	} else {
		fmt.Println("Post liked successfully.")
	}
}

func (nc *NewbieController) ViewFilterPost() {
	postType := utils.PromptInput("Enter post type: ")
	posts, err := nc.PostRepo.Load()
	if err != nil {
		fmt.Println("Error loading posts:", err)
		return
	}

	for _, post := range posts {
		if post.Type == postType {
			err := nc.PostRepo.PostDisplayTable()
			if err != nil {
				fmt.Println("Error displaying post:", err)
			}
		}
	}
}
