package repositories

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

type FileRepositoryInterface interface {
	Save(interface{}) error
	Load(interface{}) error
}

type FileRepository struct {
	FilePath string
	mu       sync.Mutex
}

func NewFileRepository(filePath string) *FileRepository {
	return &FileRepository{FilePath: filePath}
}

func (r *FileRepository) Save(data interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	file, err := os.OpenFile(r.FilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}

func (r *FileRepository) Load(data interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	file, err := os.Open(r.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		// Handle the EOF error as a special case, since it indicates an empty file
		if err == io.EOF {
			return nil // No data to load, so it's not considered an error
		}
		return err
	}
	return nil
}
