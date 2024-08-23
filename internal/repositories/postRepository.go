package repositories

import (
	"errors"
	"fmt"
	"local-eyes/internal/models"
	"strings"
)

type PostRepositoryInterface interface {
	Save(post *models.Post) error
	Load() ([]*models.Post, error)
	Delete(postID int) error
	Update(post *models.Post) error
	Like(postID int) error
	SaveAll(posts []*models.Post) error
	PostDisplayTable() error
}

type PostRepository struct {
	FileRepository
}

func NewPostRepository(filepath string) *PostRepository {
	return &PostRepository{
		FileRepository: *NewFileRepository(filepath),
	}
}

func (pr *PostRepository) Save(post *models.Post) error {
	posts, err := pr.Load()
	if err != nil {
		return err
	}
	posts = append(posts, post)
	return pr.SaveAll(posts)
}

func (pr *PostRepository) Load() ([]*models.Post, error) {
	var posts []*models.Post
	err := pr.FileRepository.Load(&posts)
	return posts, err
}

func (pr *PostRepository) Delete(postID int) error {
	posts, err := pr.Load()
	if err != nil {
		return err
	}
	for i, post := range posts {
		if post.ID == postID {
			posts = append(posts[:i], posts[i+1:]...)
			return pr.SaveAll(posts)
		}
	}
	return errors.New("post not found")
}

func (pr *PostRepository) Update(post *models.Post) error {
	posts, err := pr.Load()
	if err != nil {
		return err
	}
	for i, p := range posts {
		if p.ID == post.ID {
			posts[i] = post
			return pr.SaveAll(posts)
		}
	}
	return nil
}

func (pr *PostRepository) Like(postID int) error {
	posts, err := pr.Load()
	if err != nil {
		return err
	}
	for _, post := range posts {
		if post.ID == postID {
			post.LikeCount++
			return pr.SaveAll(posts)
		}
	}
	return errors.New("post not found")
}

func (pr *PostRepository) SaveAll(posts []*models.Post) error {
	return pr.FileRepository.Save(posts)
}

func (pr *PostRepository) PostDisplayTable() error {
	// Define column headers
	headers := []string{"Id", "Title", "Content", "Type", "LikeCount"}

	// Define column widths
	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = len(header)
	}

	posts, err := pr.Load()
	if err != nil {
		return err
	}

	// Compute the maximum width needed for each column based on all posts
	for _, postElement := range posts {
		values := []string{
			fmt.Sprintf("%v", postElement.ID),
			postElement.Title,
			postElement.Content,
			string(postElement.Type),
			fmt.Sprintf("%d", postElement.LikeCount),
		}
		for i, value := range values {
			if len(value) > colWidths[i] {
				colWidths[i] = len(value)
			}
		}
	}
	// Print the headers
	for i, header := range headers {
		fmt.Printf("%-*s ", colWidths[i], header)
	}
	fmt.Println()

	// Print the separator line
	fmt.Println(strings.Repeat("-", sum(colWidths)+len(colWidths)-1))
	// Print the post details
	for _, postElement := range posts {
		values := []string{
			fmt.Sprintf("%d", postElement.ID),
			postElement.Title,
			postElement.Content,
			string(postElement.Type),
			fmt.Sprintf("%d", postElement.LikeCount),
		}
		for i, value := range values {
			fmt.Printf("%-*s ", colWidths[i], value)
		}
		fmt.Println()
	}
	fmt.Println()
	return nil
}

// sum calculates the sum of integers in a slice
func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
