package repositories

import (
	"encoding/json"
	"os"
	"sync"
)

type Repository interface {
	Save(interface{}) error
	Load() ([]interface{}, error)
	Delete(string) error
	Update(interface{}) error
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
	if err := decoder.Decode(data); err != nil {
		return err
	}
	return nil
}

func (r *FileRepository) Delete(id string) error {
	// Implementation depends on the structure of data and how it is managed
	return nil
}

func (r *FileRepository) Update(data interface{}) error {
	// Implementation depends on the structure of data and how it is managed
	return nil
}
