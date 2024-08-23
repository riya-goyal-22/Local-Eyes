package repositories_test

import (
	"local-eyes/internal/repositories"
	"os"
	"testing"
)

func TestFileRepository_SaveLoad(t *testing.T) {
	filePath := "test_data.json"
	defer os.Remove(filePath)

	repo := repositories.NewFileRepository(filePath)

	// Test data
	data := map[string]string{"key": "value"}

	// Save data
	err := repo.Save(data)
	if err != nil {
		t.Fatalf("Failed to save data: %v", err)
	}

	// Load data
	var loadedData map[string]string
	err = repo.Load(&loadedData)
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	// Compare
	if loadedData["key"] != data["key"] {
		t.Errorf("Loaded data mismatch. Expected %v, got %v", data, loadedData)
	}
}
