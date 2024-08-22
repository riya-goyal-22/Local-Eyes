package controllers

import (
	"fmt"
	"local-eyes/internal/models"
	"local-eyes/internal/repositories"
	"local-eyes/utils"
)

type ResidentController struct {
	User     *models.User
	PostRepo *repositories.PostRepository
	notify   *NotificationController
}

func NewResidentController(user *models.User, postRepo *repositories.PostRepository, notify *NotificationController) *ResidentController {
	return &ResidentController{User: user, PostRepo: postRepo, notify: notify}
}

func (rc *ResidentController) HandleResidentActions() {
	for {
		fmt.Println("\nResident actions:")
		fmt.Println("1. Create Post")
		fmt.Println("2. Update Post")
		fmt.Println("3. Delete Post")
		fmt.Println("4. View Post")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			rc.CreatePost()
		case 2:
			rc.ViewPost()
			rc.UpdatePost()
		case 3:
			rc.ViewPost()
			rc.DeletePost()
		case 4:
			rc.ViewPost()
		case 5:
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func (rc *ResidentController) CreatePost() {
	var title, content, pType string
	fmt.Print("Enter post title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter post content: ")
	fmt.Scanln(&content)
	fmt.Print("Enter post type: ")
	fmt.Scanln(&pType)

	post := &models.Post{
		ID:        utils.GenerateID(),
		Title:     title,
		Content:   content,
		Type:      pType,
		LikeCount: 0,
	}
	if err := rc.PostRepo.Save(post); err != nil {
		fmt.Println("Error creating post:", err)
	} else {
		rc.notify.NotifyNewPost(post)
		fmt.Println("Post created successfully.")
	}
}

func (rc *ResidentController) UpdatePost() {
	var postID, title, content, pType string
	fmt.Print("Enter post ID to update: ")
	fmt.Scanln(&postID)
	fmt.Print("Enter new title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter new content: ")
	fmt.Scanln(&content)
	fmt.Print("Enter new post type: ")
	fmt.Scanln(&pType)

	post := &models.Post{
		ID:        postID,
		Title:     title,
		Content:   content,
		Type:      pType,
		LikeCount: 0,
	}
	if err := rc.PostRepo.Update(post); err != nil {
		fmt.Println("Error updating post:", err)
	} else {
		fmt.Println("Post updated successfully.")
	}
}

func (rc *ResidentController) DeletePost() {
	var postID string
	fmt.Print("Enter post ID to delete: ")
	fmt.Scanln(&postID)
	if err := rc.PostRepo.Delete(postID); err != nil {
		fmt.Println("Error deleting post:", err)
	} else {
		fmt.Println("Post deleted successfully.")
	}

}

func (rc *ResidentController) ViewPost() {
	posts, err := rc.PostRepo.Load()
	if err != nil {
		fmt.Println("Error loading posts:", err)
	}
	for _, post := range posts {
		if post.ID == rc.User.ID {
			fmt.Printf("Post ID: %s\nTitle: %s\nContent: %s\n\n", post.ID, post.Title, post.Content)
		}
	}
}
