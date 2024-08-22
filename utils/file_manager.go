package utils

import (
	"encoding/json"
	"os"
)

const UserFile = "../../data/users.json"
const PostFile = "../../data/posts.json"

func ReadFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(v)
}

func WriteFile(filePath string, v interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(v)
}
