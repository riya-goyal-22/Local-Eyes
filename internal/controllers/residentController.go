package controllers

import (
	"fmt"
	"local-eyes/constants"
	"local-eyes/internal/models"
	"local-eyes/internal/repositories"
	"local-eyes/utils"
	"strconv"
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
		fmt.Println(constants.Cyan + "\n---------------------------------")
		fmt.Println("Resident Account")
		fmt.Println("----------------------------------" + constants.Reset)
		fmt.Println(constants.Blue + "Resident actions:")
		fmt.Println("1. Create Post")
		fmt.Println("2. Update Post")
		fmt.Println("3. Delete Post")
		fmt.Println("4. View Post")
		fmt.Println("5. Exit" + constants.Reset)

		var choice int
		fmt.Scanln(&choice)

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
	title := utils.PromptInput("Enter post title: ")
	content := utils.PromptInput("Enter post content: ")
	pType := utils.PromptInput("Enter post type: ")
	post := &models.Post{
		ID:        utils.GeneratePostId(),
		Title:     title,
		Content:   content,
		Type:      pType,
		LikeCount: 0,
		UserId:    rc.User.ID,
	}
	if err := rc.PostRepo.Save(post); err != nil {
		fmt.Println("Error creating post:", err)
	} else {
		rc.notify.NotifyNewPost(post)
		fmt.Println("Post created successfully.")
	}
}

func (rc *ResidentController) UpdatePost() {
	postId, _ := strconv.Atoi(utils.PromptInput("Enter post ID to update: "))
	title := utils.PromptInput("Enter new post title: ")
	content := utils.PromptInput("Enter new post content: ")
	pType := utils.PromptInput("Enter new post type: ")

	post := &models.Post{
		ID:        postId,
		Title:     title,
		Content:   content,
		Type:      pType,
		LikeCount: 0,
		UserId:    rc.User.ID,
	}
	if err := rc.PostRepo.Update(post); err != nil {
		fmt.Println("Error updating post:", err)
	} else {
		fmt.Println("Post updated successfully.")
	}
}

func (rc *ResidentController) DeletePost() {
	postID, err := strconv.Atoi(utils.PromptInput("Enter post ID to delete: "))
	if err != nil {
		fmt.Println("Invalid post ID")
	}
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
		if post.UserId == rc.User.ID {
			fmt.Printf("Post ID: %d\nTitle: %s\nContent: %s\nType: %s\n", post.ID, post.Title, post.Content, post.Type)
		}
	}
}
