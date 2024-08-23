package repositories_test

import (
	"io/ioutil"
	"local-eyes/internal/models"
	"local-eyes/internal/repositories"
	"os"
	"testing"
)

func TestUserRepository_SaveLoadDelete(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "test_users_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Ensure the temp file is removed after the test

	repo := repositories.NewUserRepository(tempFile.Name()) // Use temp file for the repository

	// Test user
	user := &models.User{ID: 1, Username: "Test User", Password: "test@example.com", Type: "newbie"}

	// Save user
	err = repo.Save(user)
	if err != nil {
		t.Fatalf("Failed to save user: %v", err)
	}

	// Load users
	users, err := repo.Load()
	if err != nil {
		t.Fatalf("Failed to load users: %v", err)
	}

	// Check if user is saved
	if len(users) != 1 || users[0].ID != user.ID {
		t.Errorf("User not saved correctly. Expected ID %v, got %v", user.ID, users[0].ID)
	}

	// Delete user
	err = repo.Delete(user.ID)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// Load users again
	users, err = repo.Load()
	if err != nil {
		t.Fatalf("Failed to load users after deletion: %v", err)
	}

	if len(users) != 0 {
		t.Errorf("User not deleted correctly. Expected 0 users, got %v", len(users))
	}
}
