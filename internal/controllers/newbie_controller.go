package controllers

import (
	"fmt"
	"local-eyes/internal/models"
	"local-eyes/internal/repositories"
)

type NewbieController struct {
	User     *models.User
	PostRepo *repositories.PostRepository
}

func NewNewbieController(user *models.User, postRepo *repositories.PostRepository) *NewbieController {
	return &NewbieController{User: user, PostRepo: postRepo}
}

func (nc *NewbieController) HandleNewbieActions() {
	for _, n := range nc.User.Notification {
		nc.User.NotificationCh <- n
	}
	<-nc.User.NotificationCh
	for {
		fmt.Println("\nNewbie actions:")
		fmt.Println("1. View Posts")
		fmt.Println("2. Like Post")
		fmt.Println("3. Exit")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			nc.ViewPosts()
		case 2:
			nc.LikePost()
		case 3:
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func (nc *NewbieController) ViewPosts() {
	posts, err := nc.PostRepo.Load()
	if err != nil {
		fmt.Println("Error loading posts:", err)
		return
	}

	for _, post := range posts {
		fmt.Printf("Post ID: %s\nTitle: %s\nContent: %s\n\n", post.ID, post.Title, post.Content)
	}
}

func (nc *NewbieController) LikePost() {
	var postID string
	fmt.Print("Enter post ID to like: ")
	fmt.Scan(&postID)

	if err := nc.PostRepo.Like(postID); err != nil {
		fmt.Println("Error liking post:", err)
	} else {
		fmt.Println("Post liked successfully.")
	}
}

func (nc *NewbieController) ViewFilterPost() {
	fmt.Println("\nEnter post type:")
	var postType string
	fmt.Scan(&postType)
	posts, err := nc.PostRepo.Load()
	if err != nil {
		fmt.Println("Error loading posts:", err)
		return
	}

	for _, post := range posts {
		if post.Type == postType {
			fmt.Printf("Post ID: %s\nTitle: %s\nContent: %s\n\n", post.ID, post.Title, post.Content)
		}
	}
}
