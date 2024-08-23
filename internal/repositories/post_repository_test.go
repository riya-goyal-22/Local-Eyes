package repositories_test

import (
	"io/ioutil"
	"local-eyes/internal/models"
	"local-eyes/internal/repositories"
	"os"
	"testing"
)

func TestPostRepository_SaveLoadDelete(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "test_posts_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Ensure the temp file is removed after the test

	repo := repositories.NewPostRepository(tempFile.Name()) // Use temp file for the repository

	// Test post
	post := &models.Post{ID: 1, Title: "Test Post", Content: "This is a test post.", Type: "example", LikeCount: 10, UserId: 42}

	// Save post
	err = repo.Save(post)
	if err != nil {
		t.Fatalf("Failed to save post: %v", err)
	}

	// Load posts
	posts, err := repo.Load()
	if err != nil {
		t.Fatalf("Failed to load posts: %v", err)
	}

	// Check if post is saved
	if len(posts) != 1 || posts[0].ID != post.ID {
		t.Errorf("Post not saved correctly. Expected ID %v, got %v", post.ID, posts[0].ID)
	}

	// Delete post
	err = repo.Delete(post.ID)
	if err != nil {
		t.Fatalf("Failed to delete post: %v", err)
	}

	// Load posts again
	posts, err = repo.Load()
	if err != nil {
		t.Fatalf("Failed to load posts after deletion: %v", err)
	}

	if len(posts) != 0 {
		t.Errorf("Post not deleted correctly. Expected 0 posts, got %v", len(posts))
	}
}
